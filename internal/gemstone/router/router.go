package router

import (
	"gemstone-backend/internal/gemstone/handler"
	"github.com/gofiber/fiber/v2"
)

func SetRouter(router fiber.Router, handler *handler.Handler) {
	router.Get("/ping", handler.Ping)
}