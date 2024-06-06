package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	url := "http://localhost:8080/get"
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}
	fmt.Println("Response Body:", string(body))

	queryParams := resp.Request.URL.Query()
	fmt.Println("Query Parameters:")
	for k,v := range queryParams {
		fmt.Printf("%s: %s\n", k,v)
	}
}