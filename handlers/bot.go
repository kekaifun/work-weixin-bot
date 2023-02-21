package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/kekaifun/work-weixin-bot/config"
	"github.com/kekaifun/work-weixin-bot/wxcrypt"
	"log"
)

type Bot struct {
}

func (b *Bot) VerifySignature(c *gin.Context) ([]byte, error) {
	verifyTimestamp := c.Query("timestamp")
	verifyNonce := c.Query("nonce")
	verifyEchoStr := c.Query("echostr")
	verifyMsgSign := c.Query("msg_signature")

	wxcpt := wxcrypt.NewWXBizMsgCrypt(config.Token, config.EncodingAESKey, "", wxcrypt.JsonType)

	echoStr, cryptErr := wxcpt.VerifyURL(verifyMsgSign, verifyTimestamp, verifyNonce, verifyEchoStr)
	log.Printf("verify,echoStr: %s", string(echoStr))
	return echoStr, cryptErr
}
