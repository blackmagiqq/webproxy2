package usecases

import (
	"encoding/json"
	"errors"
	"testing"

	"github.com/blackmagiqq/webproxy2/dto"
	"github.com/blackmagiqq/webproxy2/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type CalculatorGetServicesUseCaseSuite struct {
	suite.Suite
	requestBody dto.CalculatorGetServicesRequest
	mockService *mocks.MockAPIService
	useCase     *CalculatorGetServicesUseCase
}

func (s *CalculatorGetServicesUseCaseSuite) SetupTest() {
	s.requestBody = dto.CalculatorGetServicesRequest{}
	s.mockService = mocks.NewMockAPIService(s.T())
	s.useCase = NewCalculatorGetServicesUseCase("http://test.com", s.mockService)
}

func (s *CalculatorGetServicesUseCaseSuite) TestSuccessCallService() {
	response := dto.APIResponse{}
	responseBody := struct{ Message string }{Message: "success"}
	responseBodyJSON, _ := json.Marshal(responseBody)
	response.Body = responseBodyJSON

	s.mockService.EXPECT().Handle(s.useCase.url, s.useCase.method, map[string]string{}, &s.requestBody).Return(&response, nil)

	_, err := s.useCase.Handle(map[string]string{}, &s.requestBody)
	if err != nil {
		s.T().Fatalf("expected no error, got %v", err)
	}
	s.mockService.AssertExpectations(s.T())
}

func (s *CalculatorGetServicesUseCaseSuite) TestSuccessHandleServiceError() {
	s.mockService.On(
		"Handle",
		s.useCase.url,
		s.useCase.method,
		map[string]string{},
		&s.requestBody,
	).Return(nil, errors.New("test error"))

	_, err := s.useCase.Handle(map[string]string{}, &s.requestBody)
	assert.ErrorContains(s.T(), err, "calculator getServices useCase API service")
}

func (s *CalculatorGetServicesUseCaseSuite) TestSuccessHandleJSONParsingError() {
	s.mockService.On(
		"Handle",
		s.useCase.url,
		s.useCase.method,
		map[string]string{},
		&s.requestBody,
	).Return(&dto.APIResponse{}, nil)

	_, err := s.useCase.Handle(map[string]string{}, &s.requestBody)
	assert.ErrorContains(s.T(), err, "calculator getServices useCase unparse response")
}

func TestCalculatorGetServicesUseCaseSuite(t *testing.T) {
	suite.Run(t, new(CalculatorGetServicesUseCaseSuite))
}
