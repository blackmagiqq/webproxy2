package main

import (
	"github.com/blackmagiqq/webproxy2/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		panic("No .env file found")
	}
	r := gin.Default()
	routes.SetupRoutes(r)
	_ = r.Run()
}
