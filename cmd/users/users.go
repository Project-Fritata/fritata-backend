package main

import (
	"log"

	"github.com/Project-Fritata/fritata-backend/internal"
	"github.com/Project-Fritata/fritata-backend/pkg/users"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
)

func main() {
	internal.LoadEnv()
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
		users.SetupServiceRoutes(serviceApp)
		if err := serviceApp.Listen("localhost:8011"); err != nil {
			log.Fatalf("Error starting serviceApp: %v", err)
		}
	}()

	// Run the client app in the main goroutine
	users.SetupClientRoutes(clientApp)
	if err := clientApp.Listen("localhost:8010"); err != nil {
		log.Fatalf("Error starting clientApp: %v", err)
	}
}
