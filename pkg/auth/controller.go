package auth

import (
	"strconv"
	"time"

	"github.com/Project-Fritata/fritata-backend/internal"

	"github.com/gofiber/fiber/v3"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

func Register(c fiber.Ctx) error {
	var data map[string]string

	if err := c.Bind().JSON(&data); err != nil {
		return err
	}

	password, _ := bcrypt.GenerateFromPassword([]byte(data["password"]), 14)
	user := internal.User{
		Name:     data["name"],
		Email:    data["email"],
		Password: password,
	}

	internal.DB.Create(&user)

	if user.Id == 0 {
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": "failed to create user",
		})
	}

	return c.JSON(user)
}

func Login(c fiber.Ctx) error {
	var data map[string]string

	if err := c.Bind().JSON(&data); err != nil {
		return err
	}

	var user internal.User
	internal.DB.First(&user, "email = ?", data["email"])

	if user.Id == 0 {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": "invalid credentials",
		})
	}
	if err := bcrypt.CompareHashAndPassword(user.Password, []byte(data["password"])); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": "invalid credentials",
		})
	}

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer:    strconv.Itoa(int(user.Id)),
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(), // 1 day
	})
	token, err := claims.SignedString([]byte(internal.GetEnvVar("JWT_SECRET")))
	if err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": "could not login",
		})
	}

	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    token,
		Expires:  time.Now().Add(time.Hour * 24),
		HTTPOnly: true,
	}
	c.Cookie(&cookie)

	return c.JSON(fiber.Map{
		"message": "success",
	})
}

func Logout(c fiber.Ctx) error {
	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour),
		HTTPOnly: true,
	}

	c.Cookie(&cookie)
	return c.JSON(fiber.Map{
		"message": "success",
	})
}
