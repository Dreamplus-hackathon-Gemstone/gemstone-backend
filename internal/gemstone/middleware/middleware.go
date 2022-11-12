package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	fRecover "github.com/gofiber/fiber/v2/middleware/recover"
)

func NewGemStoneAppMiddleware(router *fiber.App) fiber.Router {
	fLogger := router.Use(logger.New(logger.Config{
		Format: "[${ip}]:${adapter} ${status} - ${method} ${path}\n",
	}))
	fCors := fLogger.Use(cors.New())
	api := fCors.Group("/api")
	return api.Group("/v1", fRecover.New())
}
