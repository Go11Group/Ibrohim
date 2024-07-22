package handler

import (
	"auth-service/pkg/logger"
	"auth-service/service"
	"database/sql"
	"log/slog"
)

type Handler struct {
	Auth *service.AuthService
	Log  *slog.Logger
}

func NewHandler(db *sql.DB) *Handler {
	return &Handler{
		Auth: service.NewAuthService(db),
		Log:  logger.NewLogger(),
	}
}
