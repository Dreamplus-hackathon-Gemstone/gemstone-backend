package handler

import (
	"gemstone-backend/internal/gemstone/domain/service"
	"gemstone-backend/internal/gemstone/global"
	"github.com/gofiber/fiber/v2"
)

type Handler struct {
	MakerHandler global.IMakerHandler
	MinerHandler global.IMinerHandler
}

func NewHandler(service *service.Service) *Handler {
	maker := NewMakerHandler(service.MakerService)
	miner := NewMinerHandler(service.MinerService)
	return &Handler{maker, miner}
}

func (h *Handler) Ping(c *fiber.Ctx) error {
	return c.SendString("pong")
}
