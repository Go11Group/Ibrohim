package model

type Filter struct {
	ID int `form:"id"`
	First_name string `form:"first-name"`
	Last_name string `form:"last-name"`
	Age int `form:"age"`
	Nationality string `form:"nationality"`
	Field string `form:"field"`
	City string `form:"city"`
	Limit int `form:"limit"`
	Offset int `form:"offset"`
}