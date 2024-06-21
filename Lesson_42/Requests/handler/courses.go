package handler

import (
	"net/http"
	"strings"
)

func courseGet(w http.ResponseWriter, r *http.Request) {
	url := "http://localhost:8080/language_learning_app/courses/" + strings.TrimPrefix(r.URL.Path, "/courses/")

	resp, err := http.Get(url)
	if err != nil {
		http.Error(w, "error fetching URL: "+err.Error(), 500)
		return
	}
	defer resp.Body.Close()

	readResponse(w, resp)
}

func coursePost(w http.ResponseWriter, r *http.Request) {
	body, contentType, err := parseFormData(r)
	if err != nil {
		http.Error(w, "error parsing json data: "+err.Error(), 400)
		return
	}

	url := "http://localhost:8080/language_learning_app/courses"
	resp, err := http.Post(url, contentType, body)
	if err != nil {
		http.Error(w, "error making POST request: "+err.Error(), 500)
		return
	}
	defer resp.Body.Close()

	readResponse(w, resp)
}

func coursePut(w http.ResponseWriter, r *http.Request) {
	body, contentType, err := parseFormData(r)
	if err != nil {
		http.Error(w, "error parsing json data: "+err.Error(), 400)
		return
	}

	url := "http://localhost:8080/language_learning_app/courses/" + strings.TrimPrefix(r.URL.Path, "/courses/")

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

func courseDelete(w http.ResponseWriter, r *http.Request) {
	url := "http://localhost:8080/language_learning_app/courses/" + strings.TrimPrefix(r.URL.Path, "/courses/")

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

func courseGetAll(w http.ResponseWriter, r *http.Request) {
	quearyParams := r.URL.Query()

	url := "http://localhost:8080/language_learning_app/courses/get-all"
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

func courseLessons(w http.ResponseWriter, r *http.Request) {
	url := "http://localhost:8080/language_learning_app/courses/" + strings.TrimPrefix(r.URL.Path, "/courses-lessons/") + "/lessons"

	resp, err := http.Get(url)
	if err != nil {
		http.Error(w, "error fetching URL: "+err.Error(), 500)
		return
	}
	defer resp.Body.Close()

	readResponse(w, resp)
}

func courseUsers(w http.ResponseWriter, r *http.Request) {
	url := "http://localhost:8080/language_learning_app/courses/" + strings.TrimPrefix(r.URL.Path, "/courses-enrollments/") + "/enrollments"

	resp, err := http.Get(url)
	if err != nil {
		http.Error(w, "error fetching URL: "+err.Error(), 500)
		return
	}
	defer resp.Body.Close()

	readResponse(w, resp)
}

func coursePopular(w http.ResponseWriter, r *http.Request) {
	quearyParams := r.URL.Query()

	url := "http://localhost:8080/language_learning_app/courses/popular"
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