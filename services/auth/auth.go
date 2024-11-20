package main

import (
	"log"

	"github.com/Project-Fritata/fritata-backend/internal/db"
	"github.com/Project-Fritata/fritata-backend/services/auth/api"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"

	"github.com/Flussen/swagger-fiber-v3"
	_ "github.com/Project-Fritata/fritata-backend/services/auth/docs"
)

// @title Fritata Auth API
// @version 1.0
// @description This is an API that handles auth in Fritata social network
// @description Other microservice APIs: [Posts API](https://50.17.141.235.nip.io/api/v1/posts/swagger/), [Users API](https://50.17.141.235.nip.io/api/v1/users/swagger/)

// @contact.name Klemen Remec
// @contact.email klemen.remec@gmail.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @schemes http
// @host localhost:8000
// @BasePath /
func main() {
	db.Connect()
	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173", "https://project-fritata.github.io", "http://localhost:8000", "http://localhost:8010", "http://localhost:8020"},
		AllowCredentials: true,
	}))

	api.SetupClientRoutes(app)
	app.Get("/api/v1/auth/swagger/*", swagger.HandlerDefault)

	if err := app.Listen(":8000"); err != nil {
		log.Fatalf("Error starting auth service: %v", err)
	}
}
