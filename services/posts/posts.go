package main

import (
	"github.com/Project-Fritata/fritata-backend/internal"
	"github.com/Project-Fritata/fritata-backend/services/posts/api"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
)

func main() {
	internal.Connect()
	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://72.144.115.35, http://localhost:5173"},
		AllowCredentials: true,
	}))

	api.SetupClientRoutes(app)

	app.Listen("0.0.0.0:8020")
}
