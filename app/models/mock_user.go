package models

func MockUser()  []*User{
	users := []*User{
		{Name:"tom"},
		{Name:"jip"},
		{Name:"tam"},
		{Name:"satit"},
		{Name:"somrod"},
		{Name:"bee"},
	}
	passwords := []string{
		"1323453",
		"asdfasfd",
		"23rsafasf",
		";alsjdfl",
		"a;dlfjka",
		"a;dlsjkf;;l",
	}
	for k, _ := range users{
		users[k].SetPass(passwords[k])
	}
	return users
}
//func MockPassword() []
