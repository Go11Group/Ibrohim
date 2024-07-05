package handler

import (
	"user_service/storage/postgres"
)

type Handler struct {
	UserRepo *postgres.UserRepo
}

func NewHandler() *Handler {
	return &Handler{}
}
