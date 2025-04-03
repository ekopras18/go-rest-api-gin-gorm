package user

import (
	"go-rest-api-gin-gorm/internal/models"
	"go-rest-api-gin-gorm/internal/repository/user"
)

type Service struct {
	repo user.Repository
}

func (s *Service) GetAllUsersService() []models.User {
	return s.repo.GetAllUsersRepository()
}
