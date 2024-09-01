package routes

import (
	"net/http"
	"os"

	"github.com/blackmagiqq/webproxy2/controllers"
	"github.com/blackmagiqq/webproxy2/services"
	"github.com/blackmagiqq/webproxy2/usecases"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	r.GET("/health/check", func(c *gin.Context) {
		c.String(http.StatusOK, "i`m alive")
	})

	r.POST("/api/v2/calculator/getServices", func(c *gin.Context) {
		apiService := services.APIService{}
		getServicesUseCase := usecases.CalculatorGetServicesUseCase{APIService: &apiService}
		calculatorController := controllers.CalculatorController{
			Host:               os.Getenv("CALCULATOR_HOST"),
			GetServicesUseCase: &getServicesUseCase,
		}
		calculatorController.GetServices(c)
	})
}
