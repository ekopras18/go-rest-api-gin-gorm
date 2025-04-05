package app

import (
	"github.com/ekopras18/go-rest-api-gin-gorm/config"
	"github.com/ekopras18/go-rest-api-gin-gorm/internal/dto/auth"
	"github.com/ekopras18/go-rest-api-gin-gorm/pkg/utils"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/ulule/limiter/v3"
	ginlimiter "github.com/ulule/limiter/v3/drivers/middleware/gin"
	"github.com/ulule/limiter/v3/drivers/store/memory"
	"log"
	"net/http"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			utils.Response(c, http.StatusUnauthorized, "Missing token", nil)
			c.Abort()
			return
		}
		claims := &dto.Claims{}
		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return config.JwtKey, nil
		})
		if err != nil || !token.Valid {
			utils.Response(c, http.StatusUnauthorized, "Invalid token", nil)
			c.Abort()
			return
		}

		if utils.IsTokenBlacklisted(tokenString) {
			utils.Response(c, http.StatusUnauthorized, "Token is expired", nil)
			c.Abort()
			return
		}
		c.Set("username", claims.Username)
		c.Next()
	}
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if c.Request.Method == http.MethodOptions {
			c.AbortWithStatus(http.StatusOK)
			return
		}

		c.Next()
	}
}

func RateLimiter() gin.HandlerFunc {
	rate, err := limiter.NewRateFromFormatted("60-M")
	if err != nil {
		log.Println("Error creating rate limiter:", err)
	}

	store := memory.NewStore()
	limiterInstance := limiter.New(store, rate)

	middleware := ginlimiter.NewMiddleware(limiterInstance, ginlimiter.WithLimitReachedHandler(func(c *gin.Context) {
		utils.Response(c, http.StatusTooManyRequests, "You have reached the maximum number of requests per minute.", nil)
	}))
	return middleware
}
