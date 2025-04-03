package repository

import (
	"go-rest-api-gin-gorm/internal/database"
	"go-rest-api-gin-gorm/internal/models"
)

type UserRepository struct{}

func (r *UserRepository) GetAllUsers() []models.User {
	var users []models.User
	database.Db.Find(&users)
	return users
}
