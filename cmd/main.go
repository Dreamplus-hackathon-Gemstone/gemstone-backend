package main

import (
	"fmt"
	"gemstone-backend/internal/app"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	fmt.Println("Gemstone app start")

	app := app.NewGemStoneApp()
	app.Run()
}

func loadEnv() {
	if err := godotenv.Load(".env"); err != nil {
		log.Panicf("Loading environment failed: %v", err)
	}

}
