package app

import (
	"context"
	appctx "gemstone-backend/internal/app/context"
	"gemstone-backend/internal/app/middleware"
	"gemstone-backend/internal/app/router"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

// IGemStoneApp implements GemStoneApp interface
type IGemStoneApp interface {
	Run()
}

// GemStoneApp implements GemStoneApp
type GemStoneApp struct {
	fiber  *fiber.App
	appCtx appctx.AppCtx
}

// NewGemStoneApp creates a new GemStoneApp
func NewGemStoneApp() IGemStoneApp {
	return &GemStoneApp{fiber: fiber.New()}
}

// HealthCheck instant
func HealthCheck(c *fiber.Ctx) error {
	return c.SendString("pong")
}

func loadEnv() error {
	return godotenv.Load(".env")
}

// Run GemStoneApp
func (app *GemStoneApp) Run() {
	// Set ENV
	if err := loadEnv(); err != nil {
		log.Panicf("Load environment failed with error: %v", err)
	}
	// Set Middleware
	v1 := middleware.NewGemStoneAppMiddleware(app.fiber)
	log.Println("Set Middleware Done")
	// Set Router
	router.SetRouter(v1)
	log.Println("Set Router Done")

	appCtx, err := appctx.Initialize(appctx.CreateConfigFromEnv())
	if err != nil {
		log.Panicf("Initialize failed: %v", err)
	}

	app.appCtx = appCtx

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
