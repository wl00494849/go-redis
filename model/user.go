package model

type User struct {
	Id       int    `json:"id"`
	UserName string `json:"userName"`
	Email    string `json:"email"`
	Password []byte `json:"pwd"`
}
