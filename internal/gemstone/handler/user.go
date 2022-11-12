package handler

import (
	"fmt"
	"gemstone-backend/internal/gemstone/global"
	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	userService global.IUserService
}

// VerifyNickname godoc
// @Title       VerifyNickname
// @Summary     Verify User Nickname
// @Description 존재하는 닉네임인지 확인.
// @Tags        auth
// @Accept      json
// @Param       nickname path string true "VerifyNickname Req"
// @Produce     json
// @Success     200 {object} global.VerifyNicknameRes
// @Failure     400 {object} global.VerifyNicknameRes
// @Router      /api/v1/auth/verify/{nickname} [get]
func (u *UserHandler) VerifyNickname(c *fiber.Ctx) error {
	req := global.VerifyNicknameReq{}
	if err := c.ParamsParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(global.VerifyNicknameRes{Success: false, Err: fmt.Errorf("%v", fiber.ErrBadRequest.Message)})
	}
	res := u.userService.VerifyNickname(req)
	if res.Success == false {
		return c.Status(fiber.StatusBadRequest).JSON(res)
	}
	return c.Status(fiber.StatusOK).JSON(res)
}

// Login godoc
// @Title       Login
// @Summary     Login
// @Description 로그인
// @Tags        auth
// @Accept      json
// @Param       LoginReq body global.LoginReq true "Login Req"
// @Produce     json
// @Success     200 {object} global.LoginRes
// @Failure     400 {object} global.LoginRes
// @Router      /api/v1/auth/login [post]
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

// Register godoc
// @Title       Register
// @Summary     Register
// @Description 회원가입
// @Tags        auth
// @Accept      json
// @Param       RegisterReq body global.RegisterReq true "Register Req"
// @Produce     json
// @Success     200 {object} global.RegisterRes
// @Failure     400 {object} global.RegisterRes
// @Router      /api/v1/auth/register [post]
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

// UpdateNickname godoc
// @Title       UpdateNickname
// @Summary     UpdateNickname
// @Description 닉네임 수정
// @Tags        auth
// @Accept      json
// @Param       UpdateNicknameReq body global.UpdateNicknameReq true "UpdateNickname Req"
// @Produce     json
// @Success     200 {object} global.UpdateNicknameRes
// @Failure     400 {object} global.UpdateNicknameRes
// @Router      /api/v1/auth/update [patch]
func (u *UserHandler) UpdateNickname(c *fiber.Ctx) error {
	req := global.UpdateNicknameReq{}
	if err := c.BodyParser(&req); err != nil {
		return c.JSON(fiber.ErrBadRequest)
	}
	res := u.userService.UpdateNickname(req)
	if !res.Success {
		return c.Status(fiber.ErrBadRequest.Code).SendString(fiber.ErrBadRequest.Message)
	}
	return c.Status(fiber.StatusOK).SendString("Nickname has successfully updated")
}

func NewUserHandler(userService global.IUserService) global.IUserHandler {
	return &UserHandler{userService: userService}
}
