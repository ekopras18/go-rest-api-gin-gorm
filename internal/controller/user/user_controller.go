package user

import (
	"github.com/ekopras18/go-rest-api-gin-gorm/internal/dto/user"
	"github.com/ekopras18/go-rest-api-gin-gorm/internal/models/user"
	"github.com/ekopras18/go-rest-api-gin-gorm/internal/service/user"
	"github.com/ekopras18/go-rest-api-gin-gorm/pkg/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

var userService = user.Service{}

// GetAllUsersHandler godoc
// @Summary Get all users
// @Description Get all users information
// @Tags Users
// @Accept application/json
// @Produce application/json
// @Success 200 {object} utils.ResponseMessage200 "Success"
// @Failure 404 {object} utils.ResponseMessage404 "Not Found"
// @Router /api/v1/users [get]
// @Security BearerAuth
func GetAllUsersHandler(c *gin.Context) {
	users := userService.GetAllUsersService()
	if len(users) == 0 {
		utils.Response(c, http.StatusNotFound, "No data found", nil)
		return
	}

	// using DTO
	var usersResult []dto.UserResponse
	for _, u := range users {
		usersResult = append(usersResult, dto.UserResponse{
			ID:       u.ID,
			Username: u.Username,
			Email:    u.Email,
		})
	}

	utils.Response(c, http.StatusOK, "Success", usersResult)
}

// GetUserHandler godoc
// @Summary Get a user by ID
// @Description Get user information by providing an ID
// @Tags Users
// @Accept application/json
// @Produce application/json
// @Param id path int true "User ID"
// @Success 200 {object} utils.ResponseMessage200 "Success"
// @Failure 404 {object} utils.ResponseMessage404 "Not Found"
// @Router /api/v1/user/{id} [get]
// @Security BearerAuth
func GetUserHandler(c *gin.Context) {
	id := c.Param("id")

	userId, err := strconv.Atoi(id)
	if err != nil {
		utils.Response(c, http.StatusBadRequest, "Invalid user ID", nil)
		return
	}
	u, err := userService.GetUserByIdService(userId)
	if err != nil || u.ID == 0 {
		utils.Response(c, http.StatusNotFound, "ID not found", nil)
		return
	}

	// using DTO
	userResult := dto.UserResponse{
		ID:       u.ID,
		Username: u.Username,
		Email:    u.Email,
	}

	utils.Response(c, http.StatusOK, "Success", userResult)
}

// UpdateUserHandler godoc
// @Summary Update a user by ID
// @Description Update user information by providing an ID
// @Tags Users
// @Accept application/json
// @Produce application/json
// @Param id path int true "User ID"
// @Param user body dto.User true "User data"
// @Success 200 {object} utils.ResponseMessage200 "Success"
// @Failure 400 {object} utils.ResponseMessage400 "Bad Request"
// @Failure 404 {object} utils.ResponseMessage404 "Not Found"
// @Router /api/v1/user/{id} [put]
// @Security BearerAuth
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

	utils.Response(c, http.StatusOK, "Updated successfully", nil)
}

// DeleteUserHandler godoc
// @Summary Delete a user by ID
// @Description Delete user information by providing an ID
// @Tags Users
// @Accept application/json
// @Produce application/json
// @Param id path int true "User ID"
// @Success 200 {object} utils.ResponseMessage200 "Success"
// @Failure 404 {object} utils.ResponseMessage404 "Not Found"
// @Router /api/v1/user/{id} [delete]
// @Security BearerAuth
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

	utils.Response(c, http.StatusOK, "Deleted successfully", nil)
}
