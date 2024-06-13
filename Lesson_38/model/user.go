package model

type User struct {
	UserID    string `form:"user_id"`
    Name      string `form:"name"`
    Birthday  string `form:"birthday"`
    Email     string `form:"email"`
    Password  string `form:"password"`
}

type UserFilter struct {
    Name string
    AgeFrom int
    AgeTo int
    Limit int
    Offset int
}