package main

import (
	_ "gemstone-backend/docs"
	"gemstone-backend/internal/gemstone"
	"github.com/gofiber/fiber/v2"
	"log"
)

// Bootstrap
// @title          Gemstone Backend
// @version        0.1
// @description    This is API Documentation Gemstone Backend
// @termsOfService http://swagger.io/terms/
// @contact.name   Backend-Tech
// @contact.email  ff.primrose@gmail.com
// @license.name   Apache 2.0
// @license.url    http://www.apache.org/licenses/LICENSE-2.0.html
// @host           127.0.0.1:8080
// @BasePath       /
func main() {

	gemStoneApp := gemstone.NewGemStoneApp(fiber.New())
	log.Println("Create GemStone App ... ")

	gemStoneApp.Run()
}
