package repository

import (
	"gemstone-backend/internal/gemstone/global"
	"gorm.io/gorm"
)

type MakerRepo[P global.Parameter, R global.ReturnType] struct {
	db *gorm.DB
}

func (m *MakerRepo[P, R]) Find(param P) (ret R) {
	//TODO implement me
	panic("implement me")
}

func (m *MakerRepo[P, R]) FindAll(param P) (ret R) {
	//TODO implement me
	panic("implement me")
}

func (m *MakerRepo[P, R]) Store(param P) (ret R) {
	//TODO implement me
	panic("implement me")
}

func (m *MakerRepo[P, R]) Update(param P) (ret R) {
	//TODO implement me
	panic("implement me")
}

func (m *MakerRepo[P, R]) Delete(param P) (ret R) {
	//TODO implement me
	panic("implement me")
}

func NewMakerRepo[P global.Parameter, R global.ReturnType](db *gorm.DB) global.IRepository[P, R] {
	return &MakerRepo[P, R]{db: db}
}
