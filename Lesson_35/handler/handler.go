package handler

import (
	"gorilla_pg/storage/postgres"
	"net/http"
	"github.com/gorilla/mux"
)

type Handler struct {
	User *postgres.UserRepo
	Problem *postgres.ProblemRepo
	UserProblem *postgres.UserProblemRepo
}

func NewHandler(h Handler) *http.Server {
	r := mux.NewRouter()
	sr := r.PathPrefix("/gorilla").Subrouter()
	sr.HandleFunc("/user/{id}", h.userGet).Methods("GET")
	sr.HandleFunc("/user/", h.userPost).Methods("POST")
	sr.HandleFunc("/user/{id}", h.userPut).Methods("PUT")
	sr.HandleFunc("/user/{id}", h.userDelete).Methods("DELETE")

	sr.HandleFunc("/problem/{id}", h.problemGet).Methods("GET")
	sr.HandleFunc("/problem/", h.problemPost).Methods("POST")
	sr.HandleFunc("/problem/{id}", h.problemPut).Methods("PUT")
	sr.HandleFunc("/problem/{id}", h.problemDelete).Methods("DELETE")

	sr.HandleFunc("/user-problems/{id}", h.userProblemGet).Methods("GET")
	sr.HandleFunc("/user-problems/", h.userProblemPost).Methods("POST")
	sr.HandleFunc("/user-problems/", h.userProblemPut).Methods("PUT")
	sr.HandleFunc("/user-problems/", h.userProblemDelete).Methods("DELETE")
	return &http.Server{Addr: ":8080", Handler: r}
}