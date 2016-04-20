package controllers

import (
	"github.com/revel/revel"
	"github.com/mrtomyum/user/app/models"
	//"github.com/mrtomyum/user/app/routes"
	//"golang.org/x/crypto/bcrypt"
	"golang.org/x/crypto/bcrypt"
)

type App struct {
	*revel.Controller
}

//func (c App) connected() *models.User {
//	if c.RenderArgs["user"] != nil {
//		return c.RenderArgs["user"].(*models.User)
//	}
//	if username, ok := c.Session["user"]; ok {
//		return c.getUser(username)
//	}te(
//	return nil
//}

func (c App) getUser(username string) *models.User {
	user := new(models.User)
	db.First(&user, "username = ?", username)
	if user == nil {
		return nil
	}
	return user
}

//func (c App) Login(username, password string, remember bool) revel.Result {
//	user := c.getUser(username)
//	if user != nil {
//		err := bcrypt.CompareHashAndPassword(user.HashedPassword, []byte(password))
//		if err != nil {
//			c.Session["user"] = username
//			if remember {
//				c.Session.SetDefaultExpiration()
//			} else {
//				c.Session.SetNoExpiration()
//			}
//			c.Flash.Success("ยินดีต้อนรับ "+ username)
//			return  c.Redirect(User.Index)
//		}
//	}
//	return c.Redirect(App.Login)
//}

func (c App) Index() revel.Result {
	//if c.connected() != nil {
	//	return c.Redirect(routes.User.Index())
	//}
	c.Flash.Error("กรุณา login ก่อน")
	return c.Render()
}

func (c App) Login(username, password string, remember bool) revel.Result {
	user := c.getUser(username)
	if user != nil {
		err := bcrypt.CompareHashAndPassword(user.HashedPassword, []byte(password))
		if err == nil {
			c.Session["user"] = username
			if remember {
				c.Session.SetDefaultExpiration()
			} else {
				c.Session.SetNoExpiration()
			}
			c.Flash.Success("Welcome, " + username)
			return c.Redirect(User.Index)
		}
	}

	c.Flash.Out["username"] = username
	c.Flash.Error("Login failed")
	return c.Redirect(App.Index)
}