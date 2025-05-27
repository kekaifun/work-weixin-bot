package handlers

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin/render"
	work_weixin "github.com/kekaifun/work-weixin-bot/model/work-weixin"
	"github.com/sbzhu/weworkapi_golang/wxbizmsgcrypt"

	"github.com/gin-gonic/gin"
)

// MessageHandler handler message request
func (b *Bot) MessageHandler(c *gin.Context) {
	verifyMsgSign := c.Query("msg_signature")
	verifyTimestamp := c.Query("timestamp")
	verifyNonce := c.Query("nonce")
	secret := b.WechatBotConfig.Secret
	enableAsyncResponse := b.WechatBotConfig.AsyncResponse

	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"Error": err,
		})
		return
	}
	log.Printf("messageHandler: %s\n", string(body))
	wxcpt := wxbizmsgcrypt.NewWXBizMsgCrypt(secret.Token, secret.EncodingAESKey, "", wxbizmsgcrypt.XmlType)

	var msg []byte
	var cryptErr *wxbizmsgcrypt.CryptError
	msg, cryptErr = wxcpt.DecryptMsg(verifyMsgSign, verifyTimestamp, verifyNonce, body)
	if cryptErr != nil {
		log.Printf("decrypt msg failed, err:%v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"Error": err,
		})
		return
	}

	log.Printf("messageHandler: decryptMessage: %s\n", string(msg))
	var msgContent work_weixin.MsgContent
	err = xml.Unmarshal(msg, &msgContent)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"Error": err,
		})
		return
	}
	log.Printf("got message: %v\n", msgContent)
	if msgContent.ChatType == "group" && msgContent.MsgType == "text" {
		msgContent.Text.Content = strings.TrimPrefix(msgContent.Text.Content, "@"+b.Name+" ")
	}
	if enableAsyncResponse {
		go b.asyncResponse(msgContent)
		c.Status(http.StatusOK)
		return
	}

	reply := work_weixin.ReplyMsgContent{
		Text: work_weixin.ResponseText{
			Content: wxbizmsgcrypt.CDATA{Value: msgContent.Text.Content},
		},
		MsgType: msgContent.MsgType,
	}
	replyXml, _ := xml.Marshal(&reply)

	fmt.Println("replyXml: " + string(replyXml))

	encryptMsg, cryptErr := wxcpt.EncryptMsg(string(replyXml), verifyTimestamp, verifyNonce)
	if cryptErr != nil {
		log.Printf("encrypt msg failed, err:%v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"Error": err,
		})
		return
	}

	c.Render(200, render.Data{
		Data: encryptMsg,
	})
}

func (b *Bot) asyncResponse(msgContent work_weixin.MsgContent) {

	jsonData, err := json.Marshal(msgContent)
	if err != nil {
		log.Printf("marshal message content failed: %v", err)
		return
	}

	resp, err := http.Post(b.WechatBotConfig.BackendServer.Url, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		log.Printf("send async request failed: %v", err)
		return
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("read async response body failed: %v", err)
		return
	}
	log.Printf("async response body: %s", string(body))
	// body: {"status":"ok","response":"你好！有什么可以帮您的吗？😊"}, 需要解析出response
	response := struct {
		Status   string `json:"status"`
		Response string `json:"response"`
	}{}
	json.Unmarshal(body, &response)
	log.Printf("response: %s", response.Response)

	reply := work_weixin.PostMessage{
		ChatId:  msgContent.ChatId,
		MsgType: msgContent.MsgType,
		Text: work_weixin.TextJson{
			Content:       response.Response,
			MentionedList: []string{msgContent.From.UserId},
		},
	}
	replyJson, _ := json.Marshal(reply)
	log.Printf("replyJson: %s", string(replyJson))

	resp, err = http.Post(msgContent.WebhookUrl, "application/json", bytes.NewBuffer(replyJson))
	if err != nil {
		log.Printf("send async request failed: %v", err)
		return
	}
	defer resp.Body.Close()
	body, err = io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("read async response body failed: %v", err)
		return
	}
	log.Printf("async response body: %s", string(body))
}
