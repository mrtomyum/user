package models

import (
	"github.com/jinzhu/gorm"
	"log"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type DB struct {
	*gorm.DB
}

func NewDB(dataBaseType, conn string) *DB {
	db, err := gorm.Open(dataBaseType, conn)
	if err != nil {
		log.Println(err.Error())
	}

	if err = db.DB().Ping(); err != nil {
		log.Println("DB Connection Error!!!")
	} else {
		log.Println("DB Connected")
	}
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)
	db.SingularTable(true)
	return &DB{db}
}
