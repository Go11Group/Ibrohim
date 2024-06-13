package handler

import (
	"encoding/json"
	"fmt"
	"gorilla_pg/model"
	"net/http"
	"strconv"
	"github.com/gorilla/mux"
)

func GetProblemID(w http.ResponseWriter, r *http.Request) (int, error) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		http.Error(w, "Invalid problem ID", http.StatusBadRequest)
		fmt.Println("Error: ", err)
		return 0, err
	}
	return id, nil
}

func ReadProblemBody(w http.ResponseWriter, r *http.Request, p model.Problem) error {
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		http.Error(w, "Invalid JSON data", http.StatusBadRequest)
		fmt.Println("Error: ", err)
		return err
	}
	return nil
}

func (h * Handler) problemGet(w http.ResponseWriter, r *http.Request) {
	id, err := GetProblemID(w, r)
	if id == 0 || err != nil {
		return
	}
	problems, err := h.Problem.GetProblem(model.Problem{ID: id})
	if err != nil {
		http.Error(w, "Problem not found", http.StatusNotFound)
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
	err := ReadProblemBody(w, r, p)
	if err != nil {
		return
	}
	err = h.Problem.CreateProblem(p)
	if err != nil {
		http.Error(w, "Failed to create problem", http.StatusInternalServerError)
		fmt.Println("Error: ", err)
		return
	}
	fmt.Fprintf(w, "New problem inserted to database")
}

func (h *Handler) problemPut(w http.ResponseWriter, r *http.Request) {
	id, err := GetProblemID(w, r)
	if id == 0 || err != nil {
		return
	}
	p := model.Problem{ID: id}
	err = ReadProblemBody(w, r, p)
	if err != nil {
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
	id, err := GetProblemID(w, r)
	if id == 0 || err != nil {
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