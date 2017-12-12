package entity

type User struct {
	Name     string
	Password string
	Email    string
	Phone    string
}

func NewUser(name string, password string, email string, phone string) *User {
	return &User{name, password, email, phone}
}

// the return result contains  user id, so i make another one
type RetUser struct {
	Id       int    `json:"id"`
	Name     string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
}

func NewRetUser(name string, password string, email string, phone string) *RetUser {
	return &RetUser{1, name, password, email, phone}
}

type UserInfo struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
