package service

import (
	"gemstone-backend/internal/gemstone/domain/repository"
	"gemstone-backend/internal/gemstone/global"
)

type Service struct {
	MakerService global.IMakerService
	MinerService global.IMinerService
}

func NewService(repo *repository.Repository) *Service {
	maker := NewMakerService(repo.MakerRepo)
	miner := NewMinerService(repo.MinerRepo)
	return &Service{MakerService: maker, MinerService: miner}
}
