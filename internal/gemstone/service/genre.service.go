package service

import "gemstone-backend/internal/gemstone/domain"

type GenreService struct {
	genreRepo domain.IRepository
}

func NewGenreService(genreRepo domain.IRepository) *GenreService {
	return &GenreService{genreRepo: genreRepo}
}

func (g *GenreService) Find() error {
	//TODO implement me
	panic("implement me")
}

func (g *GenreService) Store() error {
	//TODO implement me
	panic("implement me")
}

func (g *GenreService) Update() error {
	//TODO implement me
	panic("implement me")
}

func (g *GenreService) FindAll() error {
	//TODO implement me
	panic("implement me")
}

func (g *GenreService) Delete() error {
	//TODO implement me
	panic("implement me")
}
