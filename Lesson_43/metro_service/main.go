package main

import (
	"fmt"
	"metro-service/api"
	"metro-service/storage/postgres"
)

func main() {
	db, err := postgres.ConnectDB()
	if err != nil {
		panic(fmt.Errorf("error with database connection: %v", err))
	}

	server := api.Routes(db)
	fmt.Println("Server is listening on port 8082...")
	panic(server.ListenAndServe())
}
