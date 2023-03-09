package chatgpt

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"
)

const (
	API_URL   string = "https://api.openai.com/v1/chat/completions"
	DEF_MODEL string = "gpt-3.5-turbo" // or "text-davinci-003"...
)

func ChatGpt(config *GptConfig, gptMessages ...*GptMessage) (GptResponse, error) {
	req, err := GetRequest(config, gptMessages...)
	if err != nil {
		return GptResponse{}, err
	}
	var httpClient *http.Client = GetHttpClient(req, config.Proxy)
	resp, err := httpClient.Do(req)
	if resp != nil {
		defer resp.Body.Close()
	}
	if err != nil {
		return GptResponse{}, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		var gptErr *GptError
		json.NewDecoder(resp.Body).Decode(&gptErr)
		return GptResponse{}, fmt.Errorf(gptErr.Error.Message)
	}

	var gptRes *GptResponse
	err = json.NewDecoder(resp.Body).Decode(&gptRes)
	if err != nil {
		return GptResponse{}, err
	}
	return *gptRes, nil
}

func GetHttpClient(req *http.Request, proxy string) *http.Client {
	var httpClient *http.Client
	if proxy != "" {
		httpTransport := &http.Transport{
			Proxy: func(_ *http.Request) (*url.URL, error) {
				return url.Parse(proxy)
			},
		}
		httpClient = &http.Client{Transport: httpTransport}
	} else {
		httpClient = &http.Client{}
	}
	return httpClient
}

func GetRequest(config *GptConfig, gptMessages ...*GptMessage) (*http.Request, error) {
	var model string = DEF_MODEL
	if config.Model != "" {
		model = config.Model
	}
	reqBody := &GptRequest{model, gptMessages, config.Stream}
	reqBodyStr, err := json.Marshal(reqBody)
	if err != nil {
		return nil, err
	}
	// fmt.Println(string(reqBodyStr))
	req, err := http.NewRequest("POST", API_URL, bytes.NewReader(reqBodyStr))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+config.Authorization)
	if config.Stream {
		req.Header.Set("Accept", "text/event-stream")
	}
	return req, err
}

func LoadConfig(path string) *GptConfig {
	data, err := os.ReadFile(path)
	if err != nil {
		panic(fmt.Errorf("load config file failed: %v", err))
	}
	var config *GptConfig = &GptConfig{}
	err = json.Unmarshal(data, config)
	if err != nil {
		panic(fmt.Errorf("decode config file failed: %v", err))
	}
	return config
}
