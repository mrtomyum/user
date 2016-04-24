package controllers

import (
	"encoding/json"
	"fmt"
	m "github.com/mrtomyum/user/app/models"
	"github.com/revel/revel"
)

type Users struct {
	App
}

func (c Users) getUserByName(username string) *m.User {
	user := new(m.User)
	db.First(&user, "username = ?", username)
	if user == nil {
		return nil
	}
	return user
}

func (c Users) getUserByID(id string) *m.User {
	user := new(m.User)
	db.First(&user, "id = ?", id)
	if user == nil {
		return nil
	}
	return user
}

func (c Users) checkUser() revel.Result {
	if user := c.connected(); user == nil {
		c.Flash.Error("Please log in first")
		return c.Redirect(App.Index)
	}
	return nil
}

func (c Users) Index() revel.Result {
	c.checkUser()
	fmt.Printf("c.connetced() = %v", c.connected().Username)
	users := []m.User{}
	db.Debug().Find(&users)
	//fmt.Println(users)
	return c.Render(users)
}

func (c Users) New() revel.Result {
	return c.Render()
}

func (c Users) Add(user m.User, verifyPassword string) revel.Result {
	fmt.Printf("verifyPassword: %v <--> u.Password: %v\n", verifyPassword, user.Password)
	c.Validation.Required(verifyPassword)
	c.Validation.Required(verifyPassword == user.Password).Message("Password does not match")
	user.Validate(c.Validation)

	if c.Validation.HasErrors() {
		c.Validation.Keep()
		c.FlashParams()
		return c.Redirect(Users.New)
	}

	user.SetPass(user.Password)
	user.Password = "" // prevent plain text password to be save to database

	rows := db.Debug().Create(&user).RowsAffected
	if rows == 0 {
		c.Flash.Error("Error!!, may be Duplicate Username.")
		return c.Redirect(Users.New)
	}
	c.Flash.Success("User %v added", user.Name)
	return c.Redirect(Users.Index)
}
func (c Users) Show(id string) revel.Result {
	//user := new(m.User)
	//db.Debug().First(&user, "id = ?", id)
	user := c.getUserByID(id)
	return c.Render(user)
}

func (c Users) Edit(id string) revel.Result {
	user := c.getUserByID(id)
	fmt.Println(user)
	return c.Redirect(Users.Edit, user, id)
}

func (c Users) Save(user m.User, verifyPassword string) revel.Result {
	fmt.Printf("verifyPassword: %v <--> u.Password: %v\n", verifyPassword, user.Password)
	c.Validation.Required(verifyPassword)
	c.Validation.Required(verifyPassword == user.Password).Message("Password does not match")
	user.Validate(c.Validation)

	if c.Validation.HasErrors() {
		c.Validation.Keep()
		c.FlashParams()
		return c.Redirect(Users.New)
	}

	user.SetPass(user.Password)
	user.Password = "" // prevent plain text password to be save to database

	rows := db.Debug().Update(&user).RowsAffected
	if rows == 0 {
		c.Flash.Error("Error!!, may be Duplicate Username.")
		return c.Redirect(Users.Edit, user)
	}
	//fmt.Printf("User info: %v\n", user)
	c.Flash.Success("User %v saved", user.Name)
	return c.Redirect(Users.Show, user)
}

func (c Users) ApiPost() revel.Result {
	var user m.User
	dec := json.NewDecoder(c.Request.Body)
	dec.Decode(&user)
	fmt.Printf("The Order data: %v\n", user)
	return c.Render(user)
}

// TODO: Add User.Edit() and modify User.Index() with link in User list
// Todo: Add html template form with data from selected user for Action Edit()
