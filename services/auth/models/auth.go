package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Auth struct {
	gorm.Model `swaggerignore:"true"`
	Id         uuid.UUID `json:"id" gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	Email      string    `json:"email" gorm:"unique"`
	Password   []byte    `json:"-"`
}

func (Auth) TableName() string {
	return "auth"
}
