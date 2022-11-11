package repository

import (
	"gemstone-backend/internal/gemstone/global"
	"gorm.io/gorm"
)

type MinerRepo struct {
	db *gorm.DB
}

func NewMinerRepo(db *gorm.DB) global.IMinerRepo {
	return &MinerRepo{db: db}
}

func (m *MinerRepo) Find(account global.FindMinerReq) (ret global.FindMinerRes) {
	//TODO implement me
	panic("implement me")
}

func (m *MinerRepo) Store(param global.RegisterMinerReq) (ret global.RegisterMinerRes) {
	//TODO implement me
	panic("implement me")
}
