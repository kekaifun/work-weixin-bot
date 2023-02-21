package handlers

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/render"
	"github.com/kekaifun/work-weixin-bot/config"
	"github.com/kekaifun/work-weixin-bot/model"
	"github.com/kekaifun/work-weixin-bot/wxcrypt"
	"io"
	"log"
	"net/http"
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
	wxcpt := wxcrypt.NewWXBizMsgCrypt(config.Token, config.EncodingAESKey, "", wxcrypt.JsonType)
	var msg []byte
	msg, err = wxcpt.DecryptMsg(verifyMsgSign, verifyTimestamp, verifyNonce, body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"Error": err,
		})
		return
	}
	var msgContent model.MsgContent
	err = json.Unmarshal(msg, &msgContent)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"Error": err,
		})
		return
	}
	log.Printf("got message: %v\n", msgContent)

	toUserName := msgContent.ToUsername
	msgContent.ToUsername = msgContent.FromUsername
	msgContent.FromUsername = toUserName
	replayJson, _ := json.Marshal(&msgContent)

	encryptMsg, cryptErr := wxcpt.EncryptMsg(string(replayJson), verifyTimestamp, verifyNonce)
	if cryptErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"Error": err,
		})
		return
	}
	c.Render(200, render.Data{
		Data: encryptMsg,
	})
}
