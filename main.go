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
		case "--host":
			config.Host = args[i]
		case "--color":
			config.Color = args[i]
		}
		argKey = ""
	}

	content := argKey
	// fmt.Printf("args: %v\n", args)
	// fmt.Printf("content: %s\n", content)

	if content == "" {
		fmt.Println("Welcome to ChatGpt CLI~ 😘 ")
		fmt.Println("(Enter '/exit' to quit; Enter '/new' to restart a new Session.)")
		fmt.Print("> ")
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
		if q == "/new" {
			gptMessages = []*chatgpt.GptMessage{}
			fmt.Println("-----Restart a new session-----")
			fmt.Printf("\n> ")
			continue
		} else if q == "/exit" {
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
	// font color for AI replies
	var colorAnsi = ""
	switch config.Color {
	case "black":
		colorAnsi = COLOR_BLACK_ANSI
	case "red":
		colorAnsi = COLOR_RED_ANSI
	case "green":
		colorAnsi = COLOR_GREEN_ANSI
	case "yellow":
		colorAnsi = COLOR_YELLOW_ANSI
	case "blue":
		colorAnsi = COLOR_BLUE_ANSI
	case "fuchsin":
		colorAnsi = COLOR_FUCHSIN_ANSI
	case "white":
		colorAnsi = COLOR_WHITE_ANSI
	}

	var answer string = ""

	// steam style
	if config.Stream {
		fmt.Print("\n")
		fmt.Print(colorAnsi)
		status, err := chatgpt.ChatGptStream(config, func(s *chatgpt.GptResponseStream) {
			fmt.Printf("%s", s.Answer())
			answer += s.Answer()
		}, gptMessages...)
		fmt.Print(COLOR_END_ANSI)
		if err != nil {
			fmt.Printf("Status: %d, Error: %s\n", status, err.Error())
			fmt.Printf("Press enter to retry...")
		}
		fmt.Printf("\n")
		return answer
	}

	// normal style
	gptRes, err := chatgpt.ChatGpt(config, gptMessages...)
	if err != nil {
		fmt.Printf("Error: %s\n", err.Error())
		fmt.Printf("Press enter to retry...")
	}
	answer = gptRes.Answer()
	fmt.Printf("\n%s\n", answer)
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

const (
	COLOR_BLACK_ANSI   = "\033[30m"
	COLOR_RED_ANSI     = "\033[31m"
	COLOR_GREEN_ANSI   = "\033[32m"
	COLOR_YELLOW_ANSI  = "\033[33m"
	COLOR_BLUE_ANSI    = "\033[34m"
	COLOR_FUCHSIN_ANSI = "\033[35m"
	COLOR_CYAN_ANSI    = "\033[36m"
	COLOR_WHITE_ANSI   = "\033[37m"
	COLOR_END_ANSI     = "\033[0m"
)
