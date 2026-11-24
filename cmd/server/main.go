package main

import (
	"log"

	"github.com/gofiber/fiber/v3"
	"github.com/mohammadReza-kaspour/Complete-RBAC-System/configs"

	recoverer "github.com/gofiber/fiber/v3/middleware/recover"
)

func main() {
	cfg := configs.Load()

	app := fiber.New(*cfg.FiberConfig)
	app.Use(recoverer.New())
	log.Fatal(app.Listen(":5000"))
}
