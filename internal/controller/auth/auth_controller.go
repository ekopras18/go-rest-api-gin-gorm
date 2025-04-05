package auth

import (
	"github.com/ekopras18/go-rest-api-gin-gorm/internal/dto/auth"
	"github.com/ekopras18/go-rest-api-gin-gorm/internal/service/auth"
	"github.com/ekopras18/go-rest-api-gin-gorm/pkg/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

var authService = auth.Service{}

// LoginHandler godoc
// @Summary Login
// @Description Login to the system
// @Tags Authentication
// @Accept application/json
// @Produce application/json
// @Param credentials body dto.Credentials true "User credentials"
// @Success 200 {object} utils.ResponseMessage200 "Success"
// @Success 201 {object} utils.ResponseMessage201 "Success"
// @Failure 404 {object} utils.ResponseMessage404 "Not Found"
// @Router /auth/login [post]
func LoginHandler(c *gin.Context) {
	var credential dto.Credentials
	if err := c.BindJSON(&credential); err != nil {
		utils.Response(c, http.StatusBadRequest, "Invalid request", nil)
		return
	}
	token, err := authService.Login(credential)
	if err != nil {
		utils.Response(c, http.StatusUnauthorized, "Invalid credentials", nil)
		return
	}

	utils.Response(c, http.StatusOK, "Token generated", gin.H{"token": token})
}

// LogoutHandler godoc
// @Summary Logout
// @Description Logout to the system
// @Tags Authentication
// @Accept application/json
// @Produce application/json
// @Success 200 {object} utils.ResponseMessage200 "Success"
// @Success 201 {object} utils.ResponseMessage201 "Success"
// @Failure 404 {object} utils.ResponseMessage404 "Not Found"
// @Router /auth/logout [post]
// @Security BearerAuth
func LogoutHandler(c *gin.Context) {
	tokenString := c.GetHeader("Authorization")
	if tokenString == "" {
		utils.Response(c, http.StatusBadRequest, "Missing token", nil)
		return
	}

	utils.BlacklistToken(tokenString)

	utils.Response(c, http.StatusOK, "Successfully logged out", nil)
}

// RegisterHandler godoc
// @Summary Register
// @Description Register a new user
// @Tags Authentication
// @Accept application/json
// @Produce application/json
// @Param credentials body dto.Register true "User credentials"
// @Success 201 {object} utils.ResponseMessage201 "User registered"
// @Failure 409 {object} utils.ResponseMessage409 "Conflict"
// @Failure 400 {object} utils.ResponseMessage400 "Bad Request"
// @Router /auth/register [post]
func RegisterHandler(c *gin.Context) {
	var credential dto.Register
	if err := c.BindJSON(&credential); err != nil {
		utils.Response(c, http.StatusBadRequest, "Invalid request", nil)
		return
	}
	if err := authService.Register(credential); err != nil {
		utils.Response(c, http.StatusConflict, err.Error(), nil)
		return
	}
	utils.Response(c, http.StatusCreated, "Successfully registered", nil)
}
