package main

import (
	"github.com/ekopras18/go-rest-api-gin-gorm/config"
	"github.com/ekopras18/go-rest-api-gin-gorm/internal/app"
)

// @title REST API with Gin and Gorm
// @version 0.1.5
// @description This is a REST API using Gin and Gorm.
// @schemes http https

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization

func main() {
	// Load configuration from config file or environment variables
	config.LoadConfig()
	// Run the server application
	app.RunServer()
}
