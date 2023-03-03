package chatgpt

import "strings"

type GptResMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type GptResChoices struct {
	Index        int           `json:"index"`
	FinishReason string        `json:"finish_reason"`
	Message      GptResMessage `json:"message"`
}

type GptResUsage struct {
	PromptTokens     int `json:"prompt_tokens"`
	CompletionTokens int `json:"completion_tokens"`
	TotalTokens      int `json:"total_tokens"`
}

type GptResponse struct {
	Id      string          `json:"id"`
	Object  string          `json:"object"`
	Created int32           `json:"created"`
	Model   string          `json:"model"`
	Choices []GptResChoices `json:"choices"`
	Usage   GptResUsage     `josn:"usage"`
}

func (r *GptResponse) Answer() string {
	var arr []string = make([]string, len(r.Choices))
	for i := 0; i < len(r.Choices); i++ {
		arr[r.Choices[i].Index] = r.Choices[i].Message.Content
	}
	return strings.Join(arr, "\n")
}

type GptRequest struct {
	Model    string        `json:"model"`
	Messages []*GptMessage `json:"messages"`
}

type GptMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type GptError struct {
	Error GptErrorDetails `json:"error"`
}

type GptErrorDetails struct {
	Message string `json:"message"`
	Type    string `json:"type"`
	Param   string `json:"param"`
	Code    int    `json:"code"`
}

// func (e *GptError) Error() string {
// 	return fmt.Sprintf("%s\n", e.Error.Message)
// }

type GptConfig struct {
	Authorization string `json:"authorization"`
	Model         string `json:"model"`
	Temperature   int    `json:"temperature,omitempty"`
	MaxTokens     int    `json:"max_tokens,omitempty"`
	// 'Proxy' support socks5 and http
	Proxy string `json:"proxy"`
}
