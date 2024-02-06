package model

type User struct {
	Model
	Username string
	Email    string
}

func (u *User) TableName() string {
	return "users"
}
