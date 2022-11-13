package service

import (
	"gemstone-backend/internal/gemstone/global"
)

type ProposalService struct {
	proposalRepo global.IProposalRepo
}

func (p *ProposalService) Find(param global.FindProposalReq) (ret global.FindProposalRes) {
	return p.proposalRepo.FindByContentID(param)
}

func (p *ProposalService) FindMany(param global.FindManyProposalReq) (ret global.FindManyProposalRes) {
	return p.proposalRepo.FindMany(param)
}

func (p *ProposalService) Store(param global.StoreProposalReq) (ret global.StoreProposalRes) {
	return p.proposalRepo.Store(param)
}

func (p *ProposalService) Update(param global.UpdateProposalReq) (ret global.UpdateProposalRes) {
	return p.proposalRepo.Update(param)
}

func (p *ProposalService) Delete(param global.DeleteProposalReq) (ret global.DeleteProposalRes) {
	return p.proposalRepo.Delete(param)
}

func NewProposalService(proposalRepo global.IProposalRepo) global.IProposalService {
	return &ProposalService{proposalRepo: proposalRepo}
}
