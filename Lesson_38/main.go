package main

import (
	"fmt"
	"language_learning_app/api/handler"
	"language_learning_app/storage/postgres"
)

func main() {
	db, err := postgres.ConnectDB()
	if err != nil {
		panic(err)
	}
	uRepo := postgres.NewUserRepo(db)

	server := handler.NewHandler(handler.Handler{User: uRepo})
	fmt.Println("Server is listening on port 8080...")
	err = server.Run(":8080")
	if err != nil {
		fmt.Println(fmt.Errorf("error with running server: %v", err))
		return
	}
}