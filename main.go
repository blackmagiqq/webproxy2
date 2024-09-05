package main

import (
	"github.com/blackmagiqq/webproxy2/infrastructure/routes"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		panic("No .env file found")
	}

	r := routes.SetupRouter()
	_ = r.Run()
}
