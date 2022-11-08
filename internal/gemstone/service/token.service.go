package service

import "gemstone-backend/internal/gemstone/domain"

type TokenService struct {
	tokenRepo domain.IRepository
}

func NewTokenService(tokenRepo domain.IRepository) *TokenService {
	return &TokenService{tokenRepo: tokenRepo}
}

func (s TokenService) Find() error {
	//TODO implement me
	panic("implement me")
}

func (s TokenService) Store() error {
	//TODO implement me
	panic("implement me")
}

func (s TokenService) Update() error {
	//TODO implement me
	panic("implement me")
}

func (s TokenService) FindAll() error {
	//TODO implement me
	panic("implement me")
}

func (s TokenService) Delete() error {
	//TODO implement me
	panic("implement me")
}
