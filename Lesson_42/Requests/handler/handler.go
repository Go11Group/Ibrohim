package handler

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

func NewRoute() *http.Server {
	mux := http.NewServeMux()

	mux.HandleFunc("POST /users", userPost)
	mux.HandleFunc("GET /users/", userGet)
	mux.HandleFunc("PUT /users/", userPut)
	mux.HandleFunc("DELETE /users/", userDelete)
	mux.HandleFunc("GET /users-all", userGetAll)
	mux.HandleFunc("GET /users-courses/", userCoursesGet)
	mux.HandleFunc("GET /users-search", userSearch)

	mux.HandleFunc("POST /courses", coursePost)
	mux.HandleFunc("GET /courses/", courseGet)
	mux.HandleFunc("PUT /courses/", coursePut)
	mux.HandleFunc("DELETE /courses/", courseDelete)
	mux.HandleFunc("GET /courses-all", courseGetAll)
	mux.HandleFunc("GET /courses-lessons/", courseLessons)
	mux.HandleFunc("GET /courses-enrollments/", courseUsers)
	mux.HandleFunc("GET /courses-popular", coursePopular)

	mux.HandleFunc("POST /lessons", lessonPost)
	mux.HandleFunc("GET /lessons/", lessonGet)
	mux.HandleFunc("PUT /lessons/", lessonPut)
	mux.HandleFunc("DELETE /lessons/", lessonDelete)
	mux.HandleFunc("GET /lessons-all", lessonGetAll)

	mux.HandleFunc("POST /enrollments", enrollmentPost)
	mux.HandleFunc("GET /enrollments/", enrollmentGet)
	mux.HandleFunc("PUT /enrollments/", enrollmentPut)
	mux.HandleFunc("DELETE /enrollments/", enrollmentDelete)
	mux.HandleFunc("GET /enrollments-all", enrollmentGetAll)

	return &http.Server{Addr: "localhost:8081", Handler: mux}
}

func readResponse(w http.ResponseWriter, resp *http.Response) {
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, "error reading response: "+err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(resp.StatusCode)
	w.Write(body)
}

func parseFormData(r *http.Request) (*bytes.Buffer, string, error) {
	var jsonData map[string]interface{}
    err := json.NewDecoder(r.Body).Decode(&jsonData)
    if err != nil {
        return nil, "", errors.New("error parsing JSON: "+ err.Error())
    }

    body := &bytes.Buffer{}
    err = json.NewEncoder(body).Encode(jsonData)
	if err != nil {
		return nil, "", errors.New("error encoding JSON to buffer: "+ err.Error())
	}
	return body, "application/json", nil
}