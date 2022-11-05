package repository

import "gorm.io/gorm"

type TokenRepository struct {
	tx *gorm.DB
}

func (t TokenRepository) Find() error {
	//TODO implement me
	panic("implement me")
}

func (t TokenRepository) Store() error {
	//TODO implement me
	panic("implement me")
}

func (t TokenRepository) Update() error {
	//TODO implement me
	panic("implement me")
}

func (t TokenRepository) FindAll() error {
	//TODO implement me
	panic("implement me")
}

func (t TokenRepository) Delete() error {
	//TODO implement me
	panic("implement me")
}

func NewTokenRepository(tx *gorm.DB) *TokenRepository {
	return &TokenRepository{tx: tx}
}
