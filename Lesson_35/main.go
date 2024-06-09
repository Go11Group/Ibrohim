package main

import (
	"fmt"
	"gorilla_pg/handler"
	"gorilla_pg/storage/postgres"
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

	// PopulateUsers(uRepo)
	// PopulateProblems(pRepo)
	// PopulateUserProblems(uRepo, pRepo, upRepo)

	fmt.Println("Server is listening on port 8080...")
	err = server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}