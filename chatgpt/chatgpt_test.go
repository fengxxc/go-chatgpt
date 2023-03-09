package chatgpt

import (
	"fmt"
	"testing"
)

func TestChatgpt(t *testing.T) {
	config := LoadConfig("../config.json")
	message0 := &GptMessage{
		Role:    "system",
		Content: "You are a helpful assistant.",
	}
	message1 := &GptMessage{
		Role:    "user",
		Content: "What is the World Cup 2022 winner?",
	}
	response, err := ChatGpt(config, message0, message1)
	if err != nil {
		fmt.Printf("something went wrong: %v", err)
	}
	for _, grc := range response.Choices {
		fmt.Printf("chatgpt: %v\n", grc.Message.Content)
	}
	// emm... chatgpt don't know yet ¯\_(ツ)_/¯
}

func TestChatgptStream(t *testing.T) {
	config := LoadConfig("../config.json")
	config.Stream = true
	message0 := &GptMessage{
		Role:    "system",
		Content: "You are a helpful assistant.",
	}
	message1 := &GptMessage{
		Role:    "user",
		Content: "What is the World Cup 2022 winner?",
	}
	fmt.Printf("chatgpt: \n")
	err := ChatGptStream(config, func(gr *GptResponseStream) {
		fmt.Printf("%s", gr.Answer())
	}, message0, message1)
	if err != nil {
		fmt.Printf("something went wrong: %v", err)
	}
	// emm... chatgpt don't know yet ¯\_(ツ)_/¯
}
