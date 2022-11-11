package repository

import (
	"gemstone-backend/internal/gemstone/global"
	"gorm.io/gorm"
)

type MovieRepo[P global.Parameter, R global.ReturnType] struct {
	db *gorm.DB
}

func (m *MovieRepo[P, R]) Find(param P) (ret R) {
	//TODO implement me
	panic("implement me")
}

func (m *MovieRepo[P, R]) FindAll(param P) (ret R) {
	//TODO implement me
	panic("implement me")
}

func (m *MovieRepo[P, R]) Store(param P) (ret R) {
	//TODO implement me
	panic("implement me")
}

func (m *MovieRepo[P, R]) Update(param P) (ret R) {
	//TODO implement me
	panic("implement me")
}

func (m *MovieRepo[P, R]) Delete(param P) (ret R) {
	//TODO implement me
	panic("implement me")
}

func NewMovieRepo[P global.Parameter, R global.ReturnType](db *gorm.DB) global.IRepository[P, R] {
	return &MovieRepo[P, R]{db: db}
}
