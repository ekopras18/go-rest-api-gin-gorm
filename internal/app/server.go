package app

import (
	"github.com/ekopras18/go-rest-api-gin-gorm/internal/database"
	"github.com/ekopras18/go-rest-api-gin-gorm/pkg/utils"
	"github.com/gin-gonic/gin"
	"log"
)

func RunServer() {
	database.InitDB()
	database.InitRedis()

	// Set Gin mode
	gin.SetMode(gin.ReleaseMode)
	// Create a new Gin router
	r := gin.New()
	utils.Logger(r)
	r.Use(gin.Recovery())

	// Set trusted proxies
	//trustedProxies := []string{
	//	"192.168.1.0/24",  // Example local network
	//	"10.0.0.0/8",      // Private network
	//	"172.16.0.0/12",   // Docker networks
	//}
	err := r.SetTrustedProxies(nil)
	if err != nil {
		log.Printf("Error setting trusted proxies: %v", err)
	}

	// Load Routes
	Routes(r)

	log.Println("Starting server on :8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
