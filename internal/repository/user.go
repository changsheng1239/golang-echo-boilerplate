package repository

import (
	"echo-boilerplate/internal/model"
	"time"
)

type UserRepository interface {
	FirstByID(id int) (*model.User, error)
}

type userRepository struct {
	*Repository
}

func NewUserRepository(r *Repository) UserRepository {
	return &userRepository{
		Repository: r,
	}
}

func (r *userRepository) FirstByID(id int) (*model.User, error) {
	user := &model.User{
		Model:    model.Model{ID: id, CreatedAt: time.Now(), UpdatedAt: time.Now()},
		Username: "test",
		Email:    "test@example.com",
	}
	// TODO
	return user, nil
}
