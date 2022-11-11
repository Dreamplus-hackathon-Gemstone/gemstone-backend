package repository

import (
	"gemstone-backend/internal/gemstone/global"
	"gorm.io/gorm"
)

type Repository[P global.Parameter, R global.ReturnType] struct {
	genreRepo    global.IGenreRepo
	makerRepo    global.IExample
	movieRepo    global.IExample
	proposalRepo global.IExample
	minerRepo    global.IExample
	tokenRepo    global.IExample
}

func NewRepository[P global.Parameter, R global.ReturnType](db *gorm.DB) *Repository[P, R] {
	proposal := NewProposalRepo(db)
	genre := NewGenreRepo(db)
	maker := NewMakerRepo(db)
	miner := NewMinerRepo(db)
	token := NewTokenRepo(db)
	movie := NewMovieRepo(db)
	return &Repository[P, R]{
		genreRepo:    genre.(global.IGenreRepo),
		makerRepo:    maker.(global.IExample),
		movieRepo:    movie.(global.IExample),
		proposalRepo: proposal.(global.IExample),
		minerRepo:    miner.(global.IExample),
		tokenRepo:    token.(global.IExample),
	}
}
