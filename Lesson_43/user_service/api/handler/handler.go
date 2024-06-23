package handler

import (
	"database/sql"
	"user-service/storage/postgres"
)

type Handler struct {
	User *postgres.UserRepo
}

func NewHandler(db *sql.DB) *Handler {
	return &Handler{User: postgres.NewUserRepo(db)}
}
