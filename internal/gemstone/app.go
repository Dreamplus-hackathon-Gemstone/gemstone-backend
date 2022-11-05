package gemstone

import "github.com/gofiber/fiber/v2"

type App struct {
	Router *fiber.App
}

func NewGemStoneApp(router *fiber.App) *App {
	return &App{Router: router}
}
