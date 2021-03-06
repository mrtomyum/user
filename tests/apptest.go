package tests

import (
	"fmt"
	"github.com/revel/revel/testing"
	"github.com/mrtomyum/user/app/controllers"
	"github.com/mrtomyum/user/app/models"
)

var dbType string = "sqlite3"
var dbFile string = "./app/models/user.db"
var db = controllers.InitDB(dbType, dbFile)

type AppTest struct {
	testing.TestSuite
}

func (t *AppTest) Before() {
	println("Set up")
}

func (t *AppTest) TestThatIndexPageWorks() {
	t.Get("/")
	t.AssertOk()
	t.AssertContentType("text/html; charset=utf-8")
}

func (t *AppTest) After() {
	println("Tear down")
}

func (t *AppTest) TestCreateTableUser() {
	db.CreateTable(&models.User{})
}
func (t *AppTest) TestDropTableUser() {
	db.DropTableIfExists(&models.User{})
}

func (t *AppTest) TestMockUserTable() {
	users := models.MockUser()
	for _, u := range users {
		db.Create(&u)
	}
}

func (t *AppTest) TestAddUserAndSetPass() {
	u := models.User{Name: "nadya"}
	db.Debug().Create(&u)

	fmt.Println("u = ", &u)
	err := u.SetPass("1234")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Set Password OK...~!")
	}
	fmt.Println("u = ", &u)
	db.Save(&u)

}
func (t *AppTest) TestPasswordMismatch() {
	b := models.User{}
	db.Debug().First(&b)
	fmt.Println("b =", &b)
	pass := "123"
	err := b.VerifyPass(pass)

	if err != nil {
		fmt.Println("Password missmatch?? =>", pass)
	} else {
		fmt.Println("Password matched!")
	}
}
