# go-chatgpt

## 快速开始
```
go-chatgpt --authorization <replace your chatgpt api key> "hello"
```
注意，你的chatgpt api key不要有"Bearer "前缀

## 配置
> 前四个与chatgpt官方参数相同，详见：https://platform.openai.com/docs/guides/chat
 - `authorization` chatgpt api key (without 'Bearer ' as prefix)
 - `model` chatgpt模型，默认为“gpt-3.5-turbo”
 - `temperature` 生成文本的多样性， 0~1
 - `max_tokens` 生成文本时最多可以使用的token数
 - `proxy` 代理，支持socks5 或 http
 
参数的配置有以下两种方式
### 通过 命令参数
在启动命令中以`--`前缀，例如
```
go-chatgpt --authorization <your chatgpt api key> --model gpt-3.5-turb "Hi~chatgpt~"
```
### 通过 配置文件config.json
执行`cp config.json.template config.json`，在`config.json`文件中填写相应的值

然后直接启动，例如
```
go-chatgpt "Hi~chatgpt~"
```

优先级为 命令参数 > 配置文件
## Session模式
在完整的会话上下文中使用，直到你主动退出（或超过token限制）

启动命令不加你要问的话，即进入Session模式，如
```
go-chatgpt --authorization <replace your chatgpt api key>
```
然后就开始跟chatgpt聊天吧~，输入`exit`或结束终端即退出会话

## 作为Package
```
go get -u github.com/fengxxc/go-chatgpt
```
在你需要的地方import
```golang
import "github.com/fengxxc/go-chatgpt/chatgpt"
```
这是个完整的例子
```golang
package main

import (
	"fmt"

	"github.com/fengxxc/go-chatgpt/chatgpt"
)

func main() {
	config := &chatgpt.GptConfig{
		Authorization: "<replace your key>",
		Proxy:         "socks5://127.0.0.1:4698",
	}
	message0 := &chatgpt.GptMessage{
		Role:    "system",
		Content: "You are a helpful assistant.",
	}
	message1 := &chatgpt.GptMessage{
		Role:    "user",
		Content: "What is the World Cup 2022 winner?",
	}
	response, err := chatgpt.ChatGpt(config, message0, message1)
	if err != nil {
		fmt.Printf("something went wrong: %v", err)
	}
	for _, grc := range response.Choices {
		fmt.Printf("chatgpt: %v\n", grc.Message.Content)
	}
	// emm... chatgpt don't know yet ¯\_(ツ)_/¯
}
```