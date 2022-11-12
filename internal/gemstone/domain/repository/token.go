package repository

import (
	"gemstone-backend/internal/gemstone/global"
	"gorm.io/gorm"
)

type TokenRepo struct {
	db *gorm.DB
}

func NewTokenRepo(db *gorm.DB) global.ITokenRepo {
	return &TokenRepo{db: db}
}

func (t *TokenRepo) Find(account global.Parameter) (ret global.ReturnType) {
	//TODO implement me
	panic("implement me")
}

func (t *TokenRepo) Store(param global.Parameter) (ret global.ReturnType) {
	//TODO implement me
	panic("implement me")
}
