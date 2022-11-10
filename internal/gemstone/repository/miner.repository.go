package repository

import (
	"gemstone-backend/internal/gemstone/transactor"
)

type MinerRepository struct {
	transactor.ITransactor
}

func NewMinerRepository(ITransactor transactor.ITransactor) *MinerRepository {
	return &MinerRepository{ITransactor: ITransactor}
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
