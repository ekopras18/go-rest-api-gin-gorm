package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"rest-api-gin-gorm/internal/service"
	"rest-api-gin-gorm/pkg/utils"
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
