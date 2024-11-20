package main

import (
	"log"

	"github.com/Project-Fritata/fritata-backend/internal/db"
	"github.com/Project-Fritata/fritata-backend/services/users/api"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"

	"github.com/Flussen/swagger-fiber-v3"
	_ "github.com/Project-Fritata/fritata-backend/services/users/docs"
)

// @title Fritata Users API
// @version 1.0
// @description This is an API that handles users in Fritata social network
// @description Other microservice APIs: [Auth API](https://50.17.141.235.nip.io/api/v1/auth/swagger/), [Posts API](https://50.17.141.235.nip.io/api/v1/posts/swagger/)

// @contact.name Klemen Remec
// @contact.email klemen.remec@gmail.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @schemes http
// @host localhost:8010
// @BasePath /
func main() {
	db.Connect()

	clientApp := fiber.New()
	clientApp.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173", "https://project-fritata.github.io", "http://localhost:8000", "http://localhost:8010", "http://localhost:8020"},
		AllowCredentials: true,
	}))

	serviceApp := fiber.New()
	serviceApp.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173", "https://project-fritata.github.io", "http://localhost:8000", "http://localhost:8010", "http://localhost:8020"},
		AllowCredentials: true,
	}))

	// Run the service app in a separate goroutine
	go func() {
		api.SetupServiceRoutes(serviceApp)
		serviceApp.Listen("0.0.0.0:8011")
	}()

	// Run the client app in the main goroutine
	api.SetupClientRoutes(clientApp)
	clientApp.Get("/swagger/*", swagger.HandlerDefault)

	if err := clientApp.Listen(":8010"); err != nil {
		log.Fatalf("Error starting client users service: %v", err)
	}
}
