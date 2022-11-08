package handler

import "gemstone-backend/internal/gemstone/domain"

type GenreHandler struct {
	service domain.IService
}

func NewCategoryHandler(service domain.IService) *GenreHandler {
	return &GenreHandler{service: service}
}
