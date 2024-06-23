package model

type Station struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type StationFilter struct {
	Name          string
	Limit, Offset int
}
