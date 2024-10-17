package users

import (
	"github.com/Project-Fritata/fritata-backend/internal"
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"

	"github.com/gofiber/fiber/v3"
)

func GetUser(c fiber.Ctx) error {
	var data GetReq
	if err := c.Bind().JSON(&data); err != nil {
		return err
	}

	// Check if Id is empty
	if data.Id == uuid.Nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": "invalid id",
		})
	}

	// Check if the user exists
	var count int64
	if err := internal.DB.Model(&internal.User{}).Where("id = ?", data.Id).Count(&count).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "internal server error",
		})
	}
	if count == 0 {
		// User doesnt exist
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "user not found",
		})
	}

	var user internal.User
	internal.DB.First(&user, "id = ?", data.Id)
	if user.Id == uuid.Nil {
		c.Status(fiber.StatusNotFound)
		return c.JSON(fiber.Map{
			"message": "user not found",
		})
	}

	return c.JSON(GetRes{
		Id:          user.Id,
		Username:    user.Username,
		Pfp:         user.Pfp,
		Description: user.Description,
	})
}

func CreateUser(c fiber.Ctx) error {
	var data CreateReq
	if err := c.Bind().JSON(&data); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": "invalid credentials",
		})
	}

	// Check if Id or username is empty
	if data.Id == uuid.Nil || data.Username == "" {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": "invalid credentials",
		})
	}

	// Create new user
	user := internal.User{
		Id:       data.Id,
		Username: data.Username,
	}
	if err := internal.DB.Create(&user).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "internal server error",
		})
	}

	return c.JSON(fiber.Map{
		"message": "success",
	})
}

func UpdateUser(c fiber.Ctx) error {
	var data UpdateReq
	if err := c.Bind().JSON(&data); err != nil {
		return err
	}

	// Check cookie
	cookie := c.Cookies("jwt")
	token, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(internal.GetEnvVar("JWT_SECRET")), nil
	})
	if err != nil {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message": "unauthenticated",
		})
	}
	claims := token.Claims.(*jwt.StandardClaims)
	id, err := uuid.Parse(claims.Issuer)
	if err != nil {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message": "unauthenticated",
		})
	}

	// Update user info
	var user = internal.User{
		Id:          id,
		Username:    data.Username,
		Pfp:         data.Pfp,
		Description: data.Description,
	}
	if err := internal.DB.Save(&user).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "internal server error",
		})
	}

	return c.JSON(fiber.Map{
		"message": "success",
	})
}
