package user

import (
	"github.com/gin-gonic/gin"
	"go-rest-api-gin-gorm/internal/models"
	"go-rest-api-gin-gorm/internal/service/user"
	"go-rest-api-gin-gorm/pkg/utils"
	"net/http"
	"strconv"
)

var userService = user.Service{}

func GetAllUsersHandler(c *gin.Context) {
	usersResult := userService.GetAllUsersService()
	if len(usersResult) == 0 {
		utils.Response(c, http.StatusNotFound, "No data found", nil)
		return
	}

	utils.Response(c, http.StatusOK, "Success", usersResult)
}

func GetUserHandler(c *gin.Context) {
	id := c.Param("id")

	userId, err := strconv.Atoi(id)
	if err != nil {
		utils.Response(c, http.StatusBadRequest, "Invalid user ID", nil)
		return
	}
	userResult, err := userService.GetUserByIdService(userId)
	if err != nil || userResult.ID == 0 {
		utils.Response(c, http.StatusNotFound, "ID not found", nil)
		return
	}

	utils.Response(c, http.StatusOK, "Success", userResult)
}

func UpdateUserHandler(c *gin.Context) {
	id := c.Param("id")
	userId, err := strconv.Atoi(id)
	if err != nil {
		utils.Response(c, http.StatusBadRequest, "Invalid user ID", nil)
		return
	}

	var updateData models.User
	if err := c.ShouldBindJSON(&updateData); err != nil {
		utils.Response(c, http.StatusBadRequest, "Invalid JSON body", nil)
		return
	}

	_, err = userService.UpdateUserByIdService(userId, updateData)
	if err != nil {
		utils.Response(c, http.StatusInternalServerError, "Failed to update user", nil)
		return
	}

	utils.Response(c, http.StatusOK, "User updated successfully", nil)
}

func DeleteUserHandler(c *gin.Context) {
	id := c.Param("id")
	userId, err := strconv.Atoi(id)
	if err != nil {
		utils.Response(c, http.StatusBadRequest, "Invalid user ID", nil)
		return
	}

	err = userService.DeleteUserByIdService(userId)
	if err != nil {
		utils.Response(c, http.StatusNotFound, "Failed to delete user", nil)
		return
	}

	utils.Response(c, http.StatusOK, "User deleted successfully", nil)
}
