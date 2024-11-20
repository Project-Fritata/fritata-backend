package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Post struct {
	gorm.Model `swaggerignore:"true"`
	Id         uint      `json:"id"`
	Id_User    uuid.UUID `json:"id_user" gorm:"type:uuid"`
	Content    string    `json:"content"`
	Media      string    `json:"media"`
}
