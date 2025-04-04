package user

import (
	"go-rest-api-gin-gorm/internal/database"
	"go-rest-api-gin-gorm/internal/models"
	"time"
)

type Repository struct{}

func (r *Repository) GetAllUsersRepository() []models.User {
	var users []models.User
	database.Db.Order("id desc").Find(&users)
	return users
}

func (r *Repository) GetUserByIdRepository(id int) models.User {
	var user models.User
	database.Db.Where("id = ?", id).First(&user)
	return user
}

func (r *Repository) UpdateUserByIdRepository(id int, data models.User) (models.User, error) {
	var user models.User
	if err := database.Db.First(&user, id).Error; err != nil {
		return user, err
	}

	user.Username = data.Username
	//user.Email = data.Email
	user.UpdatedAt = time.Now()

	database.Db.Save(&user)

	return user, nil
}

func (r *Repository) DeleteUserByIdRepository(id int) error {
	var user models.User
	if err := database.Db.First(&user, id).Error; err != nil {
		return err
	}

	database.Db.Delete(&user)

	return nil
}
