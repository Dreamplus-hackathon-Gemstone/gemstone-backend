package router

import (
	"gemstone-backend/internal/app/router/user"

	"github.com/gofiber/fiber/v2"
)

// SetRouter sets the router
func SetRouter(router fiber.Router) error {
	user.SetUserRouter(router)
	// token.SetTokenRouter(app)
	// item.SetItemRouter(app)
	return nil
}
