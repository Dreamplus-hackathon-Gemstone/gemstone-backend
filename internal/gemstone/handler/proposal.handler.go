package handler

import "gemstone-backend/internal/gemstone/domain"

type ProposalHandler struct {
	service domain.Service
}

func NewItemHandler(service domain.Service) *ProposalHandler {
	return &ProposalHandler{service: service}
}
