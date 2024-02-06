package service

import (
	"echo-boilerplate/internal/model"
	"echo-boilerplate/internal/repository"
)

type UserService interface {
	GetUserByID(id int) (*model.User, error)
}

type userService struct {
	*Service
	userRepository repository.UserRepository
}

func NewUserService(s *Service, ur repository.UserRepository) UserService {
	return &userService{s, ur}
}

func (us *userService) GetUserByID(id int) (*model.User, error) {
	return us.userRepository.FirstByID(id)
}
