package repository

import (
	"gemstone-backend/internal/gemstone/global"
	"gorm.io/gorm"
)

type ProposalRepo struct {
	db *gorm.DB
}

func (p *ProposalRepo) FindByContentID(account global.Parameter) (ret global.ReturnType) {
	//TODO implement me
	panic("implement me")
}

func (p *ProposalRepo) FindMany(param global.FindManyReq) (ret global.FindManyProposalRes) {
	//TODO implement me
	panic("implement me")
}

func (p *ProposalRepo) Store(param global.Parameter) (ret global.ReturnType) {
	//TODO implement me
	panic("implement me")
}

func (p *ProposalRepo) Update(param global.Parameter) (ret global.ReturnType) {
	//TODO implement me
	panic("implement me")
}

func (p *ProposalRepo) Delete(param global.Parameter) (ret global.ReturnType) {
	//TODO implement me
	panic("implement me")
}

func NewProposalRepo(db *gorm.DB) global.IProposalRepo {
	return &ProposalRepo{db: db}
}
