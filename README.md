# go-chatgpt

## Quick start
```
go-chatgpt --authorization <replace your chatgpt api key> "hello"
```
Your chatgpt api key should not have a "Bearer " prefix

## Config
> The first four are the same as the official chatgpt parameters, see: https://platform.openai.com/docs/guides/chat
 - `authorization` chatgpt api key (without 'Bearer ' as prefix)
 - `model` same as the official chatgpt parameter, default: 'gpt-3.5-turbo'
 - `temperature` same as the official chatgpt parameters
 - `max_tokens` same as the official chatgpt parameters
 - `proxy` proxy, supports Socks5 or HTTP
 
There are two ways to configure parameters
### via command parameters
Prefix with `--` in the startup command, like this
```
go-chatgpt --authorization <your chatgpt api key> --model gpt-3.5-turb "Hi~chatgpt~"
```
### via config.json
Run `cp config.json.template config.json`

Fill in the appropriate values in the `config.json` file

Then run command, like this
```
go-chatgpt "Hi~chatgpt~"
```

The priority is 'command parameters' > 'config.json'
## Session mode
Use in a complete conversational context until you actively exit (or exceed token limit)

If you don’t add what you want to ask to the start command, you enter session mode, such as
```
go-chatgpt --authorization <replace your chatgpt api key>
```
Then start talking to Chatgpt~

Enter `exit` or end the terminal to exit the session

## As a Package for go
```
go get github.com/fengxxc/go-chatgpt
```
import it where you need
```golang
import "github.com/fengxxc/go-chatgpt/chatgpt"
```
this is a complete example
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