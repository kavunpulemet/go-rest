package auth

type User struct {
	Id       int    `json:"-" db:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type UsersList struct {
	Id     int
	UserId int
	ListId int
}
