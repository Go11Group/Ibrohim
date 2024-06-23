package model

type Transaction struct {
    ID         string `json:"id"`
    CardID     string `json:"card_id"`
    Amount     int    `json:"amount"`
    TerminalID string `json:"terminal_id,omitempty"` // omitempty will omit the field if it's empty (null in JSON)
    Type       string `json:"type"`
}

type TransactionFilter struct {
	CardID        string
	Amount        int
	TerminalID    string
	Type          string
	Limit, Offset int
}
