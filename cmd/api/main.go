package main

import (
	"go-rest-api-gin-gorm/config"
	"go-rest-api-gin-gorm/internal/app"
)

func main() {
	// Load configuration from config file or environment variables
	config.LoadConfig()
	// Run the server application
	app.RunServer()
}
