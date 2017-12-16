package entities

type User struct {
	UID      int    `orm:"id,auto-inc" json:"id"`
	Name     string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
}

func NewUser(user User) *User {
	if len(user.Name) == 0 {
		panic("name should be not null!")
	}
	return &user
}

type UserInfo struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
