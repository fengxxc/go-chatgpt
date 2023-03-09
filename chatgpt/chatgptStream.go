package chatgpt

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

func ChatGptStream(config *GptConfig, callback func(*GptResponseStream), gptMessages ...*GptMessage) error {
	req, err := GetRequest(config, gptMessages...)
	if err != nil {
		return err
	}
	var httpClient *http.Client = GetHttpClient(req, config.Proxy)
	resp, err := httpClient.Do(req)
	if resp != nil {
		defer resp.Body.Close()
	}
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		var gptErr *GptError
		json.NewDecoder(resp.Body).Decode(&gptErr)
		return fmt.Errorf(gptErr.Error.Message)
	}
	for {
		event, err := readEvent(resp.Body)
		if err != nil {
			fmt.Println(err)
			break
		}
		if strings.TrimSpace(event.Data) == "[DONE]" {
			break
		}
		// fmt.Printf("Event Name: %s\n", event.Name)
		// fmt.Printf("Event Data: %s\n", event.Data)
		var gptRes *GptResponseStream
		err = json.NewDecoder(strings.NewReader(event.Data)).Decode(&gptRes)
		if err != nil {
			return err
		}
		callback(gptRes)
	}
	return nil
}

func readEvent(body io.ReadCloser) (*event, error) {
	e := &event{}
	line, err := readLine(body)
	if err != nil {
		return nil, err
	}

	switch {
	case line == "":
		return nil, fmt.Errorf("empty line")
	case line[0] == ':':
		return nil, nil
	case len(line) >= 6 && line[:6] == "event:":
		e.Name = line[6:]
	case len(line) >= 5 && line[:5] == "data:":
		e.Data = line[5:]
	default:
		return nil, fmt.Errorf("invalid line: %s", line)
	}

	for {
		line, err := readLine(body)
		if err != nil {
			return nil, err
		}

		switch {
		case line == "":
			return e, nil
		case line[0] == ':':
			continue
		case len(line) >= 6 && line[:6] == "event:":
			return nil, fmt.Errorf("event name changed")
		case len(line) >= 5 && line[:5] == "data:":
			e.Data += "\n" + line[5:]
		default:
			return nil, fmt.Errorf("invalid line: %s", line)
		}
	}

}

type event struct {
	Name string
	Data string
}

func readLine(body io.ReadCloser) (string, error) {
	lineBuf := make([]byte, 0, 1024)
	for {
		b := make([]byte, 1)
		if _, err := body.Read(b); err != nil {
			return "", err
		}
		if b[0] == '\r' {
			continue
		}
		if b[0] == '\n' {
			break
		}
		lineBuf = append(lineBuf, b[0])
	}
	return string(lineBuf), nil
}
