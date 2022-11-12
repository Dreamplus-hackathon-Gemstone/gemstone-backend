package gemstone

import (
	"context"
	"gemstone-backend/internal/gemstone/config"
	"gemstone-backend/internal/gemstone/domain"
	"gemstone-backend/internal/gemstone/domain/repository"
	"gemstone-backend/internal/gemstone/domain/service"
	"gemstone-backend/internal/gemstone/handler"
	"gemstone-backend/internal/gemstone/middleware"
	"gemstone-backend/internal/gemstone/router"
	"github.com/gofiber/fiber/v2"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type App struct {
	router *fiber.App
}

func NewGemStoneApp(router *fiber.App) *App {
	return &App{router: router}
}

func (app *App) Run() {
	// Load Configs ...
	if err := config.LoadEnv(); err != nil {
		log.Panicf("Panic occured while loading environment")
	}
	cfg := config.NewConfig()
	log.Println("Config Done!")

	//Get Transactor struct.
	//It includes gorm.DB with access permission for every table.
	db := domain.NewClient(cfg)
	log.Println("Create GORM Client")

	// Get Domain Repository
	newRepository := repository.NewRepository(db)
	log.Println("Create Repository ... ", newRepository)

	// Get Domain Service
	svc := service.NewService(newRepository)
	log.Println("Create Service ... ")

	// Set Handler ...
	h := handler.NewHandler(svc)
	log.Println("Set Handler ...")

	// Create Fiber App

	// Add middleware for router
	r := middleware.NewGemStoneAppMiddleware(app.router)
	log.Println("Applying middleware ... ")

	router.SetRouter(r, h)
	log.Println("Setting router ... ")

	if err := app.router.Listen(":8080"); err != nil {
		log.Panicf("Listen on %s, paniced: %v", "8080", err)
	}

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")

	_, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := app.router.Shutdown(); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}
}
