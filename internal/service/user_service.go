package service

import (
	"go-rest-api-gin-gorm/internal/models"
	"go-rest-api-gin-gorm/internal/repository"
)

type UserService struct {
	repo repository.UserRepository
}

func (s *UserService) GetUsers() []models.User {
	return s.repo.GetAllUsers()
}
