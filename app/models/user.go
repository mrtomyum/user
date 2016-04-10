package models

import (
	"github.com/elithrar/simple-scrypt"
	"github.com/jinzhu/gorm"
	"log"
)

type User struct {
	gorm.Model
	Name     string
	Password []byte
}

func (u *User) SetPass(p string) error {
	hash, err := scrypt.GenerateFromPassword([]byte(p), scrypt.DefaultParams)
	if err != nil {
		log.Fatalln(err)
		return err
	}
	u.Password = hash
	return nil
}

func (u *User) VerifyPass(p string) error {
	err := scrypt.CompareHashAndPassword(u.Password, []byte(p))
	if err != nil {
		return err
	}
	return nil
}
