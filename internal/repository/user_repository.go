package repository

import (
	"rest-api-gin-gorm/internal/database"
	"rest-api-gin-gorm/internal/entities"
)

type UserRepository struct{}

func (r *UserRepository) GetAllUsers() []entities.User {
	var users []entities.User
	database.Db.Find(&users)
	return users
}
