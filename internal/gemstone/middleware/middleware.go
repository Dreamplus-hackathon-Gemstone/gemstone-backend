package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/csrf"
	"github.com/gofiber/fiber/v2/middleware/logger"
	fRecover "github.com/gofiber/fiber/v2/middleware/recover"
)

func NewGemStoneAppMiddleware(router *fiber.App) fiber.Router {
	fLogger := router.Use(logger.New(logger.Config{
		Format: "[${ip}]:${port} ${status} - ${method} ${path}\n",
	}))
	fCsrf := fLogger.Use(csrf.New())
	api := fCsrf.Group("/api", cors.New())
	return api.Group("/v1", fRecover.New())
}
