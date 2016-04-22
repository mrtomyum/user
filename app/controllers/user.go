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

func (c User) Save(user m.User, verifyPassword string) revel.Result {
	//c.Params.Bind(&user, "user")
	fmt.Printf("verifyPassword: %v <--> u.Password: %v\n", verifyPassword, user.Password)
	c.Validation.Required(verifyPassword)
	c.Validation.Required(verifyPassword == user.Password).Message("Password does not match")
	user.Validate(c.Validation)

	if c.Validation.HasErrors() {
		c.Validation.Keep()
		c.FlashParams()
		return c.Redirect(User.New)
	}

	user.SetPass(user.Password)
	user.Password = "" // prevent plain text password to be save to database
	db.Create(&user)
	fmt.Printf("User info: %v\n", user)
	c.Flash.Success("User %v saved", user.Name)
	return c.Redirect(User.Index)
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

// TODO: Add User.Edit() and modify User.Index() with link in User list
// Todo: Add html template form with data from selected user for Action Edit()