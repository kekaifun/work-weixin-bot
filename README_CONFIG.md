# 配置文件使用说明

## 功能概述

现在支持通过 `--config` 命令行参数指定配置文件路径，程序会自动读取并解析YAML格式的配置文件。

## 使用方法

### 1. 使用默认配置文件
```bash
# 使用默认配置文件 config/config.yaml
./work-weixin-bot
```

### 2. 指定自定义配置文件
```bash
# 使用自定义配置文件
./work-weixin-bot --config /path/to/your/config.yaml
```

### 3. 查看帮助信息
```bash
./work-weixin-bot --help
```

## 配置文件格式

配置文件使用YAML格式，示例如下：

```yaml
# 服务监听地址
addr: 0.0.0.0:8099

# 企业微信机器人配置
wechat_bot:
  secret:
    token: your-token-here
    encoding_aes_key: your-aes-key-here
  serve_path: /wechat-bot
  backend_server:
    url: http://localhost:8099/wechat-bot
  async_response: true

# OpenAI代理配置（可选）
openai:
  base_url: https://api.openai.com
  serve_path: /openai
```

## 配置项说明

- `addr`: 服务器监听地址和端口
- `wechat_bot.secret.token`: 企业微信机器人的Token
- `wechat_bot.secret.encoding_aes_key`: 企业微信机器人的AES密钥
- `wechat_bot.serve_path`: 企业微信机器人的服务路径
- `wechat_bot.backend_server.url`: 后端服务器地址
- `wechat_bot.async_response`: 是否启用异步响应
- `openai.base_url`: OpenAI API基础URL（可选）
- `openai.serve_path`: OpenAI代理服务路径（可选）

## 错误处理

- 如果指定的配置文件不存在，程序会输出警告信息并使用默认配置继续运行
- 如果配置文件格式错误，程序会输出错误信息并使用默认配置继续运行
- 程序启动时会显示当前使用的配置文件路径和监听地址

## 示例

```bash
# 编译程序
go build -o bin/work-weixin-bot main.go

# 使用默认配置启动
./bin/work-weixin-bot

# 使用自定义配置启动
./bin/work-weixin-bot --config ./my-config.yaml
``` 