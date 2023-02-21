package main

import (
	"github.com/gin-gonic/gin"
	"github.com/kekaifun/work-weixin-bot/handlers"
	"github.com/kekaifun/work-weixin-bot/handlers/chatgpt"
	"log"
	"net/http"
	"os"
)

func main() {
	engine := gin.Default()
	bot := &handlers.Bot{
		Name: "汪汪队长",
	}

	openAIHandler := &chatgpt.OpenAIHandler{
		CacheFile: map[string]*os.File{},
	}

	engine.GET("", bot.VerifyHandler)
	engine.POST("", bot.MessageHandler)
	engine.GET("/hello", bot.HelloHandler)
	engine.POST("/v1/chat/completions", openAIHandler.Completions)
	engine.OPTIONS("/v1/chat/completions", openAIHandler.Completions)

	err := http.ListenAndServe("0.0.0.0:8099", engine)
	if err != nil {
		log.Fatalf("start server failed,%v", err)
	}
}
