package service

import (
	"gemstone-backend/internal/gemstone/global"
)

type MakerService struct {
	makerRepo global.IMakerRepo
}

func (m *MakerService) VerifyNickname(param global.VerifyNicknameReq) global.VerifyNicknameRes {
	return m.makerRepo.FindByNickname(param)
}

func (m *MakerService) UpdateNickname(param global.UpdateMakerNicknameReq) global.UpdateMakerNicknameRes {
	//TODO implement me
	panic("implement me")
}

func (m *MakerService) Login(param global.MakerLoginReq) global.LoginRes {
	//TODO implement me
	panic("implement me")
}

func (m *MakerService) Register(param global.RegisterMakerReq) global.RegisterMakerRes {
	//TODO implement me
	panic("implement me")
}

func NewMakerService(makerRepo global.IMakerRepo) global.IMakerService {
	return &MakerService{makerRepo: makerRepo}
}
