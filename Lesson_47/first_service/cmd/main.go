package main

import (
	"log"
	"net"
	tpb "weather-transport/genproto/transport"
	wpb "weather-transport/genproto/weather"
	"weather-transport/service"
	"weather-transport/storage/postgres"

	"google.golang.org/grpc"
)

func main() {
	db, err := postgres.ConnectDB()
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	defer lis.Close()

	server := grpc.NewServer()
	wpb.RegisterWeatherServiceServer(server, service.NewWeatherService(db))
	tpb.RegisterTransportServiceServer(server, service.NewTransportService(db))

	log.Println("Server is listening on port 50051...")
	if err := server.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
