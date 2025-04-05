package utils

import (
	"context"
	"fmt"
	"github.com/ekopras18/go-rest-api-gin-gorm/internal/database"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
	"os"
	"time"
)

func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

func CheckPasswordHash(password, hash string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password)) == nil
}

type ResponseMessage200 struct {
	Status  int    `json:"status" example:"200"`
	Message string `json:"message" example:"success"`
}

type ResponseMessage201 struct {
	Status  int    `json:"status" example:"201"`
	Message string `json:"message" example:"success"`
}

type ResponseMessage400 struct {
	Status  int    `json:"status" example:"400"`
	Message string `json:"message" example:"Bad Request"`
}

type ResponseMessage404 struct {
	Status  int    `json:"status" example:"404"`
	Message string `json:"message" example:"not found"`
}

type ResponseMessage409 struct {
	Status  int    `json:"status" example:"409"`
	Message string `json:"message" example:"Conflict"`
}

func Response(c *gin.Context, statusCode int, message string, data interface{}) {
	response := gin.H{
		//"status":  statusCode == 200 || statusCode == 201,
		"code":    statusCode,
		"message": message,
	}

	if data != nil {
		response["data"] = data
	}

	c.JSON(statusCode, response)
}

func CustomErrorHandler(c *gin.Context) {
	c.Next()
	println(len(c.Errors))
	if len(c.Errors) > 0 {
		Response(c, http.StatusInternalServerError, "Internal Server Error", nil)
	}
}

func NotFoundHandler(c *gin.Context) {
	Response(c, http.StatusNotFound, "Resource not found", nil)
}

func Logger(r *gin.Engine) {
	f, err := os.Create("server.log")
	if err != nil {
		log.Fatal(err)
	}
	gin.DefaultWriter = f

	r.Use(gin.LoggerWithConfig(gin.LoggerConfig{
		Output:    f,
		SkipPaths: []string{"/ping"},
		Formatter: func(param gin.LogFormatterParams) string {
			// Format custom log
			return fmt.Sprintf("[%s] \"%s %s %s\" %d %s \"%s\"\n",
				param.TimeStamp.Format(time.RFC3339),
				param.Method,
				param.Path,
				param.Request.Proto,
				param.StatusCode,
				param.Latency,
				param.ClientIP,
			)
		},
	}))
}

func BlacklistToken(token string) {
	ctx := context.Background()
	err := database.RedisClient.Set(ctx, "blacklist:"+token, true, 24*time.Hour).Err()
	if err != nil {
		return
	}
}

func IsTokenBlacklisted(token string) bool {
	ctx := context.Background()
	val, err := database.RedisClient.Get(ctx, "blacklist:"+token).Result()
	return err == nil && val == "1"
}
