package model

type User struct {
	Id       int    `json:"id" example:"1"`
	Name     string `json:"name" example:"ZAK"`
	Username string `json:"username" binding:"required" example:"ZAK"`
	Password string `json:"password" binding:"required" example:"qwerty"`
}

type UserExampleRegistration struct {
	Username string `json:"username" binding:"required" example:"ZAK"`
	Password string `json:"password" binding:"required" example:"qwerty"`
}
