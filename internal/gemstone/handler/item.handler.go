package handler

import "gemstone-backend/internal/gemstone/domain"

type ItemHandler struct {
	service domain.Service
}

func NewItemHandler(service domain.Service) *ItemHandler {
	return &ItemHandler{service: service}
}
