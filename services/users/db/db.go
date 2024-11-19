package db

import (
	"github.com/Project-Fritata/fritata-backend/internal/apierrors"
	"github.com/Project-Fritata/fritata-backend/internal/db"
	"github.com/Project-Fritata/fritata-backend/services/users/models"
	"github.com/gofiber/fiber/v3/log"
)

func DbUserIdExists(id string) (bool, error) {
	var count int64
	if err := db.DB.Model(&models.User{}).Where("id = ?", id).Count(&count).Error; err != nil {
		log.Errorf("Error checking if user with id %s exists in DB: %w", id, err)
		return false, apierrors.DefaultError()
	}
	return count > 0, nil
}
func DbUserUsernameExists(username string) (bool, error) {
	var count int64
	if err := db.DB.Model(&models.User{}).Where("username = ?", username).Count(&count).Error; err != nil {
		log.Errorf("Error checking if user with username %s exists in DB: %w", username, err)
		return false, apierrors.DefaultError()
	}
	return count > 0, nil
}

func DbGetUserById(id string) (models.User, error) {
	var user models.User
	if err := db.DB.Model(&models.User{}).Where("id = ?", id).First(&user).Error; err != nil {
		log.Errorf("Error getting user with id %s from DB: %w", id, err)
		return models.User{}, apierrors.DefaultError()
	}
	return user, nil
}
func DbGetUserByUsername(username string) (models.User, error) {
	var user models.User
	if err := db.DB.Model(&models.User{}).Where("username = ?", username).First(&user).Error; err != nil {
		log.Errorf("Error getting user with username %s from DB: %w", username, err)
		return models.User{}, apierrors.DefaultError()
	}
	return user, nil
}

func DbUpdateUser(user models.User) error {
	if err := db.DB.Model(&user).Updates(&user).Error; err != nil {
		log.Errorf("Error updating user in DB: %+v\n%w", user, err)
		return apierrors.DefaultError()
	}
	return nil
}

func DbCreateUser(user models.User) error {
	if err := db.DB.Create(&user).Error; err != nil {
		log.Errorf("Error creating user in DB: %+v\n%w", user, err)
		return apierrors.DefaultError()
	}
	return nil
}
