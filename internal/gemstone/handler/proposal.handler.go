package handler

import "gemstone-backend/internal/gemstone/domain"

type ProposalHandler struct {
	service domain.IService
}

func NewItemHandler(service domain.IService) *ProposalHandler {
	return &ProposalHandler{service: service}
}
