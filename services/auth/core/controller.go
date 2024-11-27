package core

import (
	"fmt"

	"github.com/Project-Fritata/fritata-backend/internal/apierrors"
	"github.com/Project-Fritata/fritata-backend/internal/cookies"
	"github.com/Project-Fritata/fritata-backend/internal/uservalidation"
	"github.com/Project-Fritata/fritata-backend/services/auth/db"
	"github.com/Project-Fritata/fritata-backend/services/auth/models"

	"github.com/gofiber/fiber/v3"
	"golang.org/x/crypto/bcrypt"
)

// Register godoc
//
// @Summary Register
// @Description Register a new account
// @Accept json
// @Produce json
// @Param account body models.RegisterReq true "Account"
// @Success 200 {array} models.RegisterRes
// @Failure 400 {object} apierrors.ErrorResponse "Bad request / Invalid credentials"
// @Failure 500 {object} apierrors.ErrorResponse
// @Router /api/v1/auth/register [post]
func Register(c fiber.Ctx) error {
	var data models.RegisterReq
	if err := c.Bind().JSON(&data); err != nil {
		return apierrors.InvalidRequest(c, fmt.Errorf("cannot parse request body"))
	}

	// Check if email or password is empty
	if data.Email == "" || data.Password == "" {
		return apierrors.InvalidRequest(c, fmt.Errorf("cannot register user with empty email or password"))
	}

	// Check if the username is valid
	if !uservalidation.ValidateInput(data.Username) {
		return apierrors.InvalidRequest(c, fmt.Errorf("invalid email"))
	}

	// Check if the email is already registered
	emailRegistered, err := db.DbEmailRegistered(data.Email)
	if err != nil {
		return apierrors.InternalServerError(c, err)
	}
	if emailRegistered {
		// Email already exists, return a generic error
		return apierrors.InvalidCredentials(c, apierrors.DefaultError())
	}

	// Hash the password
	password, _ := bcrypt.GenerateFromPassword([]byte(data.Password), 14)
	auth := models.Auth{
		Email:    data.Email,
		Password: password,
	}

	// Create new auth
	if err := db.DbCreateAuthUser(auth, data.Username); err != nil {
		return apierrors.InternalServerError(c, err)
	}

	return c.JSON(models.RegisterRes{
		Message: "success",
	})
}

// Login godoc
//
// @Summary Login
// @Description Login into an existing account
// @Accept json
// @Produce json
// @Param account body models.LoginReq true "Account"
// @Success 200 {array} models.LoginRes
// @Failure 400 {object} apierrors.ErrorResponse "Bad request / Invalid credentials"
// @Failure 500 {object} apierrors.ErrorResponse
// @Router /api/v1/auth/login [post]
func Login(c fiber.Ctx) error {
	var data models.LoginReq
	if err := c.Bind().JSON(&data); err != nil {
		return apierrors.InvalidRequest(c, fmt.Errorf("cannot parse request body"))
	}

	// Check if email or passowrd is empty
	if data.Email == "" || data.Password == "" {
		return apierrors.InvalidRequest(c, fmt.Errorf("cannot login with empty email or password"))
	}

	// Check if the email is already registered
	emailRegistered, err := db.DbEmailRegistered(data.Email)
	if err != nil {
		return apierrors.InternalServerError(c, err)
	}
	if !emailRegistered {
		// Email isn't registered
		return apierrors.InvalidCredentials(c, fmt.Errorf("email not registered"))
	}

	// Get auth by email
	auth, err := db.DbGetAuthByEmail(data.Email)
	if err != nil {
		return apierrors.InternalServerError(c, err)
	}

	// Check if the password is correct
	if err := bcrypt.CompareHashAndPassword(auth.Password, []byte(data.Password)); err != nil {
		return apierrors.InvalidCredentials(c, fmt.Errorf("invalid password"))
	}

	// Create and set the cookie
	if err := cookies.CreateSetCookie(c, auth.Id); err != nil {
		return err
	}

	return c.JSON(models.LoginRes{
		Message: "success",
	})
}

// Logout godoc
//
// @Summary Logout
// @Description Logout from logged in account
// @Accept json
// @Produce json
// @Success 200 {array} models.LoginRes
// @Router /api/v1/auth/logout [post]
func Logout(c fiber.Ctx) error {
	cookies.RemoveCookie(c)

	return c.JSON(models.LogoutRes{
		Message: "success",
	})
}
