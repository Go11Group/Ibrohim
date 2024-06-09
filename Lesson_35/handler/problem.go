package handler

import (
	"encoding/json"
	"fmt"
	"gorilla_pg/model"
	"net/http"
	"strconv"
	"github.com/gorilla/mux"
)

func (h * Handler) problemGet(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		http.Error(w, "Invalid product ID", http.StatusBadRequest)
		fmt.Println("Error: ", err)
		return
	}
	problems, err := h.Problem.GetProblem(model.Problem{ID: id})
	if err != nil {
		http.Error(w, "Product not found", http.StatusNotFound)
		fmt.Println("Error: ", err)
		return
	}
	err = json.NewEncoder(w).Encode(problems)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		fmt.Println("Error: ", err)
		return
	}
}

func (h *Handler) problemPost(w http.ResponseWriter, r *http.Request) {
	p := model.Problem{}
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		http.Error(w, "Invalid JSON data", http.StatusBadRequest)
		fmt.Println("Error: ", err)
		return
	}
	err = h.Problem.CreateProblem(p)
	if err != nil {
		http.Error(w, "Failed to create problem", http.StatusInternalServerError)
		fmt.Println("Error: ", err)
		return
	}
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "New problem inserted to database")
}

func (h *Handler) problemPut(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		http.Error(w, "Invalid problem ID", http.StatusBadRequest)
		fmt.Println("Error: ", err)
		return
	}
	p := model.Problem{ID: id}
	err = json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		http.Error(w, "Invalid JSON data", http.StatusBadRequest)
		fmt.Println("Error: ", err)
		return
	}
	err = h.Problem.UpdateProblem(p)
	if err != nil {
		http.Error(w, "Failed to update problem", http.StatusInternalServerError)
		fmt.Println("Error: ", err)
		return
	}
	fmt.Fprintf(w, "Problem with ID %d updated", p.ID)
}

func (h *Handler) problemDelete(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		http.Error(w, "Invalid problem ID", http.StatusBadRequest)
		fmt.Println("Error: ", err)
		return
	}
	err = h.Problem.DeleteProblem(id)
	if err != nil {
		http.Error(w, "Failed to delete problem", http.StatusInternalServerError)
		fmt.Println("Error: ", err)
		return
	}
	fmt.Fprintf(w, "Problem with ID %d deleted", id)
}