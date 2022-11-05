package service

import (
	"gemstone-backend/internal/gemstone/domain"
	"gemstone-backend/internal/gemstone/repository"
)

type Service struct {
	ItemService     domain.Service
	CategoryService domain.Service
	MakerService    domain.Service
	MinerService    domain.Service
	TokenService    domain.Service
}

func NewService(repository *repository.Repository) *Service {
	item := NewItemService(repository.ItemRepo)
	token := NewItemService(repository.TokenRepo)
	maker := NewItemService(repository.MakerRepo)
	miner := NewItemService(repository.MinerRepo)
	category := NewItemService(repository.CategoryRepo)
	return &Service{item, category, maker, miner, token}
}
