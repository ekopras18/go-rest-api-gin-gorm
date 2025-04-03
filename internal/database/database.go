package database

import (
	"fmt"
	"go-rest-api-gin-gorm/config"
	"go-rest-api-gin-gorm/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var Db *gorm.DB

func InitDB() {
	// Load configuration
	config.LoadConfig()

	// Build DSN from config
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		config.AppConfig.DBHost,
		config.AppConfig.DBUser,
		config.AppConfig.DBPassword,
		config.AppConfig.DBName,
		config.AppConfig.DBPort,
		config.AppConfig.DBSsl,
	)

	var err error

	// Connect to the database and assign it to Db
	Db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Run migrations
	if err := Db.AutoMigrate(&models.User{}); err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

}
