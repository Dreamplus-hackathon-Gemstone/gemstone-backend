package global

import "github.com/gofiber/fiber/v2"

type IMinerHandler interface {
	Login(c *fiber.Ctx) error
	Register(c *fiber.Ctx) error
}

type IUserHandler interface {
	VerifyNickname(c *fiber.Ctx) error
	Login(c *fiber.Ctx) error
	Register(c *fiber.Ctx) error
	UpdateNickname(c *fiber.Ctx) error
}

type IMovieHandler interface {
	Find(c *fiber.Ctx) error
	FindMany(c *fiber.Ctx) error
	Store(c *fiber.Ctx) error
}

type IProposalHandler interface {
	Find(c *fiber.Ctx) error
	FindMany(c *fiber.Ctx) error
	Store(c *fiber.Ctx) error
	Update(c *fiber.Ctx) error
	Delete(c *fiber.Ctx) error
}

type IGenreHandler interface {
	FindMovieByGenre(c *fiber.Ctx) error
	FindProposalByGenre(c *fiber.Ctx) error
	Store(c *fiber.Ctx) error
}
