package repository

import (
	"gemstone-backend/internal/gemstone/global"
	"gorm.io/gorm"
)

type Repository struct {
	GenreRepo    global.IGenreRepo
	UserRepo     global.IUserRepo
	MovieRepo    global.IMovieRepo
	ProposalRepo global.IProposalRepo
	TokenRepo    global.ITokenRepo
}

func NewRepository(db *gorm.DB) *Repository {
	proposal := NewProposalRepo(db)
	genre := NewGenreRepo(db)
	user := NewUserRepo(db)
	token := NewTokenRepo(db)
	movie := NewMovieRepo(db)
	return &Repository{
		GenreRepo:    genre,
		MovieRepo:    movie,
		ProposalRepo: proposal,
		UserRepo:     user,
		TokenRepo:    token,
	}
}
