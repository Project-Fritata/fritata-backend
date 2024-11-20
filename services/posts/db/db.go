package db

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/Project-Fritata/fritata-backend/internal/apierrors"
	"github.com/Project-Fritata/fritata-backend/internal/db"
	"github.com/Project-Fritata/fritata-backend/services/posts/models"
	usermodels "github.com/Project-Fritata/fritata-backend/services/users/models"
	"github.com/gofiber/fiber/v3/log"
	"gorm.io/gorm"
)

func ParseQueryParameters(offset int, limit int, sortOrder *models.SortOrder, filters []string) (*gorm.DB, error) {
	query := db.DB.Model(&models.Post{})

	// Sorting
	if err := models.IsValidSortOrder(sortOrder); err != nil {
		return nil, err
	}
	if sortOrder != nil {
		if *sortOrder == models.SortOrderDesc {
			query = query.Order("created_at DESC")
		} else if *sortOrder == models.SortOrderAsc {
			query = query.Order("created_at ASC")
		}
	} else {
		// Default sort: newest first
		query = query.Order("created_at DESC")
	}

	// Filters
	if len(filters) > 0 {
		for _, filter := range filters {

			// Split filter string into components
			filterParts := strings.SplitN(filter, ":", 3)
			if len(filterParts) < 3 {
				return nil, fmt.Errorf("invalid filter format: %s", filter)
			}
			field := filterParts[0]
			operator := filterParts[1]
			value := filterParts[2]

			if err := models.IsValidFilter(field, operator, value); err != nil {
				return nil, err
			}
			switch operator {
			case models.OperatorEquals:
				query = query.Where(field+" = ?", value)
			case models.OperatorNotEquals:
				query = query.Where(field+" != ?", value)
			case models.OperatorGreaterThan:
				query = query.Where(field+" > ?", value)
			case models.OperatorLessThan:
				query = query.Where(field+" < ?", value)
			case models.OperatorContains:
				query = query.Where(field+" ILIKE ?", "%"+value+"%")
			case models.OperatorIn:
				values := strings.Split(value, ",")
				query = query.Where(field+" IN ?", values)
			}
		}
	}

	// Pagination
	return query.Offset(offset).Limit(limit), nil
}
func DbGetPosts(query *gorm.DB) ([]models.GetPostsRes, error) {
	var posts []models.Post
	if err := query.Find(&posts).Error; err != nil {
		log.Errorf("Error getting posts from DB: %w", err)
		return nil, apierrors.DefaultError()
	}

	var res []models.GetPostsRes
	for _, post := range posts {
		// Get user data for post
		client := &http.Client{}
		resp, err := client.Get("http://users:8011/api/v1/users/" + post.Id_User.String())
		if err != nil {
			log.Errorf("Error getting user data from users service for post: %+v\n%w", post, err)
			return nil, apierrors.DefaultError()
		}
		defer resp.Body.Close()
		if resp.StatusCode != http.StatusOK {
			log.Errorf("Error getting user data from users service for post: %+v\n%w", post, err)
			return nil, apierrors.DefaultError()
		}
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Errorf("Error reading body from users service for post: %+v\n%w", post, err)
			return nil, apierrors.DefaultError()
		}

		var userRes usermodels.GetRes
		err = json.Unmarshal(body, &userRes)
		if err != nil {
			log.Errorf("Error unmarshalling user data from users service response body: %s\n%w", string(body), err)
			return nil, apierrors.DefaultError()
		}

		user := usermodels.User{
			Id:          userRes.Id,
			Username:    userRes.Username,
			Pfp:         userRes.Pfp,
			Description: userRes.Description,
		}

		newRes := models.GetPostsRes{
			Post: post,
			User: user,
		}
		res = append(res, newRes)
	}

	return res, nil
}

func DbCreatePost(post models.Post) error {
	if err := db.DB.Create(&post).Error; err != nil {
		log.Errorf("Error creating post in DB: %+v\n%w", post, err)
		return apierrors.DefaultError()
	}
	return nil
}
