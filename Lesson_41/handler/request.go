package handler

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

func readResponse(resp *http.Response) {
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("error reading response body:", err)
		return
	}

	fmt.Println("Response Status:", resp.Status)
	fmt.Println("Response Body:", string(body))
}

func Get() {
	url := "http://localhost:8080/person/5"
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("error fetching URL: ", err)
		return
	}
	defer resp.Body.Close()

	readResponse(resp)
}

func Post() {
	url := "http://localhost:8080/person"
	data := `{"name": "Akbar", "age": 19, "marital_status": "not married"}`
	resp, err := http.Post(url, "application/json", strings.NewReader(data))
	if err != nil {
		fmt.Println("error making POST request:", err)
		return
	}
	defer resp.Body.Close()

	readResponse(resp)
}

func Put() {
	url, err := url.Parse("http://localhost:8080/person/7")
	if err != nil {
		fmt.Println("error parsing URL: ", err)
		return
	}
	data := []byte(`{"name":"Akbar", "age": 20, "marital_status": "married"}`)

	client := http.Client{}
	req, err := http.NewRequest("PUT", url.String(), bytes.NewBuffer(data))
	if err != nil {
		fmt.Println("error creating PUT request:", err)
		return
	}
	req.Header.Set("Content-Type", "application/json")
	
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("error making PUT request:", err)
		return
	}
	defer resp.Body.Close()

	readResponse(resp)
}

func Delete() {
	url, err := url.Parse("http://localhost:8080/person/7")
	if err != nil {
		fmt.Println("error parsing URL: ", err)
		return
	}

	client := http.Client{}
	req, err := http.NewRequest("DELETE", url.String(), nil)
	if err != nil {
		fmt.Println("error creating PUT request:", err)
		return
	}

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("error making DELETE request:", err)
		return
	}
	defer resp.Body.Close()

	readResponse(resp)
}

func GetAll() {
	url := "http://localhost:8080/people"
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("error fetching URL: ", err)
		return
	}
	defer resp.Body.Close()

	readResponse(resp)
}