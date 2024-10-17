package users

import (
	"github.com/Project-Fritata/fritata-backend/internal"
	"github.com/golang-jwt/jwt"

	"github.com/gofiber/fiber/v3"
)

func GetUser(c fiber.Ctx) error {
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

	var user internal.User
	internal.DB.First(&user, "id = ?", claims.Issuer)

	return c.JSON(user)
}
