package handler

import (
	"net/http"
	"strings"
)

func userGet(w http.ResponseWriter, r *http.Request) {
	url := "http://localhost:8080/language_learning_app/users/" + strings.TrimPrefix(r.URL.Path, "/users/")

	resp, err := http.Get(url)
	if err != nil {
		http.Error(w, "error fetching URL: "+err.Error(), 500)
		return
	}
	defer resp.Body.Close()

	readResponse(w, resp)
}

func userPost(w http.ResponseWriter, r *http.Request) {
	body, contentType, err := parseFormData(r)
	if err != nil {
		http.Error(w, "error parsing json data: "+err.Error(), 400)
		return
	}

	url := "http://localhost:8080/language_learning_app/users"
	resp, err := http.Post(url, contentType, body)
	if err != nil {
		http.Error(w, "error making POST request: "+err.Error(), 500)
		return
	}
	defer resp.Body.Close()

	readResponse(w, resp)
}

func userPut(w http.ResponseWriter, r *http.Request) {
	body, contentType, err := parseFormData(r)
	if err != nil {
		http.Error(w, "error parsing json data: "+err.Error(), 400)
		return
	}

	url := "http://localhost:8080/language_learning_app/users/" + strings.TrimPrefix(r.URL.Path, "/users/")

	client := http.Client{}
	req, err := http.NewRequest("PUT", url, body)
	if err != nil {
		http.Error(w, "error creating PUT request: "+err.Error(), 500)
		return
	}
	req.Header.Set("Content-Type", contentType)

	resp, err := client.Do(req)
	if err != nil {
		http.Error(w, "error making PUT request: "+err.Error(), 500)
		return
	}
	defer resp.Body.Close()

	readResponse(w, resp)
}

func userDelete(w http.ResponseWriter, r *http.Request) {
	url := "http://localhost:8080/language_learning_app/users/" + strings.TrimPrefix(r.URL.Path, "/users/")

	client := http.Client{}
	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		http.Error(w, "error creating DELETE request: "+err.Error(), 500)
		return
	}

	resp, err := client.Do(req)
	if err != nil {
		http.Error(w, "error making DELETE request: "+err.Error(), 500)
		return
	}
	defer resp.Body.Close()

	readResponse(w, resp)
}

func userGetAll(w http.ResponseWriter, r *http.Request) {
	quearyParams := r.URL.Query()

	url := "http://localhost:8080/language_learning_app/users/get-all"
	if len(quearyParams) > 0 {
		url += "?" + quearyParams.Encode()
	}

	resp, err := http.Get(url)
	if err != nil {
		http.Error(w, "error fetching URL: "+err.Error(), 400)
		return
	}
	defer resp.Body.Close()

	readResponse(w, resp)
}

func userCoursesGet(w http.ResponseWriter, r *http.Request) {
	url := "http://localhost:8080/language_learning_app/users/" + strings.TrimPrefix(r.URL.Path, "/users-courses/") + "/courses"

	resp, err := http.Get(url)
	if err != nil {
		http.Error(w, "error fetching URL: "+err.Error(), 500)
		return
	}
	defer resp.Body.Close()

	readResponse(w, resp)
}

func userSearch(w http.ResponseWriter, r *http.Request) {
	quearyParams := r.URL.Query()

	url := "http://localhost:8080/language_learning_app/users/search"
	if len(quearyParams) > 0 {
		url += "?" + quearyParams.Encode()
	}

	resp, err := http.Get(url)
	if err != nil {
		http.Error(w, "error fetching URL: "+err.Error(), 400)
		return
	}
	defer resp.Body.Close()

	readResponse(w, resp)
}
