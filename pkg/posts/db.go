package posts

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"github.com/Project-Fritata/fritata-backend/internal"
	"github.com/Project-Fritata/fritata-backend/pkg/users"
)

func DbGetPosts(offset int, limit int) ([]GetRes, error) {
	var posts []internal.Post
	if err := internal.DB.Model(&internal.Post{}).Offset(offset).Limit(limit).Find(&posts).Error; err != nil {
		return nil, err
	}

	var res []GetRes

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

		var userRes users.GetRes
		err = json.Unmarshal(body, &userRes)
		if err != nil {
			return nil, errors.New("failed to parse response body")
		}

		user := internal.User{
			Id:          userRes.Id,
			Username:    userRes.Username,
			Pfp:         userRes.Pfp,
			Description: userRes.Description,
		}

		newRes := GetRes{
			Post: post,
			User: user,
		}
		res = append(res, newRes)
	}

	return res, nil
}

func DbCreatePost(post internal.Post) error {
	return internal.DB.Create(&post).Error
}
