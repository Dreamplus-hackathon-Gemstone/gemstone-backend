package app

import (
	"context"
	"gemstone-backend/internal/app/middleware"
	"gemstone-backend/internal/app/router"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gofiber/fiber/v2"
)

// IGemStoneApp implements GemStoneApp interface
type IGemStoneApp interface {
	Run()
}

// GemStoneApp implements GemStoneApp
type GemStoneApp struct {
	fiber *fiber.App
}

// NewGemStoneApp creates a new GemStoneApp
func NewGemStoneApp() IGemStoneApp {
	return &GemStoneApp{fiber: fiber.New()}
}

// HealthCheck instant
func HealthCheck(c *fiber.Ctx) error {
	return c.SendString("pong")
}

// Run GemStoneApp
func (app *GemStoneApp) Run() {
	v1 := middleware.NewGemStoneAppMiddleware(app.fiber)
	log.Println("Setting middleware is done")
	router.SetRouter(v1)

	if err := app.fiber.Listen(":8080"); err != nil {
		log.Panicf("Listen on %s, paniced: %v", "8080", err)
	}

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")

	_, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := app.fiber.Shutdown(); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}
}
