package service

import "gemstone-backend/internal/gemstone/domain"

type ProposalService struct {
	itemRepo domain.IRepository
}

func NewItemService(itemRepo domain.IRepository) *ProposalService {
	return &ProposalService{itemRepo: itemRepo}
}

func (s ProposalService) Find() error {
	//TODO implement me
	panic("implement me")
}

func (s ProposalService) Store() error {
	//TODO implement me
	panic("implement me")
}

func (s ProposalService) Update() error {
	//TODO implement me
	panic("implement me")
}

func (s ProposalService) FindAll() error {
	//TODO implement me
	panic("implement me")
}

func (s ProposalService) Delete() error {
	//TODO implement me
	panic("implement me")
}
