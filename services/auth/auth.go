package main

import (
	"github.com/Project-Fritata/fritata-backend/internal"
	"github.com/Project-Fritata/fritata-backend/services/auth/api"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
)

func main() {
	internal.Connect()
	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"72.144.115.35"},
		AllowCredentials: true,
	}))

	api.SetupClientRoutes(app)

	app.Listen("0.0.0.0:8000")
}
