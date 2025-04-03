package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

var JwtKey []byte

type Config struct {
	DBHost     string
	DBUser     string
	DBPassword string
	DBName     string
	DBPort     string
	DBSsl      string
}

var AppConfig Config

func LoadConfig() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	AppConfig = Config{
		DBHost:     os.Getenv("DB_HOST"),
		DBUser:     os.Getenv("DB_USER"),
		DBPassword: os.Getenv("DB_PASSWORD"),
		DBName:     os.Getenv("DB_NAME"),
		DBPort:     os.Getenv("DB_PORT"),
		DBSsl:      os.Getenv("DB_SSL"),
	}

	JwtKey = []byte(os.Getenv("JWT_KEY"))
}
