package main

import (
	"github.com/gin-gonic/gin"
	"github.com/kekaifun/work-weixin-bot/handlers"
	"log"
	"net/http"
)

func main() {
	engine := gin.Default()
	bot := &handlers.Bot{}
	engine.GET("", bot.VerifyHandler)
	engine.POST("", bot.MessageHandler)
	engine.GET("/hello", bot.HelloHandler)

	err := http.ListenAndServe("0.0.0.0:8099", engine)
	if err != nil {
		log.Fatalf("start server failed,%v", err)
	}
}
