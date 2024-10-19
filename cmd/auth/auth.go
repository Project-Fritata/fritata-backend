package main

import (
	"github.com/Project-Fritata/fritata-backend/internal"
	"github.com/Project-Fritata/fritata-backend/pkg/auth"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
)

func main() {
	internal.LoadEnv()
	internal.Connect()
	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowCredentials: true,
	}))

	auth.SetupClientRoutes(app)

	app.Listen("localhost:8000")
}
