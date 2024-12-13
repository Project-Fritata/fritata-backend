package main

import (
	"log"

	"github.com/Project-Fritata/fritata-backend/internal/db"
	"github.com/Project-Fritata/fritata-backend/services/users/api"
	"github.com/Project-Fritata/fritata-backend/services/users/core"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"

	"github.com/Flussen/swagger-fiber-v3"
	_ "github.com/Project-Fritata/fritata-backend/services/users/docs"
)

// @title Fritata Users API
// @version 1.0
// @description This is an API that handles users in Fritata social network
// @description Other microservice APIs: [Auth API](https://20.52.101.8.nip.io/api/v1/swagger/auth/), [Posts API](https://20.52.101.8.nip.io/api/v1/swagger/posts/)

// @contact.name Klemen Remec
// @contact.email klemen.remec@gmail.com

// @license.name CC BY-NC-SA 4.0
// @license.url https://creativecommons.org/licenses/by-nc-sa/4.0/

// @host 20.52.101.8.nip.io
// @BasePath /
func main() {
	db.Connect()

	clientApp := fiber.New()
	clientApp.Use(cors.New(cors.Config{
		AllowOrigins: []string{
			"http://localhost:5173",                                                   // Dev frontend
			"http://localhost:8000", "http://localhost:8010", "http://localhost:8020", // Dev swagger
			"https://project-fritata.github.io", // Prod frontend
			"https://20.52.101.8.nip.io",        // Prod swagger
		},
		AllowCredentials: true,
	}))

	serviceApp := fiber.New()
	serviceApp.Use(cors.New(cors.Config{
		AllowOrigins: []string{
			"http://localhost:5173",                                                   // Dev frontend
			"http://localhost:8000", "http://localhost:8010", "http://localhost:8020", // Dev swagger
			"https://project-fritata.github.io", // Prod frontend
			"https://20.52.101.8.nip.io",        // Prod swagger
		},
		AllowCredentials: true,
	}))

	// Run the service app in a separate goroutine
	go func() {
		api.SetupServiceRoutes(serviceApp)
		serviceApp.Listen("0.0.0.0:8011")
	}()

	// Run the client app in the main goroutine
	api.SetupClientRoutes(clientApp)
	// Swagger
	clientApp.Get("/api/v1/swagger/users/*", swagger.HandlerDefault)
	// Health check
	clientApp.Get("/api/v1/health/users/*", core.Health)

	if err := clientApp.Listen(":8010"); err != nil {
		log.Fatalf("Error starting client users service: %v", err)
	}
}
