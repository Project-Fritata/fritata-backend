package core

import (
	"log"
	"strconv"

	"github.com/Project-Fritata/fritata-backend/internal"
	"github.com/Project-Fritata/fritata-backend/services/posts/db"
	"github.com/Project-Fritata/fritata-backend/services/posts/models"

	"github.com/gofiber/fiber/v3"
)

func GetPosts(c fiber.Ctx) error {
	var data models.GetPostsReq
	if err := c.Bind().Query(&data); err != nil {
		return internal.InvalidRequest(c)
	}

	// Get posts
	posts, err := db.DbGetPosts(data.Offset, data.Limit, data.SortOrder, data.Filters)
	if err != nil {
		log.Println(err.Error())
		return internal.InternalServerError(c)
	}

	// Default returning empty array
	if posts == nil {
		posts = make([]models.GetPostsRes, 0)
	}

	// Add x-total-count header
	c.Response().Header.Set("x-total-count", strconv.Itoa(len(posts)))
	return c.JSON(posts)
}

func CreatePost(c fiber.Ctx) error {
	var data models.CreatePostReq
	if err := c.Bind().JSON(&data); err != nil {
		return internal.InvalidRequest(c)
	}

	// Check cookie
	id, valid := internal.ValidateCookie(c)
	if !valid {
		return internal.Unauthenticated(c)
	}

	// Create new post
	post := models.Post{
		Id_User: id,
		Content: data.Content,
		Media:   data.Media,
	}
	if err := db.DbCreatePost(post); err != nil {
		return internal.InternalServerError(c)
	}

	return c.JSON(models.CreatePostRes{
		Message: "success",
	})
}
