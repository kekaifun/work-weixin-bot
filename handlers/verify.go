package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/render"
)

func (b *Bot) VerifyHandler(c *gin.Context) {
	echoStr, cryptErr := b.VerifySignature(c)
	if cryptErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"Error": "verify signature failed. " + cryptErr.Error(),
		})
		return
	}
	c.Render(200, render.Data{
		Data: echoStr,
	})
	return
}
