package service

import "gemstone-backend/internal/gemstone/domain"

type service struct {
	genreRepo domain.Repository
}

func (s service) Find() error {
	//TODO implement me
	panic("implement me")
}

func (s service) Store() error {
	//TODO implement me
	panic("implement me")
}

func (s service) Update() error {
	//TODO implement me
	panic("implement me")
}

func (s service) FindAll() error {
	//TODO implement me
	panic("implement me")
}

func (s service) Delete() error {
	//TODO implement me
	panic("implement me")
}

func newService(categoryRepo domain.Repository) *service {
	return &service{genreRepo: categoryRepo}
}
