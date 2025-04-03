package repository

import (
	"go-rest-api-gin-gorm/internal/database"
	"go-rest-api-gin-gorm/internal/entities"
)

type UserRepository struct{}

func (r *UserRepository) GetAllUsers() []entities.User {
	var users []entities.User
	database.Db.Find(&users)
	return users
}
