package core

import (
	"fmt"

	"github.com/Project-Fritata/fritata-backend/internal/apierrors"
	"github.com/Project-Fritata/fritata-backend/internal/cookies"
	"github.com/Project-Fritata/fritata-backend/services/users/db"
	"github.com/Project-Fritata/fritata-backend/services/users/models"
	"github.com/google/uuid"

	"github.com/gofiber/fiber/v3"
)

func GetUserById(c fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return apierrors.InvalidRequest(c, fmt.Errorf("cannot parse request parameters"))
	}

	// Check if the user exists
	userExists, err := db.DbUserIdExists(id.String())
	if err != nil {
		return apierrors.InternalServerError(c, apierrors.DefaultError())
	}
	if !userExists {
		// User doesnt exist
		return apierrors.UserNotFound(c, fmt.Errorf("user with id %s not found", id.String()))
	}

	// Get user
	user, err := db.DbGetUserById(id.String())
	if err != nil {
		return apierrors.InternalServerError(c, apierrors.DefaultError())
	}

	return c.JSON(models.GetRes{
		Id:          user.Id,
		Username:    user.Username,
		Pfp:         user.Pfp,
		Description: user.Description,
	})
}

// GetUserByUsername godoc
//
// @Summary GetUserByUsername
// @Description Get data for user with provided username
// @Accept json
// @Produce json
// @Param username query string true "Username"
// @Success 200 {array} models.GetRes
// @Failure 400 {object} apierrors.ErrorResponse
// @Failure 404 {object} apierrors.ErrorResponse
// @Failure 500 {object} apierrors.ErrorResponse
// @Router /api/v1/users [get]
func GetUserByUsername(c fiber.Ctx) error {
	username := c.Params("username")
	if username == "" {
		return apierrors.InvalidRequest(c, fmt.Errorf("cannot parse request parameters"))
	}

	// Check if the user exists
	userCreated, err := db.DbUserUsernameExists(username)
	if err != nil {
		return apierrors.InternalServerError(c, err)
	}
	if !userCreated {
		// User doesnt exist
		return apierrors.UserNotFound(c, fmt.Errorf("user with username %s not found", username))
	}

	// Get user
	user, err := db.DbGetUserByUsername(username)
	if err != nil {
		return apierrors.InternalServerError(c, err)
	}

	return c.JSON(models.GetRes{
		Id:          user.Id,
		Username:    user.Username,
		Pfp:         user.Pfp,
		Description: user.Description,
	})
}

// GetUserByAuth godoc
//
// @Summary GetUserByAuth
// @Description Get data for user that is logged in (based on provided JWT token)
// @Accept json
// @Produce json
// @Success 200 {array} models.GetRes
// @Failure 401 {object} apierrors.ErrorResponse
// @Failure 500 {object} apierrors.ErrorResponse
// @Router /api/v1/users [get]
func GetUserByAuth(c fiber.Ctx) error {
	// Check cookie
	id, err := cookies.ValidateCookie(c)
	if err != nil {
		return err
	}

	// Get user
	user, err := db.DbGetUserById(id.String())
	if err != nil {
		return apierrors.InternalServerError(c, err)
	}

	return c.JSON(models.GetRes{
		Id:          user.Id,
		Username:    user.Username,
		Pfp:         user.Pfp,
		Description: user.Description,
	})
}

// UpdateUser godoc
//
// @Summary UpdateUser
// @Description Update user data - username, pfp and description
// @Accept json
// @Produce json
// @Param data body models.UpdateReq true "Data"
// @Success 200 {array} models.GetRes
// @Failure 400 {object} apierrors.ErrorResponse
// @Failure 401 {object} apierrors.ErrorResponse
// @Failure 500 {object} apierrors.ErrorResponse
// @Router /api/v1/users [put]
func UpdateUser(c fiber.Ctx) error {
	var data models.UpdateReq
	if err := c.Bind().JSON(&data); err != nil {
		return apierrors.InvalidRequest(c, fmt.Errorf("cannot parse request body"))
	}

	// Check cookie
	id, err := cookies.ValidateCookie(c)
	if err != nil {
		return err
	}

	// Update user info
	var user = models.User{
		Id:          id,
		Username:    data.Username,
		Pfp:         data.Pfp,
		Description: data.Description,
	}
	if err := db.DbUpdateUser(user); err != nil {
		return apierrors.InternalServerError(c, err)
	}

	return c.JSON(fiber.Map{
		"message": "success",
	})
}

func CreateUser(c fiber.Ctx) error {
	var data models.CreateReq
	if err := c.Bind().JSON(&data); err != nil {
		return apierrors.InvalidRequest(c, fmt.Errorf("cannot parse request body"))
	}

	// Check if Id or username is empty
	if data.Id == uuid.Nil || data.Username == "" {
		return apierrors.InvalidCredentials(c, fmt.Errorf("cannot create user with empty id or username"))
	}

	// Create new user
	user := models.User{
		Id:       data.Id,
		Username: data.Username,
	}
	if err := db.DbCreateUser(user); err != nil {
		return apierrors.InternalServerError(c, err)
	}

	return c.JSON(fiber.Map{
		"message": "success",
	})
}
