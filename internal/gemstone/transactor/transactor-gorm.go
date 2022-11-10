package transactor

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type Transactor struct {
	Client *gorm.DB
}

func NewTransactor(client *gorm.DB) *Transactor {
	return &Transactor{Client: client}
}

func (t *Transactor) GetTx() *gorm.DB {
	return t.Client
}

func (t *Transactor) Execute(ctx *fiber.Ctx, run func(c *fiber.Ctx) error) error {
	return t.ExecuteTx(ctx, run)
}

func (t *Transactor) ExecuteTx(ctx *fiber.Ctx, run func(ctx *fiber.Ctx) error) error {
	return t.Client.Transaction(func(tx *gorm.DB) error {
		return run(ctx)
	})
}
