package controllers

import (
	"fmt"
	m "github.com/mrtomyum/user/app/models"
	"github.com/revel/revel"
	"encoding/json"
	//"github.com/mrtomyum/user/app/routes"
)

var dbType string = "sqlite3"
var dbFile string = "./app/models/user.db"
var db *m.DB = m.NewDB(dbType, dbFile)

type User struct {
	*revel.Controller
}

func (c User) Index() revel.Result {
	users := []m.User{}
	db.Debug().Find(&users)
	fmt.Println(users)
	return c.Render(users)
}

func (c User) New() revel.Result {
	return c.Render()
}

func (c User) NewPost() revel.Result {
	var user m.User
	c.Params.Bind(&user, "user")
	fmt.Printf("User info: %v\n", user)
	return c.RenderTemplate("User/New.html")
}

func (c User) ApiPost() revel.Result {
	var user m.User
	dec := json.NewDecoder(c.Request.Body)
	dec.Decode(&user)
	fmt.Printf("The Order data: %v\n", user)
	return c.Render(user)
}

func (c User) Show(id uint) revel.Result {
	return c.Render()
}
