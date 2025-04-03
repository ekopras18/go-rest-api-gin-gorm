package user

import (
	"github.com/gin-gonic/gin"
	"go-rest-api-gin-gorm/internal/service/user"
	"go-rest-api-gin-gorm/pkg/utils"
	"net/http"
)

var userService = user.Service{}

func GetAllUsersHandler(c *gin.Context) {
	users := userService.GetAllUsersService()
	if len(users) == 0 {
		utils.Response(c, http.StatusNotFound, "No data found", nil)
		return
	}

	utils.Response(c, http.StatusOK, "Success", users)
}
