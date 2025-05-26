package handlers

import (
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
	secret := b.WechatBotSecret

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
