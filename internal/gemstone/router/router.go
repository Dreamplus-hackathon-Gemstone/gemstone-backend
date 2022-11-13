package router

import (
	"gemstone-backend/internal/gemstone/handler"
	"github.com/gofiber/fiber/v2"
)

func SetRouter(router fiber.Router, handler *handler.Handler) {
	router.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.Status(fiber.StatusOK).SendString("Gemstone App")
	})

	router.Get("/ping", handler.Ping)

	authRouter := router.Group("/auth")
	authRouter.Get("/verify/:nickname", handler.UserHandler.VerifyNickname)
	authRouter.Post("/register", handler.UserHandler.Register)
	authRouter.Post("/login", handler.UserHandler.Login)
	authRouter.Patch("/update", handler.UserHandler.UpdateNickname)

	proposalRouter := router.Group("/proposal")
	proposalRouter.Get("/", handler.ProposalHandler.FindMany)
	proposalRouter.Get("/:cid", handler.ProposalHandler.Find)
	proposalRouter.Post("/create", handler.ProposalHandler.Store)
	proposalRouter.Patch("/update", handler.ProposalHandler.Update)
	proposalRouter.Delete("/delete", handler.ProposalHandler.Delete)
}
