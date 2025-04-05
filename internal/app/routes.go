package app

import (
	_ "github.com/ekopras18/go-rest-api-gin-gorm/docs"
	"github.com/ekopras18/go-rest-api-gin-gorm/internal/controller/auth"
	"github.com/ekopras18/go-rest-api-gin-gorm/internal/controller/user"
	"github.com/ekopras18/go-rest-api-gin-gorm/pkg/utils"
	"github.com/gin-gonic/gin"
	filesSwagger "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Routes(r *gin.Engine) {
	r.Use(CORSMiddleware())
	r.Use(RateLimiter())
	r.NoRoute(utils.NotFoundHandler)

	r.GET("/docs/*any", ginSwagger.WrapHandler(filesSwagger.Handler))

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
