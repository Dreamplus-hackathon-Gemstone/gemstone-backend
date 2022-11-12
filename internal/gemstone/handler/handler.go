package handler

import (
	"gemstone-backend/internal/gemstone/domain/service"
	"gemstone-backend/internal/gemstone/global"
	"github.com/gofiber/fiber/v2"
)

type Handler struct {
	UserHandler global.IUserHandler
}

func NewHandler(svc *service.Service) *Handler {
	user := NewUserHandler(svc.UserService)
	return &Handler{user}
}

func (h *Handler) Ping(c *fiber.Ctx) error {
	return c.SendString("pong")
}
