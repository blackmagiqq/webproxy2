package usecases

import (
	"github.com/blackmagiqq/webproxy2/services"
)

const uri = "/api/calculator/getServices"

type CalculatorGetServicesUseCase struct {
	APIService *services.APIService
}

func (useCase *CalculatorGetServicesUseCase) Handle(
	host string, body interface{},
	headers map[string]string,
) (interface{}, error) {
	services, err := useCase.APIService.Handle(
		host+uri,
		"POST",
		headers,
		body,
	)
	if err != nil {
		return nil, err
	}
	return services, nil
}
