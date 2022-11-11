package repository

import (
	"gemstone-backend/internal/gemstone/global"
	"gorm.io/gorm"
)

type ProposalRepo[P global.Parameter, R global.ReturnType] struct {
	db *gorm.DB
}

func (p *ProposalRepo[P, R]) Find(param P) (ret R) {
	//TODO implement me
	panic("implement me")
}

func (p *ProposalRepo[P, R]) FindAll(param P) (ret R) {
	//TODO implement me
	panic("implement me")
}

func (p *ProposalRepo[P, R]) Store(param P) (ret R) {
	//TODO implement me
	panic("implement me")
}

func (p *ProposalRepo[P, R]) Update(param P) (ret R) {
	//TODO implement me
	panic("implement me")
}

func (p *ProposalRepo[P, R]) Delete(param P) (ret R) {
	//TODO implement me
	panic("implement me")
}

func NewProposalRepo[P global.Parameter, R global.ReturnType](db *gorm.DB) global.IRepository[P, R] {
	return &ProposalRepo[P, R]{db: db}
}
