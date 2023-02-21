package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/render"
	"github.com/kekaifun/work-weixin-bot/config"
	"github.com/kekaifun/work-weixin-bot/wxcrypt"
	"net/http"
)

func Verify(c *gin.Context) {

	verifyTimestamp := c.Query("timestamp")
	verifyNonce := c.Query("nonce")
	verifyEchoStr := c.Query("echostr")
	verifyMsgSign := c.Query("msg_signature")

	wxcpt := wxcrypt.NewWXBizMsgCrypt(config.Token, config.EncodingAESKey, "", wxcrypt.JsonType)

	echoStr, cryptErr := wxcpt.VerifyURL(verifyMsgSign, verifyTimestamp, verifyNonce, verifyEchoStr)
	if cryptErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"Error": "verify signature failed",
		})
	}
	c.Render(200, render.Data{
		Data: echoStr,
	})
	return

	//paras := []string{config.Token, timestamp, nonce, echostr}
	//sort.Strings(paras)
	//result := strings.Join(paras, "")
	//devSignature := sha1.Sum([]byte(result))
	//msgSignature := c.Query("msg_signature")
	//
	//if string(devSignature[:]) != msgSignature {
	//	c.JSON(http.StatusInternalServerError, gin.H{
	//		"Error": "bad signature",
	//	})
	//	return
	//}
	//
	//ciphertext, err := base64.StdEncoding.DecodeString(echostr)
	//if err != nil {
	//	c.JSON(http.StatusInternalServerError, gin.H{
	//		"Error": "decode echostr failed",
	//	})
	//	return
	//}
	//aesKey, _ := base64.StdEncoding.DecodeString(config.EncodingAESKey)
	//blc, err := aes.NewCipher(aesKey)
	//if err != nil {
	//	c.JSON(http.StatusInternalServerError, gin.H{
	//		"Error": "create aes cipher failed",
	//	})
	//	return
	//}
	//iv := ciphertext[:blc.BlockSize()]
	//ciphertext = ciphertext[blc.BlockSize():]
	//mode := cipher.NewCBCDecrypter(blc, iv)
	//mode.CryptBlocks(ciphertext, ciphertext)
	//fmt.Println(string(ciphertext))
	//
	//content := ciphertext[16:]                                   //去掉前16随机字节
	//msg_len, _ := strconv.ParseInt(string(content[0:4]), 10, 32) // 取出4字节的msg_len
	//
	//msg := content[4 : msg_len+4]    // 截取msg_len 长度的msg
	//receiveid := content[msg_len+4:] //= "wx5823bf96d3bd56c7" # 剩余字节为receiveid
	//fmt.Printf("msg: %s, receiveid: %s", msg, receiveid)
	//c.Render(200, render.Data{
	//	Data: msg,
	//})
}
