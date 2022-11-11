package repository

import (
	"gemstone-backend/internal/gemstone/global"
	"gorm.io/gorm"
)

type MinerRepo[P global.Parameter, R global.ReturnType] struct {
	db *gorm.DB
}

func (m *MinerRepo[P, R]) Find(param P) (ret R) {
	//TODO implement me
	panic("implement me")
}

func (m *MinerRepo[P, R]) FindAll(param P) (ret R) {
	//TODO implement me
	panic("implement me")
}

func (m *MinerRepo[P, R]) Store(param P) (ret R) {
	//TODO implement me
	panic("implement me")
}

func (m *MinerRepo[P, R]) Update(param P) (ret R) {
	//TODO implement me
	panic("implement me")
}

func (m *MinerRepo[P, R]) Delete(param P) (ret R) {
	//TODO implement me
	panic("implement me")
}

func NewMinerRepo[P global.Parameter, R global.ReturnType](db *gorm.DB) global.IRepository[P, R] {
	return &MinerRepo[P, R]{db: db}
}
