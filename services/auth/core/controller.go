package core

import (
	"log"

	"github.com/Project-Fritata/fritata-backend/internal"
	"github.com/Project-Fritata/fritata-backend/services/auth/db"
	"github.com/Project-Fritata/fritata-backend/services/auth/models"

	"github.com/gofiber/fiber/v3"
	"golang.org/x/crypto/bcrypt"
)

func Register(c fiber.Ctx) error {
	var data models.RegisterReq
	if err := c.Bind().JSON(&data); err != nil {
		return internal.InvalidRequest(c)
	}

	// Check if email or password is empty
	if data.Email == "" || data.Password == "" {
		return internal.InvalidRequest(c)
	}

	// Check if the email is already registered
	emailRegistered, err := db.DbEmailRegistered(data.Email)
	if err != nil {
		return internal.InternalServerError(c)
	}
	if emailRegistered {
		// Email already exists, return a generic error
		return internal.InvalidCredentials(c)
	}

	// Hash the password
	password, _ := bcrypt.GenerateFromPassword([]byte(data.Password), 14)
	auth := models.Auth{
		Email:    data.Email,
		Password: password,
	}

	// Create new auth
	if err := db.DbCreateAuthUser(auth); err != nil {
		log.Println(err.Error())
		return internal.InternalServerError(c)
	}

	return c.JSON(models.RegisterRes{
		Message: "success",
	})
}

func Login(c fiber.Ctx) error {
	var data models.LoginReq
	if err := c.Bind().JSON(&data); err != nil {
		return internal.InvalidRequest(c)
	}

	// Check if email or passowrd is empty
	if data.Email == "" || data.Password == "" {
		return internal.InvalidRequest(c)
	}

	// Check if the email is already registered
	emailRegistered, err := db.DbEmailRegistered(data.Email)
	if err != nil {
		return internal.InternalServerError(c)
	}
	if !emailRegistered {
		// Email isn't registered
		return internal.InvalidCredentials(c)
	}

	// Get auth by email
	auth, err := db.DbGetAuthByEmail(data.Email)
	if err != nil {
		return internal.InternalServerError(c)
	}

	// Check if the password is correct
	if err := bcrypt.CompareHashAndPassword(auth.Password, []byte(data.Password)); err != nil {
		return internal.InvalidCredentials(c)
	}

	// Create and set the cookie
	internal.CreateSetCookie(c, auth.Id)

	return c.JSON(models.LoginRes{
		Message: "success",
	})
}

func Logout(c fiber.Ctx) error {
	internal.RemoveCookie(c)

	return c.JSON(models.LogoutRes{
		Message: "success",
	})
}
