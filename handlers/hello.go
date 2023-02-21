package handlers

import (
	"github.com/gin-gonic/gin"
)

func (b *Bot) HelloHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "this is good",
	})
}
