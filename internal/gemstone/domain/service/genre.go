package service

import "gemstone-backend/internal/gemstone/global"

type GenreService struct {
	genreRepo global.IGenreRepo
}

func NewGenreService(genreRepo global.IGenreRepo) global.IGenreService {
	return &GenreService{genreRepo: genreRepo}
}

func (g *GenreService) Store(param global.StoreGenreReq) (ret global.StoreGenreRes) {
	return g.genreRepo.Store(param)
}

func (g *GenreService) FindMovie(param global.FindMovieByGenreReq) (ret global.FindMovieByGenreRes) {
	return g.genreRepo.FindMovieByGenre(param)
}

func (g *GenreService) FindProposal(param global.FindProposalByGenreReq) (ret global.FindProposalByGenreRes) {
	return g.genreRepo.FindProposalByGenre(param)
}
