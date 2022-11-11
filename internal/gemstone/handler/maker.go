package handler

import (
	"gemstone-backend/internal/gemstone/global"
	"github.com/gofiber/fiber/v2"
)

type MakerHandler struct {
	makerService global.IMakerService
}

func NewMakerHandler(makerService global.IMakerService) global.IMakerHandler {
	return &MakerHandler{makerService: makerService}
}

func (m *MakerHandler) VerifyNickname(c *fiber.Ctx) error {
	req := global.VerifyNicknameReq{}
	if err := c.ParamsParser(&req); err != nil {
		return c.JSON(fiber.ErrBadRequest)
	}
	res := m.makerService.VerifyNickname(req)
	if !res.Success {
		return c.Status(fiber.StatusNotFound).SendString("Possible nickname")
	}
	return c.Status(fiber.StatusOK).SendString("Invalid nickname")
}

func (m *MakerHandler) Login(ctx *fiber.Ctx) error {
	//TODO implement me
	panic("implement me")
}

func (m *MakerHandler) Register(ctx *fiber.Ctx) error {
	//TODO implement me
	panic("implement me")
}

func (m *MakerHandler) UpdateNickname(ctx *fiber.Ctx) error {
	//TODO implement me
	panic("implement me")
}
