package api

import (
	"github.com/Project-Fritata/fritata-backend/services/users/core"

	"github.com/gofiber/fiber/v3"
)

func SetupClientRoutes(app *fiber.App) {
	app.Get("/api/v1/users/:username", core.GetUserByUsername)
	app.Get("/api/v1/users", core.GetUserByAuth)
	app.Put("/api/v1/users", core.UpdateUser)
}

func SetupServiceRoutes(app *fiber.App) {
	app.Get("/api/v1/users/:id", core.GetUserById)
	app.Post("/api/v1/users", core.CreateUser)
}
