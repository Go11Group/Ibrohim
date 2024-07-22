package main

import (
	"context"
	"fmt"
	"log"
	"redis-crud/models"
	"redis-crud/redis"
)

func main() {
	db, err := redis.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}

	pRepo := redis.NewPersonRepo(db)

	p := models.PersonInfo{
		Name:      "John Doe",
		Age:       21,
		IsMarried: false,
	}

	res, err := pRepo.Add(context.Background(), &p)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Add result: %v\n", *res)

	res2, err := pRepo.Read(context.Background(), res.ID)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Person: %v\n", *res2)
}
