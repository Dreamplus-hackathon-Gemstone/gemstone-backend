package repository

import (
	"gemstone-backend/internal/gemstone/transactor"
)

type TokenRepository struct {
	transactor transactor.ITransactor
}

func NewTokenRepository(transactor transactor.ITransactor) *TokenRepository {
	return &TokenRepository{transactor: transactor}
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
