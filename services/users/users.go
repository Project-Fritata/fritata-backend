package main

import (
	"log"

	"github.com/Project-Fritata/fritata-backend/internal"
	"github.com/Project-Fritata/fritata-backend/services/users/api"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
)

func main() {
	internal.Connect()

	clientApp := fiber.New()
	clientApp.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowCredentials: true,
	}))

	serviceApp := fiber.New()
	serviceApp.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowCredentials: true,
	}))

	// Run the service app in a separate goroutine
	go func() {
		api.SetupServiceRoutes(serviceApp)
		if err := serviceApp.Listen(":8011"); err != nil {
			log.Fatalf("Error starting serviceApp: %v", err)
		}
	}()

	// Run the client app in the main goroutine
	api.SetupClientRoutes(clientApp)
	if err := clientApp.Listen(":8010"); err != nil {
		log.Fatalf("Error starting clientApp: %v", err)
	}
}
