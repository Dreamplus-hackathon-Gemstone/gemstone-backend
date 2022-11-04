package user

import "github.com/gofiber/fiber/v2"

// SetUserRouter ...
func SetUserRouter(router fiber.Router) {
	router.Get("/user", func(c *fiber.Ctx) error {
		return c.SendString("user")
	})
	router.Post("/user", func(c *fiber.Ctx) error {
		return c.SendString("user")
	})
	router.Patch("/user", func(c *fiber.Ctx) error {
		return c.SendString("user")
	})
	router.Delete("/user", func(c *fiber.Ctx) error {
		return c.SendString("user")
	})
}
