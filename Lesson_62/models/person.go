package models

type Person struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Age       int    `json:"age"`
	IsMarried bool   `json:"is_married"`
}

type PersonInfo struct {
	Name      string `json:"name"`
	Age       int    `json:"age"`
	IsMarried bool   `json:"is_married"`
}
