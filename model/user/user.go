package user

import (
	"progres/db"
	entity "progres/entity/users"
)

func CreateAccount(newData entity.User) {
	newUser := entity.LinkedlistUsers{}
	newUser.Data = newData
	temp := &db.DBUsers
	if temp.Next == nil {
		temp.Next = &newUser
	} else {
		for temp.Next != nil {
			temp = temp.Next
		}
		temp.Next = &newUser
	}
}

func DataStatisUser() {
	administrator := entity.User{
		Username: "halim",
		Password: "halim123",
		Status:   "administrator",
	}

	petugas_loket := entity.User{
		Username: "alif",
		Password: "alif123",
		Status:   "petugas loket",
	}

	CreateAccount(administrator)
	CreateAccount(petugas_loket)
}

func Login(username, password string) string {
	temp := &db.DBUsers
	for temp != nil {
		if username == temp.Data.Username && password == temp.Data.Password {
			return temp.Data.Status
		}
		temp = temp.Next
	}
	return ""
}
