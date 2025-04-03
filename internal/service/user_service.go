package service

import (
	"go-rest-api-gin-gorm/internal/entities"
	"go-rest-api-gin-gorm/internal/repository"
)

type UserService struct {
	repo repository.UserRepository
}

func (s *UserService) GetUsers() []entities.User {
	return s.repo.GetAllUsers()
}
