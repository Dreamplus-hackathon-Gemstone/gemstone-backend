package repository

import (
	"gemstone-backend/internal/gemstone/global"
	"gorm.io/gorm"
)

type MovieRepo struct {
	db *gorm.DB
}

func NewMovieRepo(db *gorm.DB) global.IMovieRepo {
	return &MovieRepo{db: db}
}

func (m *MovieRepo) Find(account global.FindItemReq) (ret global.FindMovieRes) {
	//TODO implement me
	panic("implement me")
}

func (m *MovieRepo) FindMany(param global.FindManyProposalReq) (ret global.ReturnType) {
	//TODO implement me
	panic("implement me")
}

func (m *MovieRepo) Store(param global.Parameter) (ret global.ReturnType) {
	//TODO implement me
	panic("implement me")
}

func (m *MovieRepo) Update(param global.Parameter) (ret global.ReturnType) {
	//TODO implement me
	panic("implement me")
}

func (m *MovieRepo) Delete(param global.Parameter) (ret global.ReturnType) {
	//TODO implement me
	panic("implement me")
}
