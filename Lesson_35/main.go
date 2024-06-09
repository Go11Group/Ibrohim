package main

import "gorilla_pg/storage/postgres"

func main() {
	db, err := postgres.ConnectDB()
	if err != nil {
		panic(err)
	}
	defer db.Close()
}