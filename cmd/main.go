package main

import (
	"fmt"
	"gemstone-backend/internal/gemstone"
	"gemstone-backend/internal/gemstone/config"
	"log"
)

func main() {
	fmt.Println("Gemstone app start")

	gemStoneApp := gemstone.NewGemStoneApp()
	if err := config.LoadEnv(); err != nil {
		log.Panicf("Panic occured while loading environment")
	}
	config := config.NewConfig()

	//gemStoneApp.Run()
}
