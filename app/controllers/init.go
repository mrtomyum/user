package controllers

import "github.com/revel/revel"

func init() {

	//revel.OnAppStart(InitDB)
	revel.InterceptMethod(App.AddUser, revel.BEFORE)
	revel.InterceptMethod(Users.checkUser, revel.BEFORE)

}
