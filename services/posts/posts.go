package main

import (
	"log"

	"github.com/Project-Fritata/fritata-backend/internal/db"
	"github.com/Project-Fritata/fritata-backend/services/posts/api"
	"github.com/Project-Fritata/fritata-backend/services/posts/core"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"

	"github.com/Flussen/swagger-fiber-v3"
	_ "github.com/Project-Fritata/fritata-backend/services/posts/docs"
)

// @title Fritata Posts API
// @version 1.0
// @description This is an API that handles posts in Fritata social network
// @description Other microservice APIs: [Auth API](https://20.52.101.8.nip.io/api/v1/swagger/auth/), [Users API](https://20.52.101.8.nip.io/api/v1/swagger/users/)

// @contact.name Klemen Remec
// @contact.email klemen.remec@gmail.com

// @license.name Apache 2.0
// @license.url https://creativecommons.org/licenses/by-nc-sa/4.0/

// @host 20.52.101.8.nip.io
// @BasePath /
func main() {
	db.Connect()
	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowOrigins: []string{
			"http://localhost:5173",                                                   // Dev frontend
			"http://localhost:8000", "http://localhost:8010", "http://localhost:8020", // Dev swagger
			"https://project-fritata.github.io", // Prod frontend
			"https://20.52.101.8.nip.io",        // Prod swagger
		},
		AllowCredentials: true,
	}))

	api.SetupClientRoutes(app)
	// Swagger
	app.Get("/api/v1/swagger/posts/*", swagger.HandlerDefault)
	// Health check
	app.Get("/api/v1/health/posts/*", core.Health)

	if err := app.Listen(":8020"); err != nil {
		log.Fatalf("Error starting posts service: %v", err)
	}
}
