package main

import (
	"fmt"
	"gin_pg/handler"
	"gin_pg/storage/postgres"
)

func main() {
	db, err := postgres.ConnectDB()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	uRepo := postgres.NewUserRepo(db)
	pRepo := postgres.NewProblemRepo(db)
	upRepo := postgres.NewUserProblemRepo(db)
	server := handler.NewHandler(handler.Handler{User: uRepo, Problem: pRepo, UserProblem: upRepo})

	fmt.Println("Server is listening on port 8080...")
	err = server.Run(":8080")
	if err != nil {
		panic(err)
	}
}