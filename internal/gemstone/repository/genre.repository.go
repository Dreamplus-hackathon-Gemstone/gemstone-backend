package repository

import "gorm.io/gorm"

type GenreRepository struct {
	tx *gorm.DB
}

func (c GenreRepository) Find() error {
	//TODO implement me
	panic("implement me")
}

func (c GenreRepository) Store() error {
	//TODO implement me
	panic("implement me")
}

func (c GenreRepository) Update() error {
	//TODO implement me
	panic("implement me")
}

func (c GenreRepository) FindAll() error {
	//TODO implement me
	panic("implement me")
}

func (c GenreRepository) Delete() error {
	//TODO implement me
	panic("implement me")
}

func NewCategoryRepository(tx *gorm.DB) *GenreRepository {
	return &GenreRepository{tx: tx}
}
