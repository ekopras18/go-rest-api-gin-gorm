package utils

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
	"os"
	"time"
)

type Claims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

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
