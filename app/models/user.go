package models

import (
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
	"log"
)

type User struct {
	gorm.Model
	Name           string
	Username       string
	Password       string	`gorm:"-"`
	HashedPassword []byte
	Role           string
}

func (u *User) SetPass(p string) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(p), bcrypt.DefaultCost)
	if err != nil {
		log.Fatalln(err)
		return err
	}
	u.HashedPassword = hash
	return nil
}

func (u *User) VerifyPass(p string) error {
	err := bcrypt.CompareHashAndPassword(u.HashedPassword, []byte(p))
	if err != nil {
		return err
	}
	return nil
}
