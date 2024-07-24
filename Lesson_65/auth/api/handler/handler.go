package handler

import (
	"auth-service/pkg/logger"
	"auth-service/storage/postgres"
	"database/sql"
	"log/slog"
)

type Handler struct {
	RepoAdmin *postgres.AdminRepo
	RepoUser  *postgres.UserRepo
	RepoToken *postgres.TokenRepo
	Log       *slog.Logger
}

func NewHandler(db *sql.DB) *Handler {
	return &Handler{
		RepoAdmin: postgres.NewAdminRepo(db),
		RepoUser:  postgres.NewUserRepo(db),
		RepoToken: postgres.NewTokenRepo(db),
		Log:       logger.NewLogger(),
	}
}
