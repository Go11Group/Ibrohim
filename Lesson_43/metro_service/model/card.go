package model

type Card struct {
	ID     string `json:"id"`
	Number string `json:"number"`
	UserID string `json:"user_id"`
}

type CardFilter struct {
	UserID        string
	Limit, Offset int
}
