package handler

import "github.com/gofiber/fiber/v2"

type Handler struct {
	*ProposalHandler
	*MakerHandler
	*MinerHandler
	*TokenHandler
	*GenreHandler
}

func NewHandler(itemHandler *ProposalHandler, makerHandler *MakerHandler, minerHandler *MinerHandler, tokenHandler *TokenHandler, categoryHandler *GenreHandler) *Handler {
	return &Handler{ProposalHandler: itemHandler, MakerHandler: makerHandler, MinerHandler: minerHandler, TokenHandler: tokenHandler, GenreHandler: categoryHandler}
}

func (h *Handler) Ping(c *fiber.Ctx) error {
	return c.SendString("pong")
}
