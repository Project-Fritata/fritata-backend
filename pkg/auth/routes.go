package auth

import (
	"github.com/gofiber/fiber/v3"
)

func SetupRoutes(app *fiber.App) {
	app.Post("/api/v1/register", Register)
	app.Post("/api/v1/login", Login)
}
