package service

import (
	"gemstone-backend/internal/gemstone/domain"
	"gemstone-backend/internal/gemstone/repository"
)

type Service struct {
	ProposalService domain.IService
	GenreService    domain.IService
	MakerService    domain.IService
	MinerService    domain.IService
	TokenService    domain.IService
}

func NewService(repository *repository.Repository) *Service {
	proposal := NewProposalService(repository.ItemRepo)
	token := NewTokenService(repository.TokenRepo)
	maker := NewMakerService(repository.MakerRepo)
	miner := NewMinerService(repository.MinerRepo)
	category := NewGenreService(repository.CategoryRepo)
	return &Service{proposal, category, maker, miner, token}
}
