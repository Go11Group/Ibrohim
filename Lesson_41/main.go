package main

import (
	"fmt"
	"person_request_response/handler"
	"person_request_response/storage/postgres"
)

func main() {
	db, err := postgres.ConnectDB()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	pRepo := postgres.NewPersonRepo(db)
	server := handler.NewHandler(handler.Handler{Person: pRepo})
	go handler.Post()
	fmt.Println("Server is listening on port 8080...")
	err = server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}