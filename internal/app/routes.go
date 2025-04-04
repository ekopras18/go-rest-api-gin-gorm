package app

import (
	"github.com/gin-gonic/gin"
	"go-rest-api-gin-gorm/internal/controller/auth"
	"go-rest-api-gin-gorm/internal/controller/user"
	"go-rest-api-gin-gorm/pkg/utils"
)

func Routes(r *gin.Engine) {
	r.NoRoute(utils.NotFoundHandler)

	// Auth routes
	routesAuth := r.Group("/auth")
	routesAuth.POST("/register", auth.RegisterHandler)
	routesAuth.POST("/login", auth.LoginHandler)
	routesAuth.POST("/logout", auth.LogoutHandler)

	// v1 api & middleware
	routes := r.Group("/api/v1")
	routes.Use(AuthMiddleware())

	// User routes
	routes.GET("/users", user.GetAllUsersHandler)
	routes.GET("/user/:id", user.GetUserHandler)
	routes.PUT("/user/:id", user.UpdateUserHandler)
	routes.DELETE("/user/:id", user.DeleteUserHandler)
}
