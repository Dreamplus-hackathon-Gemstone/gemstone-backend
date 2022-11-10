package repository

import (
	"gemstone-backend/internal/gemstone/transactor"
)

type GenreRepository struct {
	transactor.ITransactor
}

func NewGenreRepository(ITransactor transactor.ITransactor) *GenreRepository {
	return &GenreRepository{ITransactor: ITransactor}
}

func (c *GenreRepository) Find() error {
	//TODO implement me
	panic("implement me")
}

func (c *GenreRepository) Store() error {
	//TODO implement me
	panic("implement me")
}

func (c *GenreRepository) Update() error {
	//TODO implement me
	panic("implement me")
}

func (c *GenreRepository) FindAll() error {
	//TODO implement me
	panic("implement me")
}

func (c *GenreRepository) Delete() error {
	//TODO implement me
	panic("implement me")
}
