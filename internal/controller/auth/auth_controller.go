package auth

import (
	"github.com/gin-gonic/gin"
	"go-rest-api-gin-gorm/internal/service/auth"
	"go-rest-api-gin-gorm/pkg/utils"
	"net/http"
)

var authService = auth.Service{}

func LoginHandler(c *gin.Context) {
	var creeds utils.Credentials
	if err := c.BindJSON(&creeds); err != nil {
		utils.Response(c, http.StatusBadRequest, "Invalid request", nil)
		return
	}
	token, err := authService.Login(creeds)
	if err != nil {
		utils.Response(c, http.StatusUnauthorized, "Invalid credentials", nil)
		return
	}

	utils.Response(c, http.StatusOK, "Token generated", gin.H{"token": token})
}

func LogoutHandler(c *gin.Context) {
	tokenString := c.GetHeader("Authorization")
	if tokenString == "" {
		utils.Response(c, http.StatusBadRequest, "Missing token", nil)
		return
	}

	utils.BlacklistToken(tokenString)

	utils.Response(c, http.StatusOK, "Successfully logged out", nil)
}
