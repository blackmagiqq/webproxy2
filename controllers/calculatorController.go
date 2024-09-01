package controllers

import (
	"net/http"

	"github.com/blackmagiqq/webproxy2/dto"
	"github.com/blackmagiqq/webproxy2/usecases"
	"github.com/gin-gonic/gin"
)

type CalculatorController struct {
	Host               string
	GetServicesUseCase *usecases.CalculatorGetServicesUseCase
}

func (ctrl *CalculatorController) GetServices(c *gin.Context) {
	// парсим тело запроса и валидируем
	bodyDTO := new(dto.CalculatorBody)
	if err := c.BindJSON(bodyDTO); handleError(c, err) {
		return
	}

	headers := map[string]string{
		"X-User-Lang": ctrl.getUserLang(c),
	}

	// вызываем useCase
	services, err := ctrl.GetServicesUseCase.Handle(ctrl.Host, bodyDTO, headers)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// возвращаем результат
	c.JSON(http.StatusOK, services)
}

func (ctrl *CalculatorController) getUserLang(c *gin.Context) string {
	userLang := c.GetHeader("X-User-Lang")
	if userLang == "" {
		return "RUS"
	}
	return userLang
}
