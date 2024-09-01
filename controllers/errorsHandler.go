package controllers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// handleError обработка ошибок
// максимально простая реализация.
// не раскрываем клиенту никакие детали.
func handleError(c *gin.Context, err error) bool {
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return true
	}
	return false
}
