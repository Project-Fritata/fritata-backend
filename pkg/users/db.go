package users

import "github.com/Project-Fritata/fritata-backend/internal"

func DbUserIdCreated(id string) (bool, error) {
	var count int64
	if err := internal.DB.Model(&internal.User{}).Where("id = ?", id).Count(&count).Error; err != nil {
		return false, err
	}
	return count > 0, nil
}
func DbUserUsernameCreated(username string) (bool, error) {
	var count int64
	if err := internal.DB.Model(&internal.User{}).Where("username = ?", username).Count(&count).Error; err != nil {
		return false, err
	}
	return count > 0, nil
}

func DbGetUserById(id string) (internal.User, error) {
	var user internal.User
	if err := internal.DB.Model(&internal.User{}).Where("id = ?", id).First(&user).Error; err != nil {
		return internal.User{}, err
	}
	return user, nil
}
func DbGetUserByUsername(username string) (internal.User, error) {
	var user internal.User
	if err := internal.DB.Model(&internal.User{}).Where("username = ?", username).First(&user).Error; err != nil {
		return internal.User{}, err
	}
	return user, nil
}

func DbUpdateUser(user internal.User) error {
	return internal.DB.Save(&user).Error
}

func DbCreateUser(user internal.User) error {
	return internal.DB.Create(&user).Error
}
