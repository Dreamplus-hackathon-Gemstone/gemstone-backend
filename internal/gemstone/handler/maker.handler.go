package handler

import "gemstone-backend/internal/gemstone/domain"

type MakerHandler struct {
	service domain.IService
}

func NewMakerHandler(service domain.IService) *MakerHandler {
	return &MakerHandler{service: service}
}
