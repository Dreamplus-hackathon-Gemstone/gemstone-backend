package handler

import "gemstone-backend/internal/gemstone/domain"

type MinerHandler struct {
	service domain.Service
}

func NewMinerHandler(service domain.Service) *MinerHandler {
	return &MinerHandler{service: service}
}
