package service

import "gemstone-backend/internal/gemstone/domain"

type MakerService struct {
	makerRepo domain.IRepository
}

func NewMakerService(makerRepo domain.IRepository) *MakerService {
	return &MakerService{makerRepo: makerRepo}
}

func (m *MakerService) Find() error {
	return m.makerRepo.Find()
}

func (m *MakerService) Store() error {
	//TODO implement me
	panic("implement me")
}

func (m *MakerService) Update() error {
	//TODO implement me
	panic("implement me")
}

func (m *MakerService) FindAll() error {
	//TODO implement me
	panic("implement me")
}

func (m *MakerService) Delete() error {
	//TODO implement me
	panic("implement me")
}
