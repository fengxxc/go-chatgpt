package chatgpt

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

const (
	API_URL   string = "https://api.openai.com/v1/chat/completions"
	DEF_MODEL string = "gpt-3.5-turbo" // or "text-davinci-003"...
)

func ChatGpt(config *GptConfig, gptMessages ...*GptMessage) (GptResponse, error) {
	var model string = DEF_MODEL
	if config.Model != "" {
		model = config.Model
	}
	reqBody := &GptRequest{model, gptMessages}

	var httpClient *http.Client
	if config.Proxy != "" {
		httpTransport := &http.Transport{
			Proxy: func(_ *http.Request) (*url.URL, error) {
				return url.Parse(config.Proxy)
			},
		}
		httpClient = &http.Client{Transport: httpTransport}
	} else {
		httpClient = &http.Client{}
	}
	reqBodyStr, err := json.Marshal(reqBody)
	if err != nil {
		return GptResponse{}, err
	}
	// fmt.Println(string(reqBodyStr))
	req, err := http.NewRequest("POST", API_URL, bytes.NewReader(reqBodyStr))

	if err != nil {
		return GptResponse{}, err
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+config.Authorization)

	if err != nil {
		return GptResponse{}, err
	}

	resp, err := httpClient.Do(req)
	if resp != nil {
		defer resp.Body.Close()
	}

	if err != nil {
		return GptResponse{}, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		var gptErr *GptError
		json.NewDecoder(resp.Body).Decode(&gptErr)
		return GptResponse{}, fmt.Errorf(gptErr.Error.Message)
	}

	var gptRes *GptResponse
	json.NewDecoder(resp.Body).Decode(&gptRes)
	return *gptRes, nil
}
