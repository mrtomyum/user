package controllers

import (
	"fmt"
	m "github.com/mrtomyum/user/app/models"
	"github.com/revel/revel"
)

var dbType string = "sqlite3"
var dbFile string = "./app/models/user.db"
var db *m.DB = m.NewDB(dbType, dbFile)

type User struct {
	*revel.Controller
}

func (c User) All() revel.Result {
	users := []m.User{}
	db.Debug().Find(&users)
	fmt.Println(users)
	return c.Render(users)
}

func (c User) New() revel.Result {
	return c.Render()
}

func (c User) Show(id uint) revel.Result {
	return c.Render()
}
