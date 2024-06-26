package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"
)

func QueryParameters(w http.ResponseWriter, r *http.Request) (int, int, *time.Time, error) {
	queryParams := r.URL.Query()
	userID, err := strconv.Atoi(queryParams.Get("user-id"))
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		fmt.Println("Error: ", err)
		return 0, 0, nil, err
	}
	problemID, err := strconv.Atoi(queryParams.Get("problem-id"))
	if err != nil {
		http.Error(w, "Invalid problem ID", http.StatusBadRequest)
		fmt.Println("Error: ", err)
		return 0, 0, nil, err
	}
	solvedAt, err := time.Parse("02/01/2006 15:04:05", queryParams.Get("solved-at"))
	if err != nil {
		http.Error(w, "Invalid time of solution", http.StatusBadRequest)
		fmt.Println("Error: ", err)
		return 0, 0, nil, err
	}
	return userID, problemID, &solvedAt, nil
}

func (h * Handler) userProblemGet(w http.ResponseWriter, r *http.Request) {
	userID, err := GetUserID(w, r)
	if userID == 0 || err != nil {
		return
	}
	problems, err := h.UserProblem.GetUserProblems(userID)
	if err != nil {
		http.Error(w, "User_Problems not found", http.StatusNotFound)
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

func (h *Handler) userProblemPost(w http.ResponseWriter, r *http.Request) {
	userID, problemID, solvedAt, err := QueryParameters(w, r)
	if err != nil || userID == 0 || problemID == 0 || solvedAt == nil {
		return
	}
	err = h.UserProblem.AddProblemToUser(userID, problemID, *solvedAt)
	if err != nil {
		http.Error(w, "Failed to add problem to user", http.StatusInternalServerError)
		fmt.Println("Error: ", err)
		return
	}
	fmt.Fprintf(w, "UserID %d solved the %dth problem", userID, problemID)
}

func (h *Handler) userProblemPut(w http.ResponseWriter, r *http.Request) {
	userID, problemID, solvedAt, err := QueryParameters(w, r)
	if err != nil || userID == 0 || problemID == 0 || solvedAt == nil {
		return
	}
	err = h.UserProblem.UpdateTimeOfSolution(userID, problemID, *solvedAt)
	if err != nil {
		http.Error(w, "Failed to update time of solution", http.StatusInternalServerError)
		fmt.Println("Error: ", err)
		return
	}
	fmt.Fprintf(w, "The %dth problem time of solution for UserID %d updated", problemID, userID)
}

func (h *Handler) userProblemDelete(w http.ResponseWriter, r *http.Request) {
	userID, problemID, _, err := QueryParameters(w, r)
	if err != nil || userID == 0 || problemID == 0 {
		return
	}
	err = h.UserProblem.RemoveProblemFromUser(userID, problemID)
	if err != nil {
		http.Error(w, "Failed to delete problem from user", http.StatusInternalServerError)
		fmt.Println("Error: ", err)
		return
	}
	fmt.Fprintf(w, "The %dth problem removed from UserID %d", problemID, userID)
}