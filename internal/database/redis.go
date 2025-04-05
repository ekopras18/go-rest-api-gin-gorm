package database

import (
	"context"
	"github.com/ekopras18/go-rest-api-gin-gorm/config"
	"github.com/redis/go-redis/v9"
	"log"
	"time"
)

var RedisClient *redis.Client

func InitRedis() {

	RedisClient = redis.NewClient(&redis.Options{
		Addr:     config.AppConfig.REDISHost,
		Password: config.AppConfig.REDISPassword,
		DB:       config.AppConfig.REDISDb,
	})

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := RedisClient.Ping(ctx).Result()
	if err != nil {
		log.Fatalf("Gagal terhubung ke Redis: %v", err)
	}
}
