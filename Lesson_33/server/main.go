package main

import (
	"fmt"
	"net/http"
)

func main() {
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello, World!")
	})
	filterMethod := func(method string, handler http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Method != method {
				http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
				return
			}
			handler.ServeHTTP(w, r)
		})
	}

	http.Handle("/get", filterMethod("GET", handler))
	http.Handle("/post", filterMethod("POST", handler))
    http.Handle("/put", filterMethod("PUT", handler))
    http.Handle("/delete", filterMethod("DELETE", handler))

	fmt.Println("Server is listening on port 8080...")
	http.ListenAndServe(":8080", nil)
}