package main

import (
	"context"
	"gemstone-backend/internal/gemstone"
	"gemstone-backend/internal/gemstone/config"
	"gemstone-backend/internal/gemstone/handler"
	"gemstone-backend/internal/gemstone/middleware"
	"gemstone-backend/internal/gemstone/repository"
	"gemstone-backend/internal/gemstone/router"
	"gemstone-backend/internal/gemstone/service"
	"github.com/gofiber/fiber/v2"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {

	// Load Configs ...
	if err := config.LoadEnv(); err != nil {
		log.Panicf("Panic occured while loading environment")
	}
	cfg := config.NewConfig()
	log.Println("Config Done!")

	//Get Transactor struct.
	//It includes gorm.DB with access permission for every table.
	db := repository.NewTransactor(repository.NewClient(cfg))
	log.Println("Create Transactor ... ")

	// Get Repository
	newRepository := repository.NewRepository(db.Client)
	log.Println("Create Repository ... ")

	// Get Service
	svc := service.NewService(newRepository)
	log.Println("Create Service ... ")

	// Get Handelr
	h := handler.NewHandler(handler.NewItemHandler(svc.ItemService), handler.NewMakerHandler(svc.MakerService), handler.NewMinerHandler(svc.MinerService), handler.NewTokenHandler(svc.TokenService), handler.NewCategoryHandler(svc.CategoryService))
	log.Println("Create Handler ... ")

	// Create Fiber App
	gemStoneApp := gemstone.NewGemStoneApp(fiber.New())
	log.Println("Create GemStone App ... ")

	// Add middleware for router
	r := middleware.NewGemStoneAppMiddleware(gemStoneApp.Router)
	log.Println("Applying middleware ... ")
	router.SetRouter(r, h)
	log.Println("Setting router ... ")

	if err := gemStoneApp.Router.Listen(":8080"); err != nil {
		log.Panicf("Listen on %s, paniced: %v", "8080", err)
	}

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")

	_, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := gemStoneApp.Router.Shutdown(); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}
}
