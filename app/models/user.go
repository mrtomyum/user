package models

import (
	"github.com/jinzhu/gorm"
	"github.com/revel/revel"
	"golang.org/x/crypto/bcrypt"
	"log"
	"regexp"
)

type User struct {
	gorm.Model
	Name           string `gorm:"not null`
	Username       string `gorm:"not null;unique"`
	Password       string `gorm:"-"`
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

var userRegex = regexp.MustCompile("^\\w*$")

func (u *User) Validate(v *revel.Validation) {
	v.Check(u.Username,
		revel.Required{},
		revel.MaxSize{15},
		revel.MinSize{1},
		revel.Match{userRegex},
	)
	v.Check(u.Name,
		revel.Required{},
		revel.MaxSize{100},
	)
//	ValidatePassword(v, u.Password).Key("user.Password")
}

func ValidatePassword(v *revel.Validation, password string) *revel.ValidationResult {
	return v.Check(password,
		revel.Required{},
		revel.MaxSize{15},
		revel.MinSize{5},
	)
}
