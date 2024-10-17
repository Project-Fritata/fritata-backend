package main

import (
	"github.com/Project-Fritata/fritata-backend/internal"
	"github.com/Project-Fritata/fritata-backend/pkg/auth"

	"github.com/gofiber/fiber/v3"
)

func main() {
	internal.LoadEnv()
	internal.Connect()
	app := fiber.New()

	auth.SetupRoutes(app)

	app.Listen(":8000")
}
