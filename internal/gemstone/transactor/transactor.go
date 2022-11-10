package transactor

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type ITransactor interface {
	GetTx() *gorm.DB
	Execute(ctx *fiber.Ctx, run func(c *fiber.Ctx) error) error
	ExecuteTx(ctx *fiber.Ctx, run func(ctx *fiber.Ctx) error) error
}
