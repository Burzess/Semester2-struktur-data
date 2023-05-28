package usercontroller

import "progres/model/user"

func Login(username, password string) string {
	return user.Login(username, password)
}
