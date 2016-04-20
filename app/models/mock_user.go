package models

func MockUser()  []*User{
	users := []*User{
		{Name:"เกษมอานนทวิลาศ", Username:"tom"},
		{Name:"จิราภรณ์ อานนทวิลาศ", Username:"jip"},
		{Name:"ธนันท์ อานนทวิลาศ", Username:"tam"},
		{Name:"สาธิต โฉมวัฒนา", Username:"satit"},
		{Name:"สมรถ หลักฐาน", Username:"somrod"},
		{Name:"บีบี", Username:"bee"},
	}
	passwords := []string{
		"1234",
		"1234",
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
