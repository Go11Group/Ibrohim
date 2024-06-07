package main

import (
	"fmt"
	"http_pg/handler"
	"http_pg/storage/postgres"
)

func main() {
	db, err := postgres.ConnectDB()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	uRepo := postgres.NewUserRepo(db)
	pRepo := postgres.NewProductRepo(db)
	upRepo := postgres.NewUserProductRepo(db)
	server := handler.NewHandler(handler.Handler{User: uRepo, Product: pRepo, UserProduct: upRepo})
	
	fmt.Println("Server is listening on port 8080...")
	err = server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}