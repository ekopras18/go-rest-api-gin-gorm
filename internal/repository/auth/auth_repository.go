package auth

import (
	"go-rest-api-gin-gorm/internal/database"
	"go-rest-api-gin-gorm/internal/models"
)

type Repository struct{}

func (r *Repository) GetUserByUsername(username string) (models.User, error) {
	var user models.User
	err := database.Db.Where("username = ?", username).First(&user).Error
	return user, err
}
