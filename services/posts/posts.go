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
		AllowOrigins:     []string{"http://localhost:5173", "https://project-fritata.github.io"},
		AllowCredentials: true,
	}))

	api.SetupClientRoutes(app)

	app.Listen("0.0.0.0:8020")
}
