package handler

import (
	"net/http"
	"strings"
)

func enrollmentGet(w http.ResponseWriter, r *http.Request) {
	url := "http://localhost:8080/language_learning_app/enrollments/" + strings.TrimPrefix(r.URL.Path, "/enrollments/")

	resp, err := http.Get(url)
	if err != nil {
		http.Error(w, "error fetching URL: "+err.Error(), 500)
		return
	}
	defer resp.Body.Close()

	readResponse(w, resp)
}

func enrollmentPost(w http.ResponseWriter, r *http.Request) {
	body, contentType, err := parseFormData(r)
	if err != nil {
		http.Error(w, "error parsing json data: "+err.Error(), 400)
		return
	}

	url := "http://localhost:8080/language_learning_app/enrollments"
	resp, err := http.Post(url, contentType, body)
	if err != nil {
		http.Error(w, "error making POST request: "+err.Error(), 500)
		return
	}
	defer resp.Body.Close()

	readResponse(w, resp)
}

func enrollmentPut(w http.ResponseWriter, r *http.Request) {
	body, contentType, err := parseFormData(r)
	if err != nil {
		http.Error(w, "error parsing json data: "+err.Error(), 400)
		return
	}

	url := "http://localhost:8080/language_learning_app/enrollments/" + strings.TrimPrefix(r.URL.Path, "/enrollments/")

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

func enrollmentDelete(w http.ResponseWriter, r *http.Request) {
	url := "http://localhost:8080/language_learning_app/enrollments/" + strings.TrimPrefix(r.URL.Path, "/enrollments/")

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

func enrollmentGetAll(w http.ResponseWriter, r *http.Request) {
	quearyParams := r.URL.Query()

	url := "http://localhost:8080/language_learning_app/enrollments/get-all"
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