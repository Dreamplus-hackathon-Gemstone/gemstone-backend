package repository

import "gorm.io/gorm"

type CategoryRepository struct {
	tx *gorm.DB
}

func (c CategoryRepository) Find() error {
	//TODO implement me
	panic("implement me")
}

func (c CategoryRepository) Store() error {
	//TODO implement me
	panic("implement me")
}

func (c CategoryRepository) Update() error {
	//TODO implement me
	panic("implement me")
}

func (c CategoryRepository) FindAll() error {
	//TODO implement me
	panic("implement me")
}

func (c CategoryRepository) Delete() error {
	//TODO implement me
	panic("implement me")
}

func NewCategoryRepository(tx *gorm.DB) *CategoryRepository {
	return &CategoryRepository{tx: tx}
}
