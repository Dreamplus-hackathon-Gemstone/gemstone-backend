package service

import "gemstone-backend/internal/gemstone/domain"

type MinerService struct {
	minerRepo domain.IRepository
}

func NewMinerService(minerRepo domain.IRepository) *MinerService {
	return &MinerService{minerRepo: minerRepo}
}

func (s *MinerService) Find() error {
	//TODO implement me
	panic("implement me")
}

func (s *MinerService) Store() error {
	//TODO implement me
	panic("implement me")
}

func (s *MinerService) Update() error {
	//TODO implement me
	panic("implement me")
}

func (s *MinerService) FindAll() error {
	//TODO implement me
	panic("implement me")
}

func (s *MinerService) Delete() error {
	//TODO implement me
	panic("implement me")
}
