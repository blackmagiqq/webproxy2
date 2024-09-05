package usecases

import (
	"encoding/json"
	"fmt"

	"github.com/blackmagiqq/webproxy2/dto"
)

type APIService interface {
	Handle(url string, method string, headers map[string]string, body interface{}) (*dto.APIResponse, error)
}

func NewCalculatorGetServicesUseCase(host string, apiService APIService) *CalculatorGetServicesUseCase {
	return &CalculatorGetServicesUseCase{
		url:        host + "/api/calculator/getServices",
		method:     "POST",
		apiService: apiService,
	}
}

type CalculatorGetServicesUseCase struct {
	url        string
	method     string
	apiService APIService
}

func (u *CalculatorGetServicesUseCase) Handle(
	headers map[string]string,
	body *dto.CalculatorGetServicesRequest,
) (*dto.CalculatorGetServicesResponse, error) {
	response, err := u.apiService.Handle(
		u.url,
		u.method, headers,
		body,
	)
	if err != nil {
		return nil, fmt.Errorf("calculator getServices useCase API service: %w", err)
	}

	responseDTO := new(dto.CalculatorGetServicesResponse)
	if err := json.Unmarshal(response.Body, responseDTO); err != nil {
		return nil, fmt.Errorf("calculator getServices useCase unparse response: %w", err)
	}
	return responseDTO, nil
}
