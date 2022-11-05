package repository

import (
	"gorm.io/gorm"
)

type ItemRepository struct {
	tx *gorm.DB
}

func (i ItemRepository) Find() error {
	//TODO implement me
	panic("implement me")
}

func (i ItemRepository) Store() error {
	//TODO implement me
	panic("implement me")
}

func (i ItemRepository) Update() error {
	//TODO implement me
	panic("implement me")
}

func (i ItemRepository) FindAll() error {
	//TODO implement me
	panic("implement me")
}

func (i ItemRepository) Delete() error {
	//TODO implement me
	panic("implement me")
}

func NewItemRepository(tx *gorm.DB) *ItemRepository {
	return &ItemRepository{tx: tx}
}
