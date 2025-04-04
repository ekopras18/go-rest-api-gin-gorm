package auth

import (
	"go-rest-api-gin-gorm/internal/database"
	"go-rest-api-gin-gorm/internal/models"
	"time"
)

type Repository struct{}

func (r *Repository) GetUserByUsername(username string) (models.User, error) {
	var user models.User
	err := database.Db.Where("username = ?", username).First(&user).Error
	return user, err
}

func (r *Repository) IsUsernameTaken(username string) bool {
	var user models.User
	database.Db.Where("username = ?", username).First(&user)
	return user.ID != 0
}

func (r *Repository) RegisterUserRepository(username, password string) error {
	user := models.User{Username: username, Password: password, CreatedAt: time.Now()}
	return database.Db.Create(&user).Error
}
