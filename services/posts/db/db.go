package db

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strings"

	"github.com/Project-Fritata/fritata-backend/internal"
	"github.com/Project-Fritata/fritata-backend/services/posts/models"
	usermodels "github.com/Project-Fritata/fritata-backend/services/users/models"
)

func DbGetPosts(offset int, limit int, sortOrder *models.SortOrder, filters []models.Filter) ([]models.GetPostsRes, error) {
	query := internal.DB.Model(&models.Post{})

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
			if err := models.IsValidFilter(filter); err != nil {
				return nil, err
			}
			switch filter.Operator {
			case models.OperatorEquals:
				query = query.Where(filter.Field+" = ?", filter.Value)
			case models.OperatorNotEquals:
				query = query.Where(filter.Field+" != ?", filter.Value)
			case models.OperatorGreaterThan:
				query = query.Where(filter.Field+" > ?", filter.Value)
			case models.OperatorLessThan:
				query = query.Where(filter.Field+" < ?", filter.Value)
			case models.OperatorContains:
				query = query.Where(filter.Field+" ILIKE ?", "%"+filter.Value+"%")
			case models.OperatorIn:
				values := strings.Split(filter.Value, ",")
				query = query.Where(filter.Field+" IN ?", values)
			}
		}
	}

	// Pagination
	query = query.Offset(offset).Limit(limit)

	var posts []models.Post
	if err := query.Find(&posts).Error; err != nil {
		return nil, err
	}

	var res []models.GetPostsRes
	for _, post := range posts {
		// Get user data for post
		client := &http.Client{}
		resp, err := client.Get("http://users:8011/api/v1/users/" + post.Id_User.String())
		if err != nil {
			return nil, err
		}
		defer resp.Body.Close()
		if resp.StatusCode != http.StatusOK {
			return nil, errors.New("failed to get user in external service")
		}
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, errors.New("failed to read response body")
		}

		var userRes usermodels.GetRes
		err = json.Unmarshal(body, &userRes)
		if err != nil {
			return nil, errors.New("failed to parse response body")
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
	return internal.DB.Create(&post).Error
}
