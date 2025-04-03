package auth

import (
	"github.com/gin-gonic/gin"
	"go-rest-api-gin-gorm/internal/service/auth"
	"go-rest-api-gin-gorm/pkg/utils"
	"net/http"
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
