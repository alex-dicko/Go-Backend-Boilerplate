package models

import (
	"boilerplate/utils"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model

	Email    string
	Username string
	Password string
}

// Will set an encrypted password for a user
// takes the raw password as the only parameter
func (u *User) SetPassword(password string) {
	hased_password, _ := utils.HashPassword(password)

	u.Password = hased_password
}
