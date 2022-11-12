package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	fRecover "github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/swagger"
)

func NewGemStoneAppMiddleware(router *fiber.App) fiber.Router {

	router.Get("/swagger/*", swagger.HandlerDefault)
	router.Get("/swagger/*", swagger.New(swagger.Config{ // custom
		URL:         "http://example.com/doc.json",
		DeepLinking: false,
		// Expand ("list") or Collapse ("none") tag groups by default
		DocExpansion: "none",
		// Prefill OAuth ClientId on Authorize popup
		OAuth: &swagger.OAuthConfig{
			AppName:  "OAuth Provider",
			ClientId: "21bb4edc-05a7-4afc-86f1-2e151e4ba6e2",
		},
		// Ability to change OAuth2 redirect uri location
		OAuth2RedirectUrl: "http://localhost:8080/swagger/oauth2-redirect.html",
	}))

	fLogger := router.Use(logger.New(logger.Config{
		Format: "[${ip}]:${adapter} ${status} - ${method} ${path}\n",
	}))
	fCors := fLogger.Use(cors.New())
	api := fCors.Group("/api", fRecover.New())
	return api.Group("/v1")
}
