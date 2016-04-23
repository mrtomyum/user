package controllers

import (
	"github.com/revel/revel"
	"github.com/mrtomyum/user/app/models"
	"golang.org/x/crypto/bcrypt"
	"fmt"
)

var dbType string = "sqlite3"
var dbFile string = "./app/models/user.db"
var db = InitDB(dbType, dbFile)

type App struct {
	*revel.Controller
}

func (c App) AddUser() revel.Result {
	if user := c.connected(); user != nil {
		c.RenderArgs["user"] = user
	}
	return nil
}

func (c App) connected() *models.User {
	if c.RenderArgs["user"] != nil {
		return c.RenderArgs["user"].(*models.User)
	}
	if username, ok := c.Session["user"]; ok {
		return c.getUser(username)
	}
	return nil
}

func (c App) getUser(username string) *models.User {
	user := new(models.User)
	db.First(&user, "username = ?", username)
	if user == nil {
		return nil
	}
	return user
}

func (c App) Index() revel.Result {
	if c.connected() != nil {
		fmt.Printf("c.connetced() = %v", c.connected())
		return c.Redirect(Users.Index)
	}
	currentLocale := c.Request.Locale // ตรวจสอบ Locale ของผู้ใช้
	c.RenderArgs["controllerGreeting"] = c.Message("greeting")
	c.Flash.Error("กรุณา login ก่อน")
	return c.Render(currentLocale)
}

func (c App) Login(username, password string, remember bool) revel.Result {

	user := c.getUser(username)
	if user != nil {
		err := bcrypt.CompareHashAndPassword(user.HashedPassword, []byte(password))
		if err == nil {
			c.Session["user"] = username
			if remember {
				c.Session.SetNoExpiration()
			} else {
				c.Session.SetDefaultExpiration()
			}
			c.Flash.Success("Welcome, " + username)
			return c.Redirect(Users.Index)
		}
	}

	c.Flash.Out["username"] = username
	c.Flash.Error("Login failed")
	return c.Redirect(App.Index)
}

func (c App) Logout() revel.Result {
	for k := range c.Session {
		delete(c.Session, k)
	}
	return c.Redirect(App.Index)
}
