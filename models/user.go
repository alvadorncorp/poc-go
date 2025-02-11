package models

import (
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Username string `form:"username"`
	Password string `form:"password"`
	Email    string `form:"email"`
}

func (user *User) EncryptPassword() error {
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hash)
	return nil
}

func (user *User) ComparePassword(password string) bool {
	return bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)) == nil
}
