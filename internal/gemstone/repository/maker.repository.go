package repository

import "gorm.io/gorm"

type MakerRepository struct {
	tx *gorm.DB
}

func (m MakerRepository) Find() error {
	//TODO implement me
	panic("implement me")
}

func (m MakerRepository) Store() error {
	//TODO implement me
	panic("implement me")
}

func (m MakerRepository) Update() error {
	//TODO implement me
	panic("implement me")
}

func (m MakerRepository) FindAll() error {
	//TODO implement me
	panic("implement me")
}

func (m MakerRepository) Delete() error {
	//TODO implement me
	panic("implement me")
}

func NewMakerRepository(tx *gorm.DB) *MakerRepository {
	return &MakerRepository{tx: tx}
}
