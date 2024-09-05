package routes

import (
	"net/http"
	"os"

	"github.com/blackmagiqq/webproxy2/adapters/controllers"
	"github.com/blackmagiqq/webproxy2/adapters/services"
	"github.com/blackmagiqq/webproxy2/usecases"
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	swaggerfiles "github.com/swaggo/files"
	docs "github.com/blackmagiqq/webproxy2/docs"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/health/check", func(c *gin.Context) {
		c.String(http.StatusOK, "i`m alive")
	})

	r.POST("/api/v2/calculator/getServices", func(c *gin.Context) {
		apiService := services.APIService{}
		getServicesUseCase := usecases.NewCalculatorGetServicesUseCase(os.Getenv("CALCULATOR_HOST"), &apiService)
		calculatorController := controllers.CalculatorController{
			GetServicesUseCase: getServicesUseCase,
		}
		calculatorController.GetServices(c)
	})

	docs.SwaggerInfo.BasePath = "/api/v2"
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	return r
}