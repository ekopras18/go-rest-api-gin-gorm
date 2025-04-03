package controller

import (
	"github.com/gin-gonic/gin"
	"go-rest-api-gin-gorm/internal/service"
	"go-rest-api-gin-gorm/pkg/utils"
	"net/http"
)

var userService = service.UserService{}

func GetUsersHandler(c *gin.Context) {
	users := userService.GetUsers()
	if len(users) == 0 {
		utils.Response(c, http.StatusNotFound, "No data found", nil)
		return
	}

	utils.Response(c, http.StatusOK, "Success", users)
}
