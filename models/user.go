package models

import (
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID       int     `json:"id"`
	Username *string `json:"username,omitempty"`
	Password *string `json:"password,omitempty"`
	Email    *string `json:"email,omitempty"`
}

func (user *User) EncryptPassword() error {
	hash, err := bcrypt.GenerateFromPassword([]byte(*user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	stringHash := string(hash)
	user.Password = &stringHash
	return nil
}

func (user *User) ComparePassword(password string) bool {
	return bcrypt.CompareHashAndPassword([]byte(*user.Password), []byte(password)) == nil
}
