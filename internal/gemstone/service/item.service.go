package service

import "gemstone-backend/internal/gemstone/domain"

type ItemService struct {
	itemRepo domain.Repository
}

func NewItemService(itemRepo domain.Repository) *ItemService {
	return &ItemService{itemRepo: itemRepo}
}

func (s ItemService) Find() error {
	//TODO implement me
	panic("implement me")
}

func (s ItemService) Store() error {
	//TODO implement me
	panic("implement me")
}

func (s ItemService) Update() error {
	//TODO implement me
	panic("implement me")
}

func (s ItemService) FindAll() error {
	//TODO implement me
	panic("implement me")
}

func (s ItemService) Delete() error {
	//TODO implement me
	panic("implement me")
}
