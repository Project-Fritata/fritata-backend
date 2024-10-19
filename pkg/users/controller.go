package users

import (
	"github.com/Project-Fritata/fritata-backend/internal"
	"github.com/google/uuid"

	"github.com/gofiber/fiber/v3"
)

func GetUserById(c fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return internal.InvalidRequest(c)
	}

	// Check if Id is empty
	if id == uuid.Nil {
		return internal.InvalidRequest(c)
	}

	// Check if the user exists
	userCreated, err := DbUserIdCreated(id.String())
	if err != nil {
		return internal.InternalServerError(c)
	}
	if !userCreated {
		// User doesnt exist
		return internal.UserNotFound(c)
	}

	// Get user
	user, err := DbGetUserById(id.String())
	if err != nil {
		return internal.InternalServerError(c)
	}
	if user.Id == uuid.Nil {
		return internal.InvalidRequest(c)
	}

	return c.JSON(GetRes{
		Id:          user.Id,
		Username:    user.Username,
		Pfp:         user.Pfp,
		Description: user.Description,
	})
}

func GetUserByUsername(c fiber.Ctx) error {
	username := c.Params("username")
	// Check if Username is empty
	if username == "" {
		return internal.InvalidRequest(c)
	}

	// Check if the user exists
	userCreated, err := DbUserUsernameCreated(username)
	if err != nil {
		return internal.InternalServerError(c)
	}
	if !userCreated {
		// User doesnt exist
		return internal.UserNotFound(c)
	}

	// Get user
	user, err := DbGetUserByUsername(username)
	if err != nil {
		return internal.InternalServerError(c)
	}
	if user.Id == uuid.Nil {
		return internal.InvalidRequest(c)
	}

	return c.JSON(GetRes{
		Id:          user.Id,
		Username:    user.Username,
		Pfp:         user.Pfp,
		Description: user.Description,
	})
}

func UpdateUser(c fiber.Ctx) error {
	var data UpdateReq
	if err := c.Bind().JSON(&data); err != nil {
		return internal.InvalidRequest(c)
	}

	// Check cookie
	id, err := internal.ValidateCookie(c)
	if err != nil {
		return err
	}

	// Update user info
	var user = internal.User{
		Id:          id,
		Username:    data.Username,
		Pfp:         data.Pfp,
		Description: data.Description,
	}
	if err := DbUpdateUser(user); err != nil {
		return internal.InternalServerError(c)
	}

	return c.JSON(fiber.Map{
		"message": "success",
	})
}

func CreateUser(c fiber.Ctx) error {
	var data CreateReq
	if err := c.Bind().JSON(&data); err != nil {
		return internal.InvalidRequest(c)
	}

	// Check if Id or username is empty
	if data.Id == uuid.Nil || data.Username == "" {
		return internal.InvalidCredentials(c)
	}

	// Create new user
	user := internal.User{
		Id:       data.Id,
		Username: data.Username,
	}
	if err := DbCreateUser(user); err != nil {
		return internal.InternalServerError(c)
	}

	return c.JSON(fiber.Map{
		"message": "success",
	})
}
