package handler

import (
	"gemstone-backend/internal/gemstone/global"
	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	userService global.IUserService
}

func (u *UserHandler) VerifyNickname(c *fiber.Ctx) error {
	req := global.VerifyNicknameReq{}
	if err := c.ParamsParser(&req); err != nil {
		return c.JSON(fiber.ErrBadRequest)
	}
	res := u.userService.VerifyNickname(req)
	if !res.Success {
		return c.Status(fiber.StatusBadRequest).JSON("Invalid nickname")
	}
	return c.Status(fiber.StatusOK).SendString("Possible nickname")
}

func (u *UserHandler) Login(c *fiber.Ctx) error {
	req := global.LoginReq{}
	if err := c.BodyParser(&req); err != nil {
		return c.JSON(fiber.ErrBadRequest)
	}
	res := u.userService.Login(req)
	if !res.Success {
		return c.Status(fiber.ErrBadRequest.Code).SendString(fiber.ErrBadRequest.Message)
	}
	return c.Status(fiber.StatusOK).SendString("Welcome!")
}

func (u *UserHandler) Register(c *fiber.Ctx) error {
	req := global.RegisterReq{}
	if err := c.BodyParser(&req); err != nil {
		return c.JSON(fiber.ErrBadRequest)
	}
	res := u.userService.Register(req)
	if !res.Success {
		return c.Status(fiber.ErrBadRequest.Code).SendString(fiber.ErrBadRequest.Message)
	}
	return c.Status(fiber.StatusOK).SendString("Welcome to gemstone")
}

func (u *UserHandler) UpdateNickname(c *fiber.Ctx) error {
	req := global.UpdateNicknameReq{}
	if err := c.BodyParser(&req); err != nil {
		return c.JSON(fiber.ErrBadRequest)
	}
	res := u.userService.UpdateNickname(req)
	if !res.Success {
		return c.Status(fiber.ErrBadRequest.Code).SendString(fiber.ErrBadRequest.Message)
	}
	return c.Status(fiber.StatusOK).SendString("Successfully updated nickname")
}

func NewUserHandler(userService global.IUserService) global.IUserHandler {
	return &UserHandler{userService: userService}
}
