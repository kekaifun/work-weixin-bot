package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	// Addr is the address where the bot listens
	Addr      string          `yaml:"addr"`
	WechatBot WechatBotConfig `yaml:"wechat_bot"`
	// OpenAI is the configuration for the OpenAI proxy, if not set, the OpenAI proxy will not be used
	OpenAI *OpenAIProxyConfig `yaml:"openai,omitempty"`
}

type WechatBotConfig struct {
	// Secret is the secret to decrypt the message from WeChat Work
	Secret WechatBotSecret `yaml:"secret"`
	// ServePath is the path where the WeChat Work bot receives messages
	ServePath string `yaml:"serve_path"`
	// BackendServer is the backend server address for the WeChat Work bot, used to handle messages sent from WeChat Work
	BackendServer WechatBotBackendServer `yaml:"backend_server"`
	// AsyncResponse is the flag to enable asynchronous response, if true, the bot will respond asynchronously
	AsyncResponse bool `yaml:"async_response"`
}

type WechatBotSecret struct {
	Token          string `yaml:"token"`
	EncodingAESKey string `yaml:"encoding_aes_key"`
}

// WechatBotBackendServer configures the backend server address for the WeChat Work bot, used to handle messages sent from WeChat Work
type WechatBotBackendServer struct {
	Url string `yaml:"url"`
}

type OpenAIProxyConfig struct {
	BaseUrl string `yaml:"base_url"`
	// ServePath is the path where the OpenAI proxy receives messages
	ServePath string `yaml:"serve_path"`
}

// LoadConfig 从指定的配置文件路径加载配置
func LoadConfig(configPath string) (*Config, error) {
	// 检查文件是否存在
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		return nil, fmt.Errorf("配置文件不存在: %s", configPath)
	}

	// 读取配置文件
	data, err := os.ReadFile(configPath)
	if err != nil {
		return nil, fmt.Errorf("读取配置文件失败: %v", err)
	}

	// 解析yaml配置
	var config Config
	if err := yaml.Unmarshal(data, &config); err != nil {
		return nil, fmt.Errorf("解析配置文件失败: %v", err)
	}

	return &config, nil
}

// GetDefaultConfig 返回默认配置
func GetDefaultConfig() *Config {
	return &Config{
		Addr: "0.0.0.0:8099",
		WechatBot: WechatBotConfig{
			Secret: WechatBotSecret{
				Token:          "",
				EncodingAESKey: "",
			},
			ServePath: "/wechat-bot",
			BackendServer: WechatBotBackendServer{
				Url: "http://localhost:8099/wechat-bot",
			},
			AsyncResponse: true,
		},
	}
}
