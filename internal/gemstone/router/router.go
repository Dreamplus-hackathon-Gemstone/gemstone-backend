package router

import (
	"gemstone-backend/internal/gemstone/handler"
	"github.com/gofiber/fiber/v2"
)

func SetRouter(router fiber.Router, handler *handler.Handler) {
	router.Get("/ping", handler.Ping)

	authRouter := router.Group("/auth")
	authRouter.Get("/verify/:nickname", handler.UserHandler.VerifyNickname)
	authRouter.Post("/register", handler.UserHandler.Register)
	authRouter.Post("/login", handler.UserHandler.Login)
	authRouter.Patch("/update", handler.UserHandler.UpdateNickname)
}
