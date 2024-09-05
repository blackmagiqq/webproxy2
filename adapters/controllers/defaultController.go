package controllers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type defaultController struct{}

func (ctrl *defaultController) validate(c *gin.Context, buff interface{}) bool {
	if err := c.ShouldBindJSON(buff); err != nil {
		ctrl.failResponse(c, http.StatusUnprocessableEntity, "Некорректный json")
		log.Println(err.Error())
		return true
	}
	return false
}

func (ctrl *defaultController) successResponse(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, data)
}

func (ctrl *defaultController) failResponse(c *gin.Context, status int, err string) {
	c.JSON(status, gin.H{"error": err})
}
