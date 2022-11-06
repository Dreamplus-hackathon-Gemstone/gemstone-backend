package main

import (
	"gemstone-backend/internal/gemstone"
	"github.com/gofiber/fiber/v2"
	"log"
)

func main() {

	gemStoneApp := gemstone.NewGemStoneApp(fiber.New())
	log.Println("Create GemStone App ... ")

	gemStoneApp.Run()
}
