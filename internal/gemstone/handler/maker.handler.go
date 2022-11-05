package handler

import "gemstone-backend/internal/gemstone/domain"

type MakerHandler struct {
	service domain.Service
}

func NewMakerHandler(service domain.Service) *MakerHandler {
	return &MakerHandler{service: service}
}
