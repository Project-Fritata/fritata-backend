package internal

import (
	"time"

	"github.com/gofiber/fiber/v3"
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
)

func ValidateCookie(c fiber.Ctx) (uuid.UUID, error) {
	cookie := c.Cookies("jwt")
	token, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(GetEnvVar("JWT_SECRET")), nil
	})
	if err != nil || !token.Valid {
		return uuid.Nil, Unauthenticated(c)
	}

	claims := token.Claims.(*jwt.StandardClaims)
	if claims.Issuer == "" {
		return uuid.Nil, Unauthenticated(c)
	}
	id, err := uuid.Parse(claims.Issuer)
	if err != nil {
		return uuid.Nil, InternalServerError(c)
	}

	return id, nil
}

func RemoveCookie(c fiber.Ctx) {
	// Create empty cookie with expired time
	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    "",
		Secure:   true,
		SameSite: fiber.CookieSameSiteNoneMode,
		Expires:  time.Now().Add(-24 * time.Hour), // Set to past time
		HTTPOnly: true,
	}

	c.Cookie(&cookie)
}

func CreateSetCookie(c fiber.Ctx, id uuid.UUID) {
	// Create a new JWT token
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer:    id.String(),
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(), // 1 day
	})
	token, err := claims.SignedString([]byte(GetEnvVar("JWT_SECRET")))
	if err != nil {
		InternalServerError(c)
	}

	// Set the JWT cookie
	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    token,
		Secure:   true,
		SameSite: fiber.CookieSameSiteNoneMode,
		Expires:  time.Now().Add(time.Hour * 24),
		HTTPOnly: true,
	}
	c.Cookie(&cookie)
}
