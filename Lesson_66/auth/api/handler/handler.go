package handler

import (
	"auth-service/pkg/logger"
	"auth-service/storage"
	"log/slog"
)

type Handler struct {
	Storage storage.IStorage
	Log     *slog.Logger
}

func NewHandler(s storage.IStorage) *Handler {
	return &Handler{
		Storage: s,
		Log:     logger.NewLogger(),
	}
}
