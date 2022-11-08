package handler

import "gemstone-backend/internal/gemstone/domain"

type TokenHandler struct {
	service domain.IService
}

func NewTokenHandler(service domain.IService) *TokenHandler {
	return &TokenHandler{service: service}
}
