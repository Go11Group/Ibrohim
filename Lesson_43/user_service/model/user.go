package model

type User struct {
	ID          string `form:"id" json:"id"`
	Name        string `form:"name" json:"name"`
	Age         int    `form:"age" json:"age"`
	PhoneNumber string `form:"phone_number" json:"phone_number"`
}

type UserFilter struct {
	Name    string
	AgeFrom int
	AgeTo   int
	Limit   int
	Offset  int
}
