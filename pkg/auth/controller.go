package auth

import (
	"time"

	"github.com/Project-Fritata/fritata-backend/internal"
	"gorm.io/gorm"

	"github.com/gofiber/fiber/v3"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

func Register(c fiber.Ctx) error {
	var data RegisterReq
	if err := c.Bind().JSON(&data); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": "invalid credentials",
		})
	}

	// Check if email or passowrd is empty
	if data.Email == "" || data.Password == "" {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": "invalid credentials",
		})
	}

	// Check if the email is already registered
	var count int64
	if err := internal.DB.Model(&internal.Auth{}).Where("email = ?", data.Email).Count(&count).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "internal server error",
		})
	}
	if count > 0 {
		// Email already exists, return a generic error
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "invalid credentials",
		})
	}

	// Hash the password
	password, _ := bcrypt.GenerateFromPassword([]byte(data.Password), 14)
	auth := internal.Auth{
		Email:    data.Email,
		Password: password,
	}

	// Create new auth and user
	if err := internal.DB.Transaction(func(tx *gorm.DB) error {

		if err := tx.Create(&auth).Error; err != nil {
			return err
		}
		user := internal.User{
			Id:       auth.Id,
			Username: data.Email,
		}
		if err := tx.Create(&user).Error; err != nil {
			return err
		}

		return nil
	}); err != nil {
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": "failed to create user",
		})
	}

	return c.JSON(fiber.Map{
		"message": "success",
	})
}

func Login(c fiber.Ctx) error {
	var data LoginReq
	if err := c.Bind().JSON(&data); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": "invalid credentials",
		})
	}

	// Check if email or passowrd is empty
	if data.Email == "" || data.Password == "" {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": "invalid credentials",
		})
	}

	// Check if the email is already registered
	var count int64
	if err := internal.DB.Model(&internal.Auth{}).Where("email = ?", data.Email).Count(&count).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "internal server error",
		})
	}
	if count == 0 {
		// Email doesnt exist
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "invalid credentials",
		})
	}

	var user internal.Auth
	internal.DB.First(&user, "email = ?", data.Email)

	// Check if the password is correct
	if err := bcrypt.CompareHashAndPassword(user.Password, []byte(data.Password)); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": "invalid credentials",
		})
	}

	// Create a new JWT token
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer:    user.Id.String(),
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(), // 1 day
	})
	token, err := claims.SignedString([]byte(internal.GetEnvVar("JWT_SECRET")))
	if err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": "could not login",
		})
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
