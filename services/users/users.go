package main

import (
	"github.com/Project-Fritata/fritata-backend/internal"
	"github.com/Project-Fritata/fritata-backend/services/users/api"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
)

func main() {
	internal.Connect()

	clientApp := fiber.New()
	clientApp.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://127.0.0.1"},
		AllowCredentials: true,
	}))

	serviceApp := fiber.New()
	serviceApp.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://127.0.0.1"},
		AllowCredentials: true,
	}))

	// Run the service app in a separate goroutine
	go func() {
		api.SetupServiceRoutes(serviceApp)
		serviceApp.Listen("0.0.0.0:8011")
	}()

	// Run the client app in the main goroutine
	api.SetupClientRoutes(clientApp)
	clientApp.Listen("0.0.0.0:8010")
}
