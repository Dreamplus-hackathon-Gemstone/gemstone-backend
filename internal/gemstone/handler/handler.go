package handler

import "github.com/gofiber/fiber/v2"

type Handler struct {
	*ItemHandler
	*MakerHandler
	*MinerHandler
	*TokenHandler
	*CategoryHandler
}

func NewHandler(itemHandler *ItemHandler, makerHandler *MakerHandler, minerHandler *MinerHandler, tokenHandler *TokenHandler, categoryHandler *CategoryHandler) *Handler {
	return &Handler{ItemHandler: itemHandler, MakerHandler: makerHandler, MinerHandler: minerHandler, TokenHandler: tokenHandler, CategoryHandler: categoryHandler}
}

func (h *Handler) Ping(c *fiber.Ctx) error {
	return c.SendString("pong")
}
