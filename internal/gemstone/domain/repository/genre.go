package repository

import (
	"gemstone-backend/internal/gemstone/global"
	"gorm.io/gorm"
)

type GenreRepo[P global.Parameter, R global.ReturnType] struct {
	db *gorm.DB
}

func (g *GenreRepo[P, R]) Find(param P) (ret R) {
	//TODO implement me
	panic("implement me")
}

func (g *GenreRepo[P, R]) FindAll(param P) (ret R) {
	//TODO implement me
	panic("implement me")
}

func (g *GenreRepo[P, R]) Store(param P) (ret R) {
	//TODO implement me
	panic("implement me")
}

func (g *GenreRepo[P, R]) Update(param P) (ret R) {
	//TODO implement me
	panic("implement me")
}

func (g *GenreRepo[P, R]) Delete(param P) (ret R) {
	//TODO implement me
	panic("implement me")
}

func NewGenreRepo[P global.Parameter, R global.ReturnType](db *gorm.DB) global.IRepository[P, R] {
	return &GenreRepo[P, R]{db: db}
}
