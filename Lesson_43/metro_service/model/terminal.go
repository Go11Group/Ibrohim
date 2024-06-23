package model

type Terminal struct {
	ID        string `json:"id"`
	StationID string `json:"station_id"`
}

type TerminalFilter struct {
	StationID     string
	Limit, Offset int
}
