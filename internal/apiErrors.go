package internal

import "github.com/gofiber/fiber/v3"

// func InternalServerError(c fiber.Ctx, err error) error {
// 	c.Status(fiber.StatusInternalServerError)
// 	message := err.Error()
// 	if err == nil {
// 		message = "internal server error"
// 	} else {
// 		message = "internal server error: " + err.Error()
// 	}
// 	return c.JSON(fiber.Map{
// 		"message": message,
// 	})
// }

func InternalServerError(c fiber.Ctx) error {
	c.Status(fiber.StatusInternalServerError)
	return c.JSON(fiber.Map{
		"message": "internal server error",
	})
}

func InvalidRequest(c fiber.Ctx) error {
	c.Status(fiber.StatusBadRequest)
	return c.JSON(fiber.Map{
		"message": "invalid request",
	})
}

func InvalidCredentials(c fiber.Ctx) error {
	c.Status(fiber.StatusBadRequest)
	return c.JSON(fiber.Map{
		"message": "invalid credentials",
	})
}

func UserNotFound(c fiber.Ctx) error {
	c.Status(fiber.StatusNotFound)
	return c.JSON(fiber.Map{
		"message": "user not found",
	})
}

func Unauthenticated(c fiber.Ctx) error {
	c.Status(fiber.StatusUnauthorized)
	return c.JSON(fiber.Map{
		"message": "unauthenticated",
	})
}

func UnprocessableEntity(c fiber.Ctx) error {
	c.Status(fiber.StatusUnprocessableEntity)
	return c.JSON(fiber.Map{
		"message": "bad request content",
	})
}
