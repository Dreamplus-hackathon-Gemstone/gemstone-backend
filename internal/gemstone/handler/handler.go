package handler

import (
	"gemstone-backend/internal/gemstone/domain/service"
	"gemstone-backend/internal/gemstone/global"
	"github.com/gofiber/fiber/v2"
)

type Handler struct {
	UserHandler     global.IUserHandler
	ProposalHandler global.IProposalHandler
}

func NewHandler(svc *service.Service) *Handler {
	user := NewUserHandler(svc.UserService)
	proposal := NewProposalHandler(svc.ProposalService)
	return &Handler{
		UserHandler:     user,
		ProposalHandler: proposal,
	}
}

func (h *Handler) Ping(c *fiber.Ctx) error {
	return c.SendString("pong")
}
