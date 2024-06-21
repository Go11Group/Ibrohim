package main

import (
	"fmt"
	"requests/handler"
)

func main() {
	server := handler.NewRoute()
	
	fmt.Println("Server is listening on port 8081...")
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}