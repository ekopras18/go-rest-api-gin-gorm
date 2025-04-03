package auth

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"rest-api-gin-gorm/internal/service/auth"
	"rest-api-gin-gorm/pkg/utils"
)

var authService = auth.AuthService{}

func LoginHandler(c *gin.Context) {
	var creds utils.Credentials
	if err := c.BindJSON(&creds); err != nil {
		utils.Response(c, http.StatusBadRequest, "Invalid request", nil)
		return
	}
	token, err := authService.Login(creds)
	if err != nil {
		utils.Response(c, http.StatusUnauthorized, "Invalid credentials", nil)
		return
	}

	utils.Response(c, http.StatusOK, "Token generated", gin.H{"token": token})
}
