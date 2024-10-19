package posts

import (
	"log"

	"github.com/Project-Fritata/fritata-backend/internal"
	"github.com/gofiber/fiber/v3"
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
)

func GetPosts(c fiber.Ctx) error {
	var data GetReq
	if err := c.Bind().Query(&data); err != nil {
		return internal.InvalidRequest(c)
	}

	posts, err := DbGetPosts(data.Offset, data.Limit)
	if err != nil {
		log.Println(err.Error())
		return internal.InternalServerError(c)
	}

	return c.JSON(posts)
}

func CreatePost(c fiber.Ctx) error {
	var data PostReq
	if err := c.Bind().JSON(&data); err != nil {
		return internal.InvalidRequest(c)
	}

	// Check cookie
	cookie := c.Cookies("jwt")
	token, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(internal.GetEnvVar("JWT_SECRET")), nil
	})
	if err != nil {
		return internal.Unauthenticated(c)
	}
	claims := token.Claims.(*jwt.StandardClaims)
	id, err := uuid.Parse(claims.Issuer)
	if err != nil {
		return internal.InternalServerError(c)
	}

	// Create new post
	post := internal.Post{
		Id_User: id,
		Content: data.Content,
		Media:   data.Media,
	}
	if err := DbCreatePost(post); err != nil {
		return internal.InternalServerError(c)
	}

	return c.JSON(fiber.Map{
		"message": "success",
	})
}
