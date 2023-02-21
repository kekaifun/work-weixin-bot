package handlers

import (
	"bytes"
	"crypto/sha1"
	"encoding/xml"
	"fmt"
	"github.com/sbzhu/weworkapi_golang/wxbizmsgcrypt"
	"io"
	"log"
	"net/http"
	"sort"

	"github.com/gin-gonic/gin/render"

	"github.com/gin-gonic/gin"
	"github.com/kekaifun/work-weixin-bot/config"
	"github.com/kekaifun/work-weixin-bot/model"
)

// MessageHandler handler message request
func (b *Bot) MessageHandler(c *gin.Context) {
	verifyMsgSign := c.Query("msg_signature")
	verifyTimestamp := c.Query("timestamp")
	verifyNonce := c.Query("nonce")

	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"Error": err,
		})
		return
	}
	log.Printf("messageHandler: %s\n", string(body))
	wxcpt := wxbizmsgcrypt.NewWXBizMsgCrypt(config.Token, config.EncodingAESKey, "", wxbizmsgcrypt.XmlType)

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
	var msgContent model.MsgContent
	err = xml.Unmarshal(msg, &msgContent)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"Error": err,
		})
		return
	}
	log.Printf("got message: %v\n", msgContent)

	reply := model.ReplyMsgContent{
		Text: model.ResponseText{
			Content: wxbizmsgcrypt.CDATA{Value: msgContent.Text.Content},
			MentionedList: model.MentionedList{
				Item: []wxbizmsgcrypt.CDATA{{
					Value: msgContent.From.UserId,
				}},
			},
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

func calSignature(timestamp, nonce, data string) string {
	sort_arr := []string{config.Token, timestamp, nonce, data}
	sort.Strings(sort_arr)
	var buffer bytes.Buffer
	for _, value := range sort_arr {
		buffer.WriteString(value)
	}

	sha := sha1.New()
	sha.Write(buffer.Bytes())
	signature := fmt.Sprintf("%x", sha.Sum(nil))
	return string(signature)
}
