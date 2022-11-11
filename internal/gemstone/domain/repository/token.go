package repository

import (
	"gemstone-backend/internal/gemstone/global"
	"gorm.io/gorm"
)

type TokenRepo[P global.Parameter, R global.ReturnType] struct {
	db *gorm.DB
}

func (t *TokenRepo[P, R]) Find(param P) (ret R) {
	//TODO implement me
	panic("implement me")
}

func (t *TokenRepo[P, R]) FindAll(param P) (ret R) {
	//TODO implement me
	panic("implement me")
}

func (t *TokenRepo[P, R]) Store(param P) (ret R) {
	//TODO implement me
	panic("implement me")
}

func (t *TokenRepo[P, R]) Update(param P) (ret R) {
	//TODO implement me
	panic("implement me")
}

func (t *TokenRepo[P, R]) Delete(param P) (ret R) {
	//TODO implement me
	panic("implement me")
}

func NewTokenRepo[P global.Parameter, R global.ReturnType](db *gorm.DB) global.IRepository[P, R] {
	return &TokenRepo[P, R]{db: db}
}
