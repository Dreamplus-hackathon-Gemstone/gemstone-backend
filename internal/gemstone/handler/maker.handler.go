package handler

import (
	"gemstone-backend/internal/gemstone/domain"
	"github.com/gofiber/fiber/v2"
)

type MakerHandler struct {
	service domain.IService
}

func NewMakerHandler(service domain.IService) *MakerHandler {
	return &MakerHandler{service: service}
}

func (h *MakerHandler) GetLoginRequestID(c *fiber.Ctx) error {
	h.service.Find()
	return c.SendString("LoginRequestID")
}
