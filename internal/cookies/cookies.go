package cookies

import (
	"fmt"
	"time"

	"github.com/Project-Fritata/fritata-backend/internal/apierrors"
	"github.com/Project-Fritata/fritata-backend/internal/env"

	"github.com/gofiber/fiber/v3"
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
)

func ValidateCookie(c fiber.Ctx) (uuid.UUID, bool, error) {
	cookie := c.Cookies("jwt")
	token, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(env.GetEnvVar("JWT_SECRET")), nil
	})
	if err != nil || !token.Valid {
		return uuid.Nil, false, apierrors.Unauthenticated(c, fmt.Errorf("invalid token"))
	}

	claims := token.Claims.(*jwt.StandardClaims)
	if claims.Issuer == "" {
		return uuid.Nil, false, apierrors.Unauthenticated(c, fmt.Errorf("empty token issuer"))
	}
	id, err := uuid.Parse(claims.Issuer)
	if err != nil {
		return uuid.Nil, false, apierrors.InternalServerError(c, fmt.Errorf("invalid token id format"))
	}

	if id == uuid.Nil {
		return uuid.Nil, false, apierrors.Unauthenticated(c, fmt.Errorf("invalid token id format"))
	}

	return id, true, nil
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

func CreateSetCookie(c fiber.Ctx, id uuid.UUID) error {
	// Create a new JWT token
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer:    id.String(),
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(), // 1 day
	})
	token, err := claims.SignedString([]byte(env.GetEnvVar("JWT_SECRET")))
	if err != nil {
		return apierrors.InternalServerError(c, fmt.Errorf("error creating JWT token"))
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
	return nil
}
