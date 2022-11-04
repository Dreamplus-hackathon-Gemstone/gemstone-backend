package main

import (
	"fmt"
	"gemstone-backend/internal/app"
)

func main() {
	fmt.Println("Gemstone app start")

	app := app.NewGemStoneApp()
	app.Run()
}
