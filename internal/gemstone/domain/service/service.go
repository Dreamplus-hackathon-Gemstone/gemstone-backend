package service

import (
	"gemstone-backend/internal/gemstone/domain/repository"
	"gemstone-backend/internal/gemstone/global"
)

type Service struct {
	UserService     global.IUserService
	ProposalService global.IProposalService
	MovieService    global.IMovieService
	GenreService    global.IGenreService
}

func NewService(repo *repository.Repository) *Service {
	user := NewUserService(repo.UserRepo)
	proposal := NewProposalService(repo.ProposalRepo)
	movie := NewMovieService(repo.MovieRepo)
	genre := NewGenreService(repo.GenreRepo)
	return &Service{UserService: user, ProposalService: proposal, MovieService: movie, GenreService: genre}
}
