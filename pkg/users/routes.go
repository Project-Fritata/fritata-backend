package users

import (
	"github.com/gofiber/fiber/v3"
)

func SetupRoutes(app *fiber.App) {
	app.Get("/api/v1/users", GetUser)
	app.Post("/api/v1/users", CreateUser)
	app.Put("/api/v1/users", UpdateUser)
}
