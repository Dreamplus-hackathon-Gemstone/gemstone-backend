package handler

import (
	"github.com/gofiber/fiber/v2"
)

type Handler struct {
	AuthHandler *AuthHandler
}

func NewHandler(registerHandler *AuthHandler) *Handler {
	return &Handler{AuthHandler: registerHandler}
}

func (h *Handler) Ping(c *fiber.Ctx) error {
	return c.SendString("pong")
}
