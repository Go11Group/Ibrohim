package main

import (
	"api-gateway/api"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.NewClient(":50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to connect to the address: %v", err)
	}
	defer conn.Close()

	r := api.NewRouter(conn)
	panic(r.Run())
}
