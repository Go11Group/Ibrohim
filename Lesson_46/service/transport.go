package service

import (
	"database/sql"
	pb "weather-transport/genproto/transport"
	"weather-transport/storage/postgres"
)

type transportService struct {
	pb.UnimplementedTransportServiceServer
	Repo *postgres.BusRepo
}

func NewTransportService(db *sql.DB) *transportService {
	return &transportService{Repo: postgres.NewBusRepo(db)}
}
