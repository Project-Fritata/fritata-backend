package core

import (
	"fmt"
	"strconv"

	"github.com/Project-Fritata/fritata-backend/internal/apierrors"
	"github.com/Project-Fritata/fritata-backend/internal/apihealth"
	"github.com/Project-Fritata/fritata-backend/internal/cookies"
	"github.com/Project-Fritata/fritata-backend/services/posts/db"
	"github.com/Project-Fritata/fritata-backend/services/posts/models"

	"github.com/gofiber/fiber/v3"
)

// GetPosts godoc
//
// @Summary Get posts
// @Description Get posts, supports pagination, sorting and filtering
// @Accept json
// @Produce json
// @Param offset query int false "Offset"
// @Param limit query int false "Limit"
// @Param sort query models.SortOrder false "Sort order"
// @Param filters query []string false "Filters" collectionFormat(multi)
// @Success 200 {array} models.GetPostsRes
// @Failure 400 {object} apierrors.ErrorResponse
// @Failure 500 {object} apierrors.ErrorResponse
// @Router /api/v1/posts [get]
func GetPosts(c fiber.Ctx) error {
	var data models.GetPostsReq
	if err := c.Bind().Query(&data); err != nil {
		return apierrors.InvalidRequest(c, fmt.Errorf("cannot parse request parameters"))
	}

	// Get parameters
	query, err := db.ParseQueryParameters(data.Offset, data.Limit, data.SortOrder, data.Filters)
	if err != nil {
		return apierrors.InvalidRequest(c, err)
	}

	// Get posts
	posts, err := db.DbGetPosts(query)
	if err != nil {
		return apierrors.InternalServerError(c, err)
	}

	// Default returning empty array
	if posts == nil {
		posts = make([]models.GetPostsRes, 0)
	}

	// Add x-total-count header
	c.Response().Header.Set("x-total-count", strconv.Itoa(len(posts)))
	return c.JSON(posts)
}

// CreatePost godoc
//
// @Summary Create post
// @Description Create post
// @Accept json
// @Produce json
// @Param post body models.CreatePostReq true "Post"
// @Success 200 {object} models.CreatePostRes
// @Failure 400 {object} apierrors.ErrorResponse
// @Failure 401 {object} apierrors.ErrorResponse
// @Failure 422 {object} apierrors.ErrorResponse
// @Failure 500 {object} apierrors.ErrorResponse
// @Router /api/v1/posts [post]
func CreatePost(c fiber.Ctx) error {
	var data models.CreatePostReq
	if err := c.Bind().JSON(&data); err != nil {
		return apierrors.InvalidRequest(c, apierrors.DefaultError())
	}

	// Check cookie
	id, valid, err := cookies.ValidateCookie(c)
	if !valid {
		return err
	}

	// Create new post
	post := models.Post{
		Id_User: id,
		Content: data.Content,
		Media:   data.Media,
	}

	// Check moderation status
	moderationStatus, err := CheckModerationStatus(post)
	if err != nil {
		if !moderationStatus {
			return apierrors.UnprocessableEntity(c, err)
		}
		return apierrors.InternalServerError(c, err)
	}

	// Create post
	if err := db.DbCreatePost(post); err != nil {
		return apierrors.InternalServerError(c, err)
	}

	return c.JSON(models.CreatePostRes{
		Message: "success",
	})
}

// Health godoc
//
// @Summary Health
// @Description Health check
// @Accept json
// @Produce json
// @Success 200 {object} apihealth.HealthRes
// @Router /api/v1/health/posts [get]
func Health(c fiber.Ctx) error {
	return apihealth.Health(c, apihealth.Posts)
}
