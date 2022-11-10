package repository

import (
	"gemstone-backend/internal/gemstone/transactor"
)

type ProposalRepository struct {
	transactor.ITransactor
}

func NewProposalRepository(ITransactor transactor.ITransactor) *ProposalRepository {
	return &ProposalRepository{ITransactor: ITransactor}
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
