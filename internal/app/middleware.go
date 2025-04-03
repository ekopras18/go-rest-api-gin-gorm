package app

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	"rest-api-gin-gorm/config"
	"rest-api-gin-gorm/pkg/utils"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			utils.Response(c, http.StatusUnauthorized, "Missing token", nil)
			c.Abort()
			return
		}
		claims := &utils.Claims{}
		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return config.JwtKey, nil
		})
		if err != nil || !token.Valid {
			utils.Response(c, http.StatusUnauthorized, "Invalid token", nil)
			c.Abort()
			return
		}
		c.Set("username", claims.Username)
		c.Next()
	}
}
