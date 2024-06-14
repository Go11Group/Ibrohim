package model

type User struct {
	UserID   string `form:"user_id" json:"user_id"`
	Name     string `form:"name" json:"name"`
	Birthday string `form:"birthday" json:"birthday"`
	Email    string `form:"email" json:"email"`
	Password string `form:"password" json:"password"`
}

// for filtering users
type UserFilter struct {
	Name    string
	AgeFrom int
	AgeTo   int
	Limit   int
	Offset  int
}

// for json responses
type UserResponse struct {
	UserID string `json:"user_id"`
	Name   string `json:"name"`
	Email  string `json:"email"`
}
