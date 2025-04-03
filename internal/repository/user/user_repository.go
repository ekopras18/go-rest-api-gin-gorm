package user

import (
	"go-rest-api-gin-gorm/internal/database"
	"go-rest-api-gin-gorm/internal/models"
)

type Repository struct{}

func (r *Repository) GetAllUsersRepository() []models.User {
	var users []models.User
	database.Db.Find(&users)
	return users
}
