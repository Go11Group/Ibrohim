package main

import (
	"les37/handler"
	"les37/postgres"
)

func main() {
	db, err := postgres.ConnectDB()
	if err != nil {
		panic(err)
	}
	pRepo := postgres.NewPersonRepo(db)
	server := handler.NewHandler(handler.Handler{Person: pRepo})
	err = server.Run(":8080")
	if err != nil {
		panic(err)
	}
}