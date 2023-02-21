package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/kekaifun/work-weixin-bot/config"
	"github.com/sbzhu/weworkapi_golang/wxbizmsgcrypt"
	"log"
)

type Bot struct {
}

func (b *Bot) VerifySignature(c *gin.Context) ([]byte, error) {
	verifyTimestamp := c.Query("timestamp")
	verifyNonce := c.Query("nonce")
	verifyEchoStr := c.Query("echostr")
	verifyMsgSign := c.Query("msg_signature")

	wxcpt := wxbizmsgcrypt.NewWXBizMsgCrypt(config.Token, config.EncodingAESKey, "", wxbizmsgcrypt.XmlType)

	echoStr, cryptErr := wxcpt.VerifyURL(verifyMsgSign, verifyTimestamp, verifyNonce, verifyEchoStr)
	var err error
	if cryptErr != nil {
		err = fmt.Errorf("code:%d, msg:%s", cryptErr.ErrCode, cryptErr.ErrMsg)
	}
	log.Printf("verify,echoStr: %s", string(echoStr))
	return echoStr, err
}
