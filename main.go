package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/fengxxc/go-chatgpt/chatgpt"
)

const (
	CONFIG_FILE = "config.json"
)

func main() {

	// load config
	var config *chatgpt.GptConfig = &chatgpt.GptConfig{}
	if fileExists(CONFIG_FILE) {
		config = chatgpt.LoadConfig(CONFIG_FILE)
	}

	args := os.Args

	argKey := ""
	for i := 1; i < len(args); i++ {
		if argKey == "" {
			argKey = args[i]
			continue
		}
		switch argKey {
		case "--authorization":
			config.Authorization = args[i]
		case "--model":
			config.Model = args[i]
		case "--temperature":
			config.Temperature, _ = strconv.Atoi(args[i])
		case "--max_tokens":
			config.MaxTokens, _ = strconv.Atoi(args[i])
		case "--proxy":
			config.Proxy = args[i]
		}
		argKey = ""
	}

	content := argKey
	// fmt.Printf("args: %v\n", args)
	// fmt.Printf("content: %s\n", content)

	if content == "" {
		fmt.Println("Welcome to ChatGpt CLI~ 😘 ")
		fmt.Print("Start your show~ 😙 (Enter 'exit' to quit.)\n> ")
		chatOfSession(config)
	}

	chat(config, &chatgpt.GptMessage{Role: "user", Content: content})
}

func chatOfSession(config *chatgpt.GptConfig) {
	var gptMessages []*chatgpt.GptMessage
	scanner := bufio.NewScanner(os.Stdin)
	for {
		q := ""
		// fmt.Scanln(&q)
		scanner.Scan()
		q = scanner.Text()
		q = strings.TrimSpace(q)
		if q == "exit" {
			fmt.Println("See you next time~ 🤗")
			os.Exit(0)
			return
		}
		gptMessages = append(gptMessages, &chatgpt.GptMessage{Role: "user", Content: q})
		answer := chat(config, gptMessages...)
		fmt.Printf("\n> ")
		gptMessages = append(gptMessages, &chatgpt.GptMessage{Role: "assistant", Content: answer})
	}
}

func chat(config *chatgpt.GptConfig, gptMessages ...*chatgpt.GptMessage) string {
	var answer string = ""

	// steam style
	if config.Stream {
		err := chatgpt.ChatGptStream(config, func(s *chatgpt.GptResponseStream) {
			fmt.Printf("%s", s.Answer())
			answer += s.Answer()
		}, gptMessages...)
		if err != nil {
			panic(err)
		}
		fmt.Printf("\n")
		return answer
	}

	// normal style
	gptRes, err := chatgpt.ChatGpt(config, gptMessages...)
	if err != nil {
		panic(err)
	}
	answer = gptRes.Answer()
	fmt.Printf("%s\n", answer)
	return answer
	// fmt.Printf("(prompt_tokens: %d, prompt_tokens: %d, total_tokens: %d)\n",
	// 	gptRes.Usage.PromptTokens, gptRes.Usage.CompletionTokens, gptRes.Usage.TotalTokens)
}

func fileExists(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		return os.IsExist(err)
	}
	return true
}
