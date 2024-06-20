package model

type Person struct {
	ID            int    `json:"id"`
	Name          string `json:"name"`
	Age           int    `json:"age"`
	MaritalStatus string `json:"marital_status"`
}
