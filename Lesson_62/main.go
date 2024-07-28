package main

import (
	"log"
	"redis-crud/api"
	"redis-crud/redis"
)

func main() {
	db, err := redis.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}

	router := api.NewRouter(db)

	router.Run(":8080")
}
