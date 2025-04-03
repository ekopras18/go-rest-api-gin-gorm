package main

import (
	"rest-api-gin-gorm/config"
	"rest-api-gin-gorm/internal/app"
)

func main() {
	// Load configuration from config file or environment variables
	config.LoadConfig()
	// Run the server application
	app.RunServer()
}
