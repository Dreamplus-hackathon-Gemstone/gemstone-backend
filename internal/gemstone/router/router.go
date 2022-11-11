package router

import (
	"gemstone-backend/internal/gemstone/handler"
	"github.com/gofiber/fiber/v2"
)

func SetRouter(router fiber.Router, handler *handler.Handler) {
	router.Get("/ping", handler.Ping)

	minerRouter := router.Group("/miner")
	minerRouter.Post("/login", handler.MinerHandler.Login)
	minerRouter.Post("/register", handler.MinerHandler.Register)

	makerRouter := router.Group("/maker")
	makerRouter.Get("/verify/:nickname", handler.MakerHandler.VerifyNickname)
	makerRouter.Post("/login", handler.MakerHandler.Login)
	makerRouter.Post("/register", handler.MakerHandler.Register)
	makerRouter.Patch("/update", handler.MakerHandler.UpdateNickname)
}
