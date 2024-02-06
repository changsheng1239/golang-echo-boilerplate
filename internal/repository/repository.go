package repository

import "echo-boilerplate/pkg/log"

type Repository struct {
	logger *log.Logger
}

func NewRepository(logger *log.Logger) *Repository {
	return &Repository{logger}
}
