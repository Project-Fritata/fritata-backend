package posts

import (
	"log"
	"strconv"

	"github.com/Project-Fritata/fritata-backend/internal"
	"github.com/gofiber/fiber/v3"
)

func GetPosts(c fiber.Ctx) error {
	var data GetReq
	if err := c.Bind().Query(&data); err != nil {
		return internal.InvalidRequest(c)
	}

	posts, err := DbGetPosts(data.Offset, data.Limit, data.SortOrder, data.Filters)
	if err != nil {
		log.Println(err.Error())
		return internal.InternalServerError(c)
	}

	if posts == nil {
		posts = make([]GetRes, 0) // Assuming Post is your struct type
	}

	// Add x-total-count header
	c.Response().Header.Set("x-total-count", strconv.Itoa(len(posts)))
	return c.JSON(posts)
}

func CreatePost(c fiber.Ctx) error {
	var data PostReq
	if err := c.Bind().JSON(&data); err != nil {
		return internal.InvalidRequest(c)
	}

	// Check cookie
	id, err := internal.ValidateCookie(c)
	if err != nil {
		return err
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
