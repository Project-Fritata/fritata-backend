package db

import (
	"github.com/Project-Fritata/fritata-backend/internal"
	"github.com/Project-Fritata/fritata-backend/services/users/models"
)

func DbUserIdCreated(id string) (bool, error) {
	var count int64
	if err := internal.DB.Model(&models.User{}).Where("id = ?", id).Count(&count).Error; err != nil {
		return false, err
	}
	return count > 0, nil
}
func DbUserUsernameCreated(username string) (bool, error) {
	var count int64
	if err := internal.DB.Model(&models.User{}).Where("username = ?", username).Count(&count).Error; err != nil {
		return false, err
	}
	return count > 0, nil
}

func DbGetUserById(id string) (models.User, error) {
	var user models.User
	if err := internal.DB.Model(&models.User{}).Where("id = ?", id).First(&user).Error; err != nil {
		return models.User{}, err
	}
	return user, nil
}
func DbGetUserByUsername(username string) (models.User, error) {
	var user models.User
	if err := internal.DB.Model(&models.User{}).Where("username = ?", username).First(&user).Error; err != nil {
		return models.User{}, err
	}
	return user, nil
}

func DbUpdateUser(user models.User) error {
	return internal.DB.Save(&user).Error
}

func DbCreateUser(user models.User) error {
	return internal.DB.Create(&user).Error
}
