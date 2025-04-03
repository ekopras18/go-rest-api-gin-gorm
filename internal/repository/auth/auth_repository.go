package auth

import (
	"go-rest-api-gin-gorm/internal/database"
	"go-rest-api-gin-gorm/internal/entities"
)

type AuthRepository struct{}

func (r *AuthRepository) GetUserByUsername(username string) (entities.User, error) {
	var user entities.User
	err := database.Db.Where("username = ?", username).First(&user).Error
	return user, err
}
