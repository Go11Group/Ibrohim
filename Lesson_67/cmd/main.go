package main

import (
	"log"
	"net"
	"product-service/config"
	mongodb "product-service/storage/mongoDB"

	"google.golang.org/grpc"
)

func main() {
	cfg := config.Load()
	db, err := mongodb.ConnectDB()
	if err != nil {
		log.Fatalf("error while connecting to database: %v", err)
	}
	defer db.Close()

	lis, err := net.Listen("tcp", cfg.PRODUCT_SERVICE_PORT)
	if err != nil {
		log.Fatalf("error while listening: %v", err)
	}
	defer lis.Close()

	// u := service.NewUserService(db)
	// a := service.NewAdminService(db)
	server := grpc.NewServer()
	// pbu.RegisterUserServer(server, u)
	// pba.RegisterAdminServer(server, a)

	log.Printf("Service is listening on port %s...\n", cfg.PRODUCT_SERVICE_PORT)
	if err := server.Serve(lis); err != nil {
		log.Fatalf("error while serving product service: %s", err)
	}
}
