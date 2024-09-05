package services

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/blackmagiqq/webproxy2/dto"
)

// ApiService представляет сервис для работы с API.
type APIService struct{}

// handle обрабатывает запросы в зависимости от метода.
func (s *APIService) Handle(
	url string,
	method string,
	headers map[string]string,
	body interface{},
) (*dto.APIResponse, error) {
	switch method {
	case "GET":
		return s.get(url, headers)
	case "POST":
		return s.post(url, headers, body)
	default:
		return nil, errors.New("unsupported method")
	}
}

// get отправляет GET-запрос на указанный URL.
func (s *APIService) get(url string, headers map[string]string) (*dto.APIResponse, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	for key, value := range headers {
		req.Header.Set(key, value)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return s.processResponse(resp)
}

// post отправляет POST-запрос на указанный URL с телом запроса.
func (s *APIService) post(url string, headers map[string]string, body interface{}) (*dto.APIResponse, error) {
	jsonData, err := json.Marshal(body)
	if err != nil {
		return nil, fmt.Errorf("create JSON at post method: %w", err)
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, fmt.Errorf("create new http request at post method: %w", err)
	}

	for key, value := range headers {
		req.Header.Set(key, value)
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("send post request: %w", err)
	}
	defer resp.Body.Close()

	return s.processResponse(resp)
}

// processResponse обрабатывает HTTP-ответ и возвращает структуру Response.
func (s *APIService) processResponse(resp *http.Response) (*dto.APIResponse, error) {
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("read body from response: %w", err)
	}

	if resp.StatusCode >= http.StatusBadRequest {
		return nil, errors.New("proxy error: " + string(body))
	}

	return &dto.APIResponse{
		StatusCode: resp.StatusCode,
		Body:       body,
		Headers:    resp.Header,
	}, nil
}
