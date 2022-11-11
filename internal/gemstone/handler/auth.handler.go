package handler

import (
	"gemstone-backend/internal/gemstone/domain/service"
	"gemstone-backend/internal/gemstone/global"
	"github.com/gofiber/fiber/v2"
)

type IAuthHandler interface {
	VerifyNickname(c *fiber.Ctx) error
	RegisterMaker(c *fiber.Ctx) error
	RegisterMiner(c *fiber.Ctx) error
	Login(c *fiber.Ctx) error
}

type AuthHandler[P global.Parameter, R global.ReturnType] struct {
	makerService service.MakerService[P, R]
	minerService service.MinerService[P, R]
}

func (a *AuthHandler[P, R]) VerifyNickname(c *fiber.Ctx) error {
	//TODO implement me
	panic("implement me")
}

func (a *AuthHandler[P, R]) RegisterMaker(c *fiber.Ctx) error {
	//TODO implement me
	panic("implement me")
}

func (a *AuthHandler[P, R]) RegisterMiner(c *fiber.Ctx) error {
	//TODO implement me
	panic("implement me")
}

func (a *AuthHandler[P, R]) Login(c *fiber.Ctx) error {
	//TODO implement me
	panic("implement me")
}

func NewAuthHandler[P global.Parameter, R global.ReturnType](makerService service.MakerService[P, R], minerService service.MinerService[P, R]) IAuthHandler {
	return &AuthHandler[P, R]{makerService: makerService, minerService: minerService}
}
