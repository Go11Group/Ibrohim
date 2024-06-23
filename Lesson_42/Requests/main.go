package main

import (
	"fmt"
	"github.com/golangN11/Ibrohim/Lesson_42/Requests/handler"
)

func main() {
	server := handler.NewRoute()
	
	fmt.Println("Server is listening on port 8081...")
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}