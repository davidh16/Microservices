package model

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"time"
)

type User struct {
	Id       int64
	Name     string
	Surname  string
	Email    string
	Password string
	Valid    bool
	ValidAt  time.Time
}

func (u *User) BeforeSave(tx *gorm.DB) (err error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(u.Password), 10)
	if err != nil {
		return err
	}
	u.Password = string(hash)
	return
}
