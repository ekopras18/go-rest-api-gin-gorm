package user

import (
	"github.com/ekopras18/go-rest-api-gin-gorm/internal/models/user"
	"github.com/ekopras18/go-rest-api-gin-gorm/internal/repository/user"
)

type Service struct {
	repo user.Repository
}

func (s *Service) GetAllUsersService() []models.User {
	return s.repo.GetAllUsersRepository()
}

func (s *Service) GetUserByIdService(id int) (models.User, error) {
	return s.repo.GetUserByIdRepository(id), nil
}

func (s *Service) UpdateUserByIdService(id int, data models.User) (models.User, error) {
	return s.repo.UpdateUserByIdRepository(id, data)
}

func (s *Service) DeleteUserByIdService(id int) error {
	return s.repo.DeleteUserByIdRepository(id)
}
