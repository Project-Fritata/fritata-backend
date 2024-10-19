package users

import (
	"github.com/gofiber/fiber/v3"
)

func SetupClientRoutes(app *fiber.App) {
	app.Get("/api/v1/users/:username", GetUserByUsername)
	app.Put("/api/v1/users", UpdateUser)
}

func SetupServiceRoutes(app *fiber.App) {
	app.Get("/api/v1/users/:id", GetUserById)
	app.Post("/api/v1/users", CreateUser)
}
