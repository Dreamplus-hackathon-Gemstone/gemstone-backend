package database

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// ITransactor ...
type ITransactor interface {
	Execute(ctx *fiber.Ctx, run func(tx *gorm.DB, ctx *fiber.Ctx) error) error
	ExecuteContext(ctx *fiber.Ctx, tx *gorm.DB, run func(tx *gorm.DB, ctx *fiber.Ctx) error) error
	GetTx() *gorm.DB
}

// Transactor ...
type Transactor struct {
	db *gorm.DB
}

// NewTransactor ...
func NewTransactor(db *gorm.DB) ITransactor {
	return &Transactor{db: db}
}

// GetTx ...
func (t *Transactor) GetTx() *gorm.DB {
	return t.db.Begin()
}

// ExecuteContext ...
func (t *Transactor) ExecuteContext(ctx *fiber.Ctx, tx *gorm.DB, run func(tx *gorm.DB, ctx *fiber.Ctx) error) error {
	return run(tx, ctx)
}

// Execute ...
func (t *Transactor) Execute(ctx *fiber.Ctx, run func(tx *gorm.DB, ctx *fiber.Ctx) error) error {
	tx := t.GetTx()
	return t.ExecuteContext(ctx, tx, run)
}
