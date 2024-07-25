package models

type NewLocation struct {
	UserId     string `json:"user_id"`
	Address    string `json:"address"`
	City       string `json:"city"`
	State      string `json:"state"`
	Country    string `json:"country"`
	PostalCode string `json:"postal_code"`
}

type LocationDetails struct {
	Id         string `json:"id"`
	UserId     string `json:"user_id"`
	Address    string `json:"address"`
	City       string `json:"city"`
	State      string `json:"state"`
	Country    string `json:"country"`
	PostalCode string `json:"postal_code"`
	CreatedAt  string `json:"created_at"`
}

type UpdateResp struct {
	Id         string `json:"id"`
	UserId     string `json:"user_id"`
	Address    string `json:"address"`
	City       string `json:"city"`
	State      string `json:"state"`
	Country    string `json:"country"`
	PostalCode string `json:"postal_code"`
	UpdatedAt  string `json:"updated_at"`
}
