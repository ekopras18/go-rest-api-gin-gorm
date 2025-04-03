package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
)

var JwtKey []byte

type Config struct {
	DBHost        string
	DBUser        string
	DBPassword    string
	DBName        string
	DBPort        string
	DBSsl         string
	REDISHost     string
	REDISPassword string
	REDISDb       int
}

var AppConfig Config

func LoadConfig() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	redisDb, err := strconv.Atoi(os.Getenv("REDIS_DB"))
	if err != nil {
		log.Fatalf("Error converting REDIS_DB to int: %v", err)
	}

	AppConfig = Config{
		DBHost:        os.Getenv("DB_HOST"),
		DBUser:        os.Getenv("DB_USER"),
		DBPassword:    os.Getenv("DB_PASSWORD"),
		DBName:        os.Getenv("DB_NAME"),
		DBPort:        os.Getenv("DB_PORT"),
		DBSsl:         os.Getenv("DB_SSL"),
		REDISHost:     os.Getenv("REDIS_HOST"),
		REDISPassword: os.Getenv("REDIS_PASSWORD"),
		REDISDb:       redisDb,
	}

	JwtKey = []byte(os.Getenv("JWT_KEY"))
}
