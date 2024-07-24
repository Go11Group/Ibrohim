package main

import (
	"auth-service/api"
	"auth-service/config"
	pba "auth-service/genproto/admin"
	pbu "auth-service/genproto/user"
	"auth-service/service"
	"auth-service/storage/postgres"
	"database/sql"
	"log"
	"net"
	"sync"

	"google.golang.org/grpc"
)

func main() {
	db, err := postgres.ConnectDB()
	if err != nil {
		log.Fatalf("error while connecting to database: %v", err)
	}
	cfg := config.Load()

	var wg *sync.WaitGroup
	wg.Add(2)

	go RunService(wg, db, cfg)
	go RunRouter(wg, db, cfg)

	wg.Wait()
}

func RunService(wg *sync.WaitGroup, db *sql.DB, cfg *config.Config) {
	defer wg.Done()

	lis, err := net.Listen("tcp", cfg.AUTH_SERVICE_PORT)
	if err != nil {
		log.Fatalf("error while listening: %v", err)
	}
	defer lis.Close()

	u := service.NewUserService(db)
	a := service.NewAdminService(db)
	server := grpc.NewServer()
	pbu.RegisterUserServer(server, u)
	pba.RegisterAdminServer(server, a)

	log.Printf("Service is listening on port %s...\n", cfg.AUTH_SERVICE_PORT)
	if err := server.Serve(lis); err != nil {
		log.Fatalf("error while serving auth service: %s", err)
	}
}

func RunRouter(wg *sync.WaitGroup, db *sql.DB, cfg *config.Config) {
	defer wg.Done()

	r := api.NewRouter(db)
	log.Printf("Router is running on port %s...\n", cfg.AUTH_ROUTER_PORT)
	r.Run(config.Load().AUTH_ROUTER_PORT)
}
