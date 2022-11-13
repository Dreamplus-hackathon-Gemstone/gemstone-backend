package service

import "gemstone-backend/internal/gemstone/global"

type MovieService struct {
	movieRepo global.IMovieRepo
}

func NewMovieService(movieRepo global.IMovieRepo) global.IMovieService {
	return &MovieService{movieRepo: movieRepo}
}

func (m *MovieService) FindByContentID(param global.FindMovieReq) (ret global.FindMovieRes) {
	return m.movieRepo.FindByContentID(param)
}

func (m *MovieService) FindMany(param global.FindManyMovieReq) (ret global.FindManyMovieRes) {
	return m.movieRepo.FindMany(param)
}

func (m *MovieService) Store(param global.StoreMovieReq) (ret global.StoreMovieRes) {
	return m.movieRepo.Store(param)
}
