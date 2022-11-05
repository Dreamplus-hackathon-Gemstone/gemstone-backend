package handler

import "gemstone-backend/internal/gemstone/domain"

type CategoryHandler struct {
	service domain.Service
}

func NewCategoryHandler(service domain.Service) *CategoryHandler {
	return &CategoryHandler{service: service}
}
