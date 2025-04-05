package auth

import (
	"github.com/ekopras18/go-rest-api-gin-gorm/internal/database"
	"github.com/ekopras18/go-rest-api-gin-gorm/internal/models/user"
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

func (r *Repository) IsEmailTaken(email string) bool {
	var user models.User
	database.Db.Where("email = ?", email).First(&user)
	return user.ID != 0
}

func (r *Repository) RegisterUserRepository(username, password string) error {
	user := models.User{Username: username, Password: password, CreatedAt: time.Now()}
	return database.Db.Create(&user).Error
}
