package handler

import (
	"database/sql"
	"metro-service/storage/postgres"
)

type Handler struct {
	Card        *postgres.CardRepo
	Station     *postgres.StationRepo
	Terminal    *postgres.TerminalRepo
	Transaction *postgres.TransactionRepo
}

func NewHandler(db *sql.DB) *Handler {
	return &Handler{
		Card:        postgres.NewCardRepo(db),
		Station:     postgres.NewStationRepo(db),
		Terminal:    postgres.NewTerminalRepo(db),
		Transaction: postgres.NewTransactionRepo(db),
	}
}
