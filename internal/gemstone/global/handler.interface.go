package global

import "github.com/gofiber/fiber/v2"

type IMinerHandler interface {
	Login(ctx *fiber.Ctx) error
	Register(ctx *fiber.Ctx) error
}

type IMakerHandler interface {
	VerifyNickname(ctx *fiber.Ctx) error
	Login(ctx *fiber.Ctx) error
	Register(ctx *fiber.Ctx) error
	UpdateNickname(ctx *fiber.Ctx) error
}

type IMovieHandler interface {
	Find(ctx *fiber.Ctx) error
	FindMany(ctx *fiber.Ctx) error
	Store(ctx *fiber.Ctx) error
	Update(ctx *fiber.Ctx) error
	Delete(ctx *fiber.Ctx) error
}

type IProposalHandler interface {
	Find(ctx *fiber.Ctx) error
	FindMany(ctx *fiber.Ctx) error
	Store(ctx *fiber.Ctx) error
	Update(ctx *fiber.Ctx) error
	Delete(ctx *fiber.Ctx) error
}
