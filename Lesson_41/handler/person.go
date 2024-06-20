package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"person_request_response/model"
	"strconv"
	"strings"
)

func (h *Handler) Get(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(strings.TrimPrefix(r.URL.Path, "/person/"))
	if err != nil {
		http.Error(w, "invalid person id", http.StatusBadRequest)
		fmt.Println("error: ", err)
		return
	}
	p, err := h.Person.Read(id)
	if err != nil {
		http.Error(w, "person not found", http.StatusNotFound)
		fmt.Println("error: ", err)
		return
	}
	err = json.NewEncoder(w).Encode(p)
	if err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
		fmt.Println("error: ", err)
		return
	}
}

func (h *Handler) Post(w http.ResponseWriter, r *http.Request) {
	var p model.Person
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		http.Error(w, "invalid JSON data", http.StatusBadRequest)
		fmt.Println("error: ", err)
		return
	}
	err = h.Person.Create(p)
	if err != nil {
		http.Error(w, "failed to create person info", http.StatusInternalServerError)
		fmt.Println("error: ", err)
		return
	}
	w.Write([]byte("New person inserted to database"))
}

func (h *Handler) Put(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(strings.TrimPrefix(r.URL.Path, "/person/"))
	if err != nil {
		http.Error(w, "invalid person id", http.StatusBadRequest)
		fmt.Println("error: ", err)
		return
	}
	var p model.Person
	err = json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		http.Error(w, "invalid JSON data", http.StatusBadRequest)
		fmt.Println("error: ", err)
		return
	}
	err = h.Person.Update(id, p)
	if err != nil {
		http.Error(w, "failed to update person", http.StatusInternalServerError)
		fmt.Println("error: ", err)
		return
	}
	fmt.Fprintf(w, "person with id %d updated", id)
}

func (h *Handler) Delete(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(strings.TrimPrefix(r.URL.Path, "/person/"))
	if err != nil {
		http.Error(w, "invalid person id", http.StatusBadRequest)
		fmt.Println("error: ", err)
		return
	}
	err = h.Person.Delete(id)
	if err != nil {
		http.Error(w, "failed to delete person", http.StatusInternalServerError)
		fmt.Println("error: ", err)
		return
	}
	fmt.Fprintf(w, "person with id %d deleted", id)
}

func (h *Handler) GetAll(w http.ResponseWriter, r *http.Request) {
	people, err := h.Person.ReadAll()
	if err != nil {
		http.Error(w, "failed to retrieve people", http.StatusInternalServerError)
		fmt.Println("error: ", err)
		return
	}
	err = json.NewEncoder(w).Encode(people)
	if err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
		fmt.Println("error: ", err)
		return
	}
}