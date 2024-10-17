package auth

import (
	"github.com/Project-Fritata/fritata-backend/internal"

	"github.com/gofiber/fiber/v3"
	"golang.org/x/crypto/bcrypt"
)

var bcrypt_password = internal.GetEnvVar("BCRYPT_PASSWORD")

func Register(c fiber.Ctx) error {
	var data map[string]string

	if err := c.Bind().JSON(&data); err != nil {
		return err
	}

	password, _ := bcrypt.GenerateFromPassword([]byte(data[bcrypt_password]), 14)
	user := internal.User{
		Name:     data["name"],
		Email:    data["email"],
		Password: password,
	}

	internal.DB.Create(&user)

	return c.JSON(user)
}

func Login(c fiber.Ctx) error {
	var data map[string]string

	if err := c.Bind().JSON(&data); err != nil {
		return err
	}

	var user internal.User
	internal.DB.First(&user, "email = ?", data["email"])

	if user.ID == 0 {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": "invalid credentials",
		})
	}

	if err := bcrypt.CompareHashAndPassword(user.Password, []byte(data[bcrypt_password])); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": "invalid credentials",
		})
	}

	return c.JSON(user)
}
