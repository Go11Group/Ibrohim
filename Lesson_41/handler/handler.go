package handler

import (
	"net/http"
	"person_request_response/storage/postgres"
)

type Handler struct {
	Person *postgres.PersonRepo
}

func NewHandler(h Handler) *http.Server {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /user", h.Get)
	mux.HandleFunc("POST /user", h.Post)
	mux.HandleFunc("PUT /user", h.Put)
	mux.HandleFunc("DELETE /user", h.Delete)
	mux.HandleFunc("GET /users", h.GetAll)
	return &http.Server{Addr: "localhost:8080", Handler: mux}
}