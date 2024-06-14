package main

import (
	"fmt"
	"language_learning_app/api/handler"
	"language_learning_app/storage/postgres"
)

func main() {
	db, err := postgres.ConnectDB() // making a connection to database
	if err != nil {
		panic(fmt.Errorf("error with database connection: %v", err))
	}

	server := handler.NewRoute(*handler.NewHandler(db)) // making a connection to server
	fmt.Println("Server is listening on port 8080...")
	err = server.Run(":8080") // running server
	if err != nil {
		fmt.Println(fmt.Errorf("error with running server: %v", err))
		return
	}
}
