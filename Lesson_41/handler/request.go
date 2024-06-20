package handler

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func Get() {
	url := "http://localhost:8080/person/5"
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("error fetching URL: ", err)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("error reading response body:", err)
		return
	}

	fmt.Println("Response Status:", resp.Status)
	fmt.Println("Response Body:")
	fmt.Println(string(body))
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

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("error reading response body:", err)
		return
	}
	fmt.Println("Response Status:", resp.Status)
	fmt.Println("Response Body:", string(body))
}