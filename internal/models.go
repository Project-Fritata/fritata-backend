package internal

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Auth struct {
	gorm.Model
	Id       uuid.UUID `json:"id" gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	Email    string    `json:"email" gorm:"unique"`
	Password []byte    `json:"-"`
}

func (Auth) TableName() string {
	return "auth"
}

type User struct {
	gorm.Model
	Id          uuid.UUID `json:"id" gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	Username    string    `json:"username"`
	Pfp         string    `json:"pfp"`
	Description string    `json:"description"`
}

type Post struct {
	gorm.Model
	Id      uint      `json:"id"`
	Id_User uuid.UUID `json:"id_user" gorm:"type:uuid"`
	User    User      `gorm:"foreignKey:Id_User;references:Id"`
	Content string    `json:"content"`
	Media   string    `json:"media"`
}
