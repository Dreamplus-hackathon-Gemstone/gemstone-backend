package repository

import (
	"gemstone-backend/internal/gemstone/global"
	"gorm.io/gorm"
)

type GenreRepo struct {
	db *gorm.DB
}

func (g *GenreRepo) FindByGenre(param global.Parameter) (ret global.ReturnType) {
	//TODO implement me
	panic("implement me")
}

func (g *GenreRepo) Store(param global.Parameter) (ret global.ReturnType) {
	//TODO implement me
	panic("implement me")
}

func NewGenreRepo(db *gorm.DB) global.IGenreRepo {
	return &GenreRepo{db: db}
}
