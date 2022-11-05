package handler

import "gemstone-backend/internal/gemstone/domain"

type TokenHandler struct {
	service domain.Service
}

func NewTokenHandler(service domain.Service) *TokenHandler {
	return &TokenHandler{service: service}
}
