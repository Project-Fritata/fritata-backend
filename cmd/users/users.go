package main

import (
	"github.com/Project-Fritata/fritata-backend/internal"
	"github.com/Project-Fritata/fritata-backend/pkg/users"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
)

func main() {
	internal.LoadEnv()
	internal.Connect()
	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost"},
		AllowCredentials: true,
	}))

	users.SetupRoutes(app)

	app.Listen(":8001")
}
