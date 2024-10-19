package auth

import (
	"github.com/gofiber/fiber/v3"
)

func SetupClientRoutes(app *fiber.App) {
	app.Post("/api/v1/auth/register", Register)
	app.Post("/api/v1/auth/login", Login)
	app.Post("/api/v1/auth/logout", Logout)
}
