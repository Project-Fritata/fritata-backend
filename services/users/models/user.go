package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Id          uuid.UUID `json:"id" gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	Username    string    `json:"username" gorm:"unique"`
	Pfp         string    `json:"pfp"`
	Description string    `json:"description"`
}
