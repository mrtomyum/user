package models

func MockUser()  []*User{
	users := []*User{
		{Name:"เกษม อานนทวิลาศ", Username:"tom", Role:"System Admin"},
		{Name:"จิราภรณ์ อานนทวิลาศ", Username:"jip", Role:"Manager"},
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
		"abc",
	}
	for k, _ := range users{
		users[k].SetPass(passwords[k])
	}
	return users
}
