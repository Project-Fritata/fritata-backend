package api

import (
	"github.com/Project-Fritata/fritata-backend/services/auth/core"

	"github.com/gofiber/fiber/v3"
)

func SetupClientRoutes(app *fiber.App) {
	app.Post("/api/v1/auth/register", core.Register)
	app.Post("/api/v1/auth/login", core.Login)
	app.Post("/api/v1/auth/logout", core.Logout)
}
