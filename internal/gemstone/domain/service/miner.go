package service

import (
	"gemstone-backend/internal/gemstone/global"
)

type MinerService struct {
	minerRepo global.IMinerRepo
}

func NewMinerService(minerRepo global.IMinerRepo) global.IMinerService {
	return &MinerService{minerRepo: minerRepo}
}

func (m *MinerService) Login(param global.MinerLoginReq) global.LoginRes {
	//TODO implement me
	panic("implement me")
}

func (m *MinerService) Register(param global.RegisterMinerReq) global.RegisterMinerRes {
	//TODO implement me
	panic("implement me")
}
