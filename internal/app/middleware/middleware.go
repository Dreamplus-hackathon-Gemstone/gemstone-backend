package middleware

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/csrf"
	"github.com/gofiber/fiber/v2/middleware/logger"
	frecover "github.com/gofiber/fiber/v2/middleware/recover"
)

// NewGemStoneAppMiddleware ...
func NewGemStoneAppMiddleware(app *fiber.App) fiber.Router {
	defer func() {
		if r := recover(); r != nil {
			log.Panic("NewGemStoneAppMiddleware Paniced :", r)
		}
	}()
	logger := app.Use(logger.New(logger.Config{
		Format: "[${ip}]:${port} ${status} - ${method} ${path}\n",
	}))
	csrf := logger.Use(csrf.New())
	api := csrf.Group("/api", cors.New())
	return api.Group("/v1", frecover.New())
}
