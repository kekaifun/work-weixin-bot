package handlers

import (
	"encoding/xml"
	"github.com/gin-gonic/gin/render"
	"io"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kekaifun/work-weixin-bot/config"
	"github.com/kekaifun/work-weixin-bot/model"
	"github.com/kekaifun/work-weixin-bot/wxcrypt"
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
	wxcpt := wxcrypt.NewWXBizMsgCrypt(config.Token, config.EncodingAESKey, "", wxcrypt.XmlType)
	var msg []byte
	msg, err = wxcpt.DecryptMsg(verifyMsgSign, verifyTimestamp, verifyNonce, body)
	if err != nil {
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
		XMLName: xml.Name{Local: "xml"},
		Text: model.Text{
			Content:       msgContent.Text.Content,
			MentionedList: []model.Item{{Value: msgContent.From.UserId}},
		},
		MsgType: msgContent.MsgType,
	}
	replyXml, _ := xml.Marshal(&reply)

	encryptMsg, cryptErr := wxcpt.EncryptMsg(string(replyXml), verifyTimestamp, verifyNonce)
	if cryptErr != nil {
		log.Printf("encrypt msg failed, err:%v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"Error": err,
		})
		return
	}
	signature := wxcpt.CalSignature(verifyTimestamp, verifyNonce, string(encryptMsg))

	responseXml := model.MessageResponse{
		Encrypt:      string(encryptMsg),
		MsgSignature: signature,
		TimeStamp:    verifyTimestamp,
		Nonce:        verifyNonce,
	}
	respBytes, _ := xml.Marshal(responseXml)
	c.Render(200, render.Data{
		Data: respBytes,
	})
}
