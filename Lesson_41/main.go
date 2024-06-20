package main

import "person_request_response/storage/postgres"

func main() {
	db, err := postgres.ConnectDB()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	
}