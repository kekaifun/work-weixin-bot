package chatgpt

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httputil"
	urlpkg "net/url"
	"os"
	"strings"

	"github.com/kekaifun/work-weixin-bot/model/openai"
	"github.com/kekaifun/work-weixin-bot/pkg"

	"github.com/gin-gonic/gin"
)

var defaultIP = "127.0.0.1"

type OpenAIHandler struct {
	// cache files
	CacheFile map[string]*os.File
}

func (o *OpenAIHandler) Completions(c *gin.Context) {
	proxy := httputil.ReverseProxy{
		Director: func(request *http.Request) {
			o.recordRequest(request)

			url := "https://api.openai.com" + request.URL.Path
			u, _ := urlpkg.Parse(url)
			request.Host = u.Host
			request.URL = u
		},
		ModifyResponse: func(response *http.Response) error {
			o.recordResponse(response.Request, response)
			return nil
		},
	}

	proxy.ServeHTTP(c.Writer, c.Request)
}

func (o *OpenAIHandler) recordResponse(request *http.Request, response *http.Response) error {
	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println("completions: read response body failed")
		return err
	}
	response.Body = io.NopCloser(bytes.NewReader(responseBody))
	if request.Method != http.MethodPost {
		return nil
	}
	var chatResp = &openai.ChatResponse{}
	err = json.Unmarshal(responseBody, chatResp)
	if err != nil {
		fmt.Println("unmarshal json response failed")
		return err
	}
	userResp := chatResp.Choices[0].Message
	clientIP := pkg.ClientIP(request)
	if clientIP == "" {
		clientIP = defaultIP
	}
	clientRecordFile := strings.Replace(clientIP, ".", "_", -1)

	return o.WriteMsg("response", userResp.Content, clientRecordFile)
}

func (o *OpenAIHandler) recordRequest(request *http.Request) error {
	requestBody, err := io.ReadAll(request.Body)
	if err != nil {
		fmt.Println("completions: read request body failed")
		return err
	}
	// fmt.Println("requestBody: " + string(requestBody))
	request.Body = io.NopCloser(bytes.NewReader(requestBody))
	if request.Method != http.MethodPost {
		return nil
	}

	var chatReq = &openai.ChatRequest{}
	err = json.Unmarshal(requestBody, chatReq)
	if err != nil {
		fmt.Println("unmarshal json request failed")
		return err
	}
	msgSize := len(chatReq.Messages)
	userMsg := chatReq.Messages[msgSize-1]
	clientIP := pkg.ClientIP(request)
	if clientIP == "" {
		clientIP = defaultIP
	}
	clientRecordFile := strings.Replace(clientIP, ".", "_", -1)

	return o.WriteMsg("request", userMsg.Content, clientRecordFile)
}

func (o *OpenAIHandler) WriteMsg(msgType string, msg string, fileName string) error {
	if f, ok := o.CacheFile[fileName]; ok {
		f.WriteString(fmt.Sprintf("%s: %s\n", msgType, msg))
		return nil
	}

	f, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("create file failed, err: %v\n", err)
		return err
	}
	f.WriteString(fmt.Sprintf("%s: %s\n", msgType, msg))
	o.CacheFile[fileName] = f
	return nil
}
