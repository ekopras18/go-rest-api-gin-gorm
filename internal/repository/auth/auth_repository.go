package auth

import (
	"go-rest-api-gin-gorm/internal/database"
	"go-rest-api-gin-gorm/internal/models"
)

type AuthRepository struct{}

func (r *AuthRepository) GetUserByUsername(username string) (models.User, error) {
	var user models.User
	err := database.Db.Where("username = ?", username).First(&user).Error
	return user, err
}
