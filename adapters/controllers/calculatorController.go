package controllers

import (
	"log"
	"net/http"

	"github.com/blackmagiqq/webproxy2/adapters/presenters"
	"github.com/blackmagiqq/webproxy2/adapters/requests"
	"github.com/blackmagiqq/webproxy2/dto"
	"github.com/gin-gonic/gin"
)

type CalculatorGetServicesUseCase interface {
	Handle(
		headers map[string]string,
		body *dto.CalculatorGetServicesRequest,
	) (*dto.CalculatorGetServicesResponse, error)
}

type CalculatorController struct {
	defaultController
	GetServicesUseCase CalculatorGetServicesUseCase
}

// @BasePath /api/v2

// @Summary Тарифы калькулятора
// @Description Список доступных тарифов по заданным параметрам
// @Tags Калькулятор
// @Accept json
// @Param body body requests.CalculatorGetServices true "Запрос"
// @Produce json
// @Success 200 {object} presenters.CalculatorGetServices
// @Router /calculator/getServices [post]
func (ctrl *CalculatorController) GetServices(c *gin.Context) {
	// валидируем запрос и получаем данные в формате клиента
	requestData := new(requests.CalculatorGetServices)
	if validationIsFail := ctrl.validate(c, requestData); validationIsFail {
		return
	}

	inputDTO := ctrl.mapRequestToInputDTO(requestData)
	headers := ctrl.getHeaders(requestData)

	// вызываем useCase
	outputDTO, err := ctrl.GetServicesUseCase.Handle(headers, inputDTO)
	if err != nil {
		ctrl.failResponse(c, http.StatusInternalServerError, "Внутренняя ошибка сервера")
		log.Println(err.Error())
		return
	}

	// мапим DTO в формат клиента
	presenter := &presenters.CalculatorGetServices{}
	view := presenter.FromDTO(outputDTO)

	// возвращаем результат
	ctrl.successResponse(c, view)
}

func (ctrl *CalculatorController) getHeaders(request *requests.CalculatorGetServices) map[string]string {
	headers := make(map[string]string)
	userLang := request.Language
	if userLang == "" {
		headers["X-User-Lang"] = "RUS"
	} else {
		headers["X-User-Lang"] = userLang
	}
	return headers
}

func (ctrl *CalculatorController) mapRequestToInputDTO(
	request *requests.CalculatorGetServices,
) *dto.CalculatorGetServicesRequest {
	DTO := &dto.CalculatorGetServicesRequest{
		Sender: dto.CalculatorGetServicesSender{
			CityID:         request.Sender.CityID,
			ContragentID:   request.Sender.ContragentID,
			ContragentType: request.Sender.ContragentType,
			ContractID:     request.Sender.ContractID,
		},
		Receiver: dto.CalculatorGetServicesReceiver{
			CityID:         request.Receiver.CityID,
			ContragentID:   request.Receiver.ContragentID,
			ContragentType: request.Receiver.ContragentType,
			ContractID:     request.Receiver.ContractID,
		},
		Payer: dto.CalculatorGetServicesPayer{
			ContragentID: request.Payer.ContragentID,
			ContractID:   request.Payer.ContractID,
			PayerType:    request.Payer.PayerType,
		},
		OrderParam: dto.CalculatorGetServicesOrderParam{
			OrderTypeCode:            request.OrderParam.OrderTypeCode,
			AdditionalOrderTypeCodes: request.OrderParam.AdditionalOrderTypeCodes,
			CashOnDeliveryIndividual: request.OrderParam.CashOnDeliveryIndividual,
			IsClientReturn:           request.OrderParam.IsClientReturn,
			HaveFirstOrderForReturn:  request.OrderParam.HaveFirstOrderForReturn,
			SellerUUID:               request.OrderParam.SellerUUID,
			OrderCost:                request.OrderParam.OrderCost,
			CalcMode:                 request.OrderParam.CalcMode,
		},
		InterfaceCode: request.InterfaceCode,
		CurrencyMark:  request.CurrencyMark,
		CalcDate:      request.CalcDate,
		Filters: dto.CalculatorGetServicesFilters{
			CalcServicesID: request.Filters.CalcServicesID,
			CalcModes:      request.Filters.CalcModes,
		},
	}
	for _, pkg := range request.Packages {
		DTO.Packages = append(DTO.Packages, dto.CalculatorGetServicesPackage{
			Length: pkg.Length,
			Width:  pkg.Width,
			Height: pkg.Height,
			Weight: pkg.Weight,
		})
	}
	return DTO
}
