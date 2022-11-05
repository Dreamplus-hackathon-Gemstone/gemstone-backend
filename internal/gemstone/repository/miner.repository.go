package repository

import "gorm.io/gorm"

type MinerRepository struct {
	tx *gorm.DB
}

func (m MinerRepository) Find() error {
	//TODO implement me
	panic("implement me")
}

func (m MinerRepository) Store() error {
	//TODO implement me
	panic("implement me")
}

func (m MinerRepository) Update() error {
	//TODO implement me
	panic("implement me")
}

func (m MinerRepository) FindAll() error {
	//TODO implement me
	panic("implement me")
}

func (m MinerRepository) Delete() error {
	//TODO implement me
	panic("implement me")
}

func NewMinerRepository(tx *gorm.DB) *MinerRepository {
	return &MinerRepository{tx: tx}
}
