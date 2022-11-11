package handler

import (
	"gemstone-backend/internal/gemstone/global"
	"github.com/gofiber/fiber/v2"
)

type MinerHandler struct {
	minerService global.IMinerService
}

func NewMinerHandler(minerService global.IMinerService) global.IMinerHandler {
	return &MinerHandler{minerService: minerService}
}

func (m *MinerHandler) Login(c *fiber.Ctx) error {
	//TODO implement me
	panic("implement me")
}

func (m *MinerHandler) Register(c *fiber.Ctx) error {
	//TODO implement me
	panic("implement me")
}
