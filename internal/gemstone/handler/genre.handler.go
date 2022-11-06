package handler

import "gemstone-backend/internal/gemstone/domain"

type GenreHandler struct {
	service domain.Service
}

func NewCategoryHandler(service domain.Service) *GenreHandler {
	return &GenreHandler{service: service}
}
