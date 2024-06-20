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
	mux.HandleFunc("GET /person/", h.Get)
	mux.HandleFunc("POST /person", h.Post)
	mux.HandleFunc("PUT /person/", h.Put)
	mux.HandleFunc("DELETE /person/", h.Delete)
	mux.HandleFunc("GET /people", h.GetAll)
	return &http.Server{Addr: "localhost:8080", Handler: mux}
}