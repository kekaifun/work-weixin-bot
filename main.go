package main

import (
	"flag"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/kekaifun/work-weixin-bot/config"
	"github.com/kekaifun/work-weixin-bot/handlers"
	"github.com/kekaifun/work-weixin-bot/handlers/chatgpt"
)

func main() {
	// 定义命令行参数
	var configPath string
	flag.StringVar(&configPath, "config", "config/config.yaml", "配置文件路径")
	flag.Parse()

	// 加载配置
	cfg, err := config.LoadConfig(configPath)
	if err != nil {
		log.Printf("加载配置文件失败: %v，使用默认配置", err)
		cfg = config.GetDefaultConfig()
	}

	log.Printf("使用配置文件: %s", configPath)
	log.Printf("服务监听地址: %s", cfg.Addr)

	engine := gin.Default()
	bot := &handlers.Bot{
		Name: "汪汪队长",
	}

	openAIHandler := &chatgpt.OpenAIHandler{
		CacheFile: map[string]*os.File{},
	}

	// 使用配置中的路径设置路由
	engine.GET(cfg.WechatBot.ServePath, bot.VerifyHandler)
	engine.POST(cfg.WechatBot.ServePath, bot.MessageHandler)
	engine.GET("/hello", bot.HelloHandler)

	// 如果配置了OpenAI代理，则添加相关路由
	if cfg.OpenAI != nil {
		log.Printf("OpenAI代理路径: %s", cfg.OpenAI.ServePath)
		engine.POST(cfg.OpenAI.ServePath, openAIHandler.Completions)
		engine.OPTIONS(cfg.OpenAI.ServePath, openAIHandler.Completions)
	}

	// 使用配置中的地址启动服务器
	err = http.ListenAndServe(cfg.Addr, engine)
	if err != nil {
		log.Fatalf("启动服务器失败: %v", err)
	}
}
