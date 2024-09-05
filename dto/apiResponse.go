package dto

import "net/http"

type APIResponse struct {
	StatusCode int
	Body       []byte
	Headers    http.Header
}
