package service

import (
	"gemstone-backend/internal/gemstone/domain/repository"
	"gemstone-backend/internal/gemstone/global"
)

type Service struct {
	UserService global.IUserService
}

func NewService(repo *repository.Repository) *Service {
	user := NewUserService(repo.UserRepo)
	return &Service{user}
}
