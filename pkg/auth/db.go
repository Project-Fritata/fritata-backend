package auth

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/Project-Fritata/fritata-backend/internal"
	"github.com/Project-Fritata/fritata-backend/pkg/users"
	"gorm.io/gorm"
)

func DbEmailRegistered(email string) (bool, error) {
	var count int64
	if err := internal.DB.Model(&internal.Auth{}).Where("email = ?", email).Count(&count).Error; err != nil {
		return false, err
	}
	return count > 0, nil
}

func DbCreateAuthUser(auth internal.Auth) error {
	return internal.DB.Transaction(func(tx *gorm.DB) error {

		// Create new auth
		if err := tx.Model(&internal.Auth{}).Create(&auth).Error; err != nil {
			return err
		}

		// Send request to create new user
		reqBody, err := json.Marshal(
			users.CreateReq{
				Id:       auth.Id,
				Username: auth.Email,
			},
		)
		if err != nil {
			return err
		}

		client := &http.Client{}
		resp, err := client.Post("http://localhost:8011/api/v1/users", "application/json", bytes.NewBuffer(reqBody))
		if err != nil {
			return err
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			return errors.New("failed to create user in external service")
		}

		return nil
	})
}

func DbGetAuthByEmail(email string) (internal.Auth, error) {
	var auth internal.Auth
	if err := internal.DB.Model(&internal.Auth{}).Where("email = ?", email).First(&auth).Error; err != nil {
		return internal.Auth{}, err
	}
	return auth, nil
}
