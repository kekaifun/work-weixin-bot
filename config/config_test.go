package config

import (
	"os"
	"testing"
)

func TestLoadConfig(t *testing.T) {
	// 创建临时配置文件
	tempConfig := `addr: 127.0.0.1:9000
wechat_bot:
  secret:
    token: test-token
    encoding_aes_key: test-key
  serve_path: /test-bot
  backend_server:
    url: http://localhost:9000/test-bot
  async_response: false
openai:
  base_url: https://test.openai.com
  serve_path: /test-openai
`

	// 写入临时文件
	tmpFile, err := os.CreateTemp("", "test-config-*.yaml")
	if err != nil {
		t.Fatalf("创建临时文件失败: %v", err)
	}
	defer os.Remove(tmpFile.Name())

	if _, err := tmpFile.WriteString(tempConfig); err != nil {
		t.Fatalf("写入临时文件失败: %v", err)
	}
	tmpFile.Close()

	// 测试加载配置
	cfg, err := LoadConfig(tmpFile.Name())
	if err != nil {
		t.Fatalf("加载配置失败: %v", err)
	}

	// 验证配置值
	if cfg.Addr != "127.0.0.1:9000" {
		t.Errorf("期望地址为 127.0.0.1:9000，实际为 %s", cfg.Addr)
	}

	if cfg.WechatBot.Secret.Token != "test-token" {
		t.Errorf("期望token为 test-token，实际为 %s", cfg.WechatBot.Secret.Token)
	}

	if cfg.OpenAI == nil {
		t.Error("OpenAI配置不应为nil")
	} else if cfg.OpenAI.BaseUrl != "https://test.openai.com" {
		t.Errorf("期望OpenAI BaseUrl为 https://test.openai.com，实际为 %s", cfg.OpenAI.BaseUrl)
	}
}

func TestLoadConfigFileNotExists(t *testing.T) {
	_, err := LoadConfig("not-exists.yaml")
	if err == nil {
		t.Error("期望加载不存在的文件时返回错误")
	}
}

func TestGetDefaultConfig(t *testing.T) {
	cfg := GetDefaultConfig()
	if cfg.Addr != "0.0.0.0:8099" {
		t.Errorf("期望默认地址为 0.0.0.0:8099，实际为 %s", cfg.Addr)
	}
}
