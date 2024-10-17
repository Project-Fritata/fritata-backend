package internal

type User struct {
	ID       uint
	Name     string
	Email    string `gorm:"unique"`
	Password []byte
}
