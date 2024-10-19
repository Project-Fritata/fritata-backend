package auth

import (
	"log"
	"time"

	"github.com/Project-Fritata/fritata-backend/internal"

	"github.com/gofiber/fiber/v3"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

func Register(c fiber.Ctx) error {
	var data RegisterReq
	if err := c.Bind().JSON(&data); err != nil {
		return internal.InvalidRequest(c)
	}

	// Check if email or password is empty
	if data.Email == "" || data.Password == "" {
		return internal.InvalidRequest(c)
	}

	// Check if the email is already registered
	emailRegistered, err := DbEmailRegistered(data.Email)
	if err != nil {
		return internal.InternalServerError(c)
	}
	if emailRegistered {
		// Email already exists, return a generic error
		return internal.InvalidCredentials(c)
	}

	// Hash the password
	password, _ := bcrypt.GenerateFromPassword([]byte(data.Password), 14)
	auth := internal.Auth{
		Email:    data.Email,
		Password: password,
	}

	// Create new auth
	if err := DbCreateAuthUser(auth); err != nil {
		log.Println(err.Error())
		return internal.InternalServerError(c)
	}

	return c.JSON(fiber.Map{
		"message": "success",
	})
}

func Login(c fiber.Ctx) error {
	var data LoginReq
	if err := c.Bind().JSON(&data); err != nil {
		return internal.InvalidRequest(c)
	}

	// Check if email or passowrd is empty
	if data.Email == "" || data.Password == "" {
		return internal.InvalidCredentials(c)
	}

	// Check if the email is already registered
	emailRegistered, err := DbEmailRegistered(data.Email)
	if err != nil {
		return internal.InternalServerError(c)
	}
	if !emailRegistered {
		// Email isn't registered
		return internal.InvalidCredentials(c)
	}

	// Get auth by email
	auth, err := DbGetAuthByEmail(data.Email)
	if err != nil {
		return internal.InternalServerError(c)
	}

	// Check if the password is correct
	if err := bcrypt.CompareHashAndPassword(auth.Password, []byte(data.Password)); err != nil {
		return internal.InvalidCredentials(c)
	}

	// Create a new JWT token
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer:    auth.Id.String(),
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(), // 1 day
	})
	token, err := claims.SignedString([]byte(internal.GetEnvVar("JWT_SECRET")))
	if err != nil {
		return internal.InternalServerError(c)
	}

	// Set the JWT cookie
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
	// Create empty cookie with expired time
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
