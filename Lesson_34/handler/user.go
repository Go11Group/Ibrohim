package handler

import (
	"encoding/json"
	"fmt"
	"http_pg/model"
	"net/http"
	"strconv"
	"strings"
)

func userGet(h Handler, w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(strings.TrimPrefix(r.URL.Path, "/user/"))
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		fmt.Println("Error: ", err)
		return
	}
	users, err := h.User.GetUser(model.User{ID: id})
	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		fmt.Println("Error: ", err)
		return
	}
	err = json.NewEncoder(w).Encode(users)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		fmt.Println("Error: ", err)
		return
	}
}

func userPost(h Handler, w http.ResponseWriter, r *http.Request) {
	u := model.User{}
	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		http.Error(w, "Invalid JSON data", http.StatusBadRequest)
		fmt.Println("Error: ", err)
		return
	}
	err = h.User.CreateUser(u)
	if err != nil {
		http.Error(w, "Failed to create user", http.StatusInternalServerError)
		fmt.Println("Error: ", err)
		return
	}
	w.Write([]byte("New user inserted to database"))
}

func userPut(h Handler, w http.ResponseWriter, r *http.Request) {
	u := model.User{}
	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		http.Error(w, "Invalid JSON data", http.StatusBadRequest)
		fmt.Println("Error: ", err)
		return
	}
	err = h.User.UpdateUser(u)
	if err != nil {
		http.Error(w, "Failed to update user", http.StatusInternalServerError)
		fmt.Println("Error: ", err)
		return
	}
	fmt.Fprintf(w, "User with ID %d updated", u.ID)
}

func userDelete(h Handler, w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(strings.TrimPrefix(r.URL.Path, "/user/"))
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		fmt.Println("Error: ", err)
		return
	}
	err = h.User.DeleteUser(id)
	if err != nil {
		http.Error(w, "Failed to delete user", http.StatusInternalServerError)
		fmt.Println("Error: ", err)
		return
	}
	fmt.Fprintf(w, "User with ID %d deleted", id)
}