package app

import (
	"github.com/gin-gonic/gin"
	"go-rest-api-gin-gorm/internal/controller/auth"
	"go-rest-api-gin-gorm/internal/controller/user"
	"go-rest-api-gin-gorm/internal/database"
	"go-rest-api-gin-gorm/internal/models"
	"go-rest-api-gin-gorm/pkg/utils"
	"net/http"
	"time"
)

func register(c *gin.Context) {
	var credential utils.Credentials
	if err := c.BindJSON(&credential); err != nil {
		utils.Response(c, http.StatusBadRequest, "Invalid request", nil)
		return
	}

	var existingUser models.User
	if err := database.Db.Where("username = ?", credential.Username).First(&existingUser).Error; err == nil && existingUser.ID != 0 {
		utils.Response(c, http.StatusConflict, "Username already exists", nil)
		return
	}

	hashedPassword, err := utils.HashPassword(credential.Password)
	if err != nil {
		utils.Response(c, http.StatusInternalServerError, "Error hashing password", nil)
		return
	}

	userData := models.User{Username: credential.Username, Password: hashedPassword, CreatedAt: time.Now()}
	if err := database.Db.Create(&userData).Error; err != nil {
		utils.Response(c, http.StatusInternalServerError, "Could not create user", nil)
		return
	}

	utils.Response(c, http.StatusCreated, "User registered", nil)
}

func Routes(r *gin.Engine) {
	r.NoRoute(utils.NotFoundHandler)

	// Auth routes
	routesAuth := r.Group("/auth")
	routesAuth.POST("/register", register)
	routesAuth.POST("/login", auth.LoginHandler)
	routesAuth.POST("/logout", auth.LogoutHandler)

	routes := r.Group("/api/v1")
	routes.Use(AuthMiddleware())

	// User routes
	routes.GET("/users", user.GetAllUsersHandler)
}
