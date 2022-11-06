package repository

import (
	"gorm.io/gorm"
)

type ProposalRepository struct {
	tx *gorm.DB
}

func (i ProposalRepository) Find() error {
	//TODO implement me
	panic("implement me")
}

func (i ProposalRepository) Store() error {
	//TODO implement me
	panic("implement me")
}

func (i ProposalRepository) Update() error {
	//TODO implement me
	panic("implement me")
}

func (i ProposalRepository) FindAll() error {
	//TODO implement me
	panic("implement me")
}

func (i ProposalRepository) Delete() error {
	//TODO implement me
	panic("implement me")
}

func NewItemRepository(tx *gorm.DB) *ProposalRepository {
	return &ProposalRepository{tx: tx}
}
