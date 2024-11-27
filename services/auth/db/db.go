package db

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/Project-Fritata/fritata-backend/internal/apierrors"
	"github.com/Project-Fritata/fritata-backend/internal/db"
	"github.com/Project-Fritata/fritata-backend/services/auth/models"
	usermodels "github.com/Project-Fritata/fritata-backend/services/users/models"
	"github.com/gofiber/fiber/v3/log"

	"gorm.io/gorm"
)

func DbEmailRegistered(email string) (bool, error) {
	var count int64
	if err := db.DB.Model(&models.Auth{}).Where("email = ?", email).Count(&count).Error; err != nil {
		return false, apierrors.DefaultError()
	}
	return count > 0, nil
}

func DbCreateAuthUser(auth models.Auth, username string) error {
	return db.DB.Transaction(func(tx *gorm.DB) error {

		// Create new auth
		if err := tx.Model(&models.Auth{}).Create(&auth).Error; err != nil {
			log.Errorf("Error creating auth in DB: %+v\n%w", auth, err)
			return apierrors.DefaultError()
		}

		// Send request to create new user
		createReq := usermodels.CreateReq{
			Id:       auth.Id,
			Username: username,
		}
		reqBody, err := json.Marshal(createReq)
		if err != nil {
			log.Errorf("Error marshalling create user request: %+v\n%w", createReq, err)
			return apierrors.DefaultError()
		}

		client := &http.Client{}
		resp, err := client.Post("http://users:8011/api/v1/users", "application/json", bytes.NewBuffer(reqBody))
		if err != nil {
			log.Errorf("Error creating user in external service: %+v\n%w", createReq, err)
			return apierrors.DefaultError()
		}
		if resp.StatusCode != http.StatusOK {
			log.Errorf("Error creating user in external service: %+v\n%w", createReq, err)
			return apierrors.DefaultError()
		}

		return nil
	})
}

func DbGetAuthByEmail(email string) (models.Auth, error) {
	var auth models.Auth
	if err := db.DB.Model(&models.Auth{}).Where("email = ?", email).First(&auth).Error; err != nil {
		log.Errorf("Error getting auth by email from DB: %w", err)
		return models.Auth{}, apierrors.DefaultError()
	}
	return auth, nil
}
