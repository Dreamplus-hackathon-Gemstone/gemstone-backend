package handler

import "gemstone-backend/internal/gemstone/domain"

type MinerHandler struct {
	service domain.IService
}

func NewMinerHandler(service domain.IService) *MinerHandler {
	return &MinerHandler{service: service}
}
