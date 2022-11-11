package repository

import (
	"gemstone-backend/internal/gemstone/global"
	"gorm.io/gorm"
)

type Repository struct {
	GenreRepo    global.IGenreRepo
	MakerRepo    global.IMakerRepo
	MovieRepo    global.IMovieRepo
	ProposalRepo global.IProposalRepo
	MinerRepo    global.IMinerRepo
	TokenRepo    global.ITokenRepo
}

func NewRepository(db *gorm.DB) *Repository {
	proposal := NewProposalRepo(db)
	genre := NewGenreRepo(db)
	maker := NewMakerRepo(db)
	miner := NewMinerRepo(db)
	token := NewTokenRepo(db)
	movie := NewMovieRepo(db)
	return &Repository{
		GenreRepo:    genre,
		MakerRepo:    maker,
		MovieRepo:    movie,
		ProposalRepo: proposal,
		MinerRepo:    miner,
		TokenRepo:    token,
	}
}
