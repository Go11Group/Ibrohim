package handler

import (
	"api-gateway/config"
	pba "api-gateway/genproto/admin"
	pbu "api-gateway/genproto/user"
	"api-gateway/pkg"
	"api-gateway/pkg/logger"
	"log/slog"
)

type Handler struct {
	UserClient  pbu.UserClient
	AdminClient pba.AdminClient
	Log         *slog.Logger
}

func NewHandler(cfg *config.Config) *Handler {
	return &Handler{
		UserClient:  pkg.NewUserClient(cfg),
		AdminClient: pkg.NewAdminClient(cfg),
		Log:         logger.NewLogger(),
	}
}
