package api

import (
	"github.com/Project-Fritata/fritata-backend/services/posts/core"

	"github.com/gofiber/fiber/v3"
)

func SetupClientRoutes(app *fiber.App) {
	app.Get("/api/v1/posts", core.GetPosts)
	app.Post("/api/v1/posts", core.CreatePost)
}
