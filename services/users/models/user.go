package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// User defines a user model for the system
// @Description User represents a user in the system
// @ID User
type User struct {
	gorm.Model  `swaggerignore:"true"`
	Id          uuid.UUID `json:"id" gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	Username    string    `json:"username" gorm:"unique"`
	Pfp         string    `json:"pfp"`
	Description string    `json:"description"`
}
