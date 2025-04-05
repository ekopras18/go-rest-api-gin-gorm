package app

import (
	"github.com/ekopras18/go-rest-api-gin-gorm/config"
	"github.com/ekopras18/go-rest-api-gin-gorm/internal/dto/auth"
	"github.com/ekopras18/go-rest-api-gin-gorm/pkg/utils"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
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
