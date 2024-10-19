package posts

import (
	"github.com/gofiber/fiber/v3"
)

func SetupClientRoutes(app *fiber.App) {
	app.Get("/api/v1/posts", GetPosts)
	app.Post("/api/v1/posts", CreatePost)
}
