package data

import (
	"immersiveApp/utils/helpers"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	TeamId      uint
	FullName    string
	Email       string `gorm:"unique"`
	Password    string
	Role        string
	PhoneNumber string
	Address     string
	Status      bool
}

func (user *User) BeforeCreate(tx *gorm.DB) (err error) {
	user.Password, err = helpers.HashPassword(user.Password)
	return
}
