package users

import (
	"github.com/gofiber/fiber/v3"
)

func SetupRoutes(app *fiber.App) {
	app.Get("/api/v1/users", GetUser)
}
