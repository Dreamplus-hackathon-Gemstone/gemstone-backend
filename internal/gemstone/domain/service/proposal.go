package service

import "gemstone-backend/internal/gemstone/global"

type ProposalService[R global.ReturnType, P global.Parameter] struct {
	proposalRepo global.IRepository[P, R]
}

func (p *ProposalService[R, P]) Find(p2 P) R {
	//TODO implement me
	panic("implement me")
}

func (p *ProposalService[R, P]) FindAll(p2 P) R {
	//TODO implement me
	panic("implement me")
}

func (p *ProposalService[R, P]) Store(p2 P) R {
	//TODO implement me
	panic("implement me")
}

func (p *ProposalService[R, P]) Update(p2 P) R {
	//TODO implement me
	panic("implement me")
}

func (p *ProposalService[R, P]) Delete(p2 P) R {
	//TODO implement me
	panic("implement me")
}
