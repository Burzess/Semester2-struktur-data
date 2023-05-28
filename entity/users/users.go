package users

type User struct {
	Username string
	Password string
	Status   string
}

type LinkedlistUsers struct {
	Data User
	Next *LinkedlistUsers
}
