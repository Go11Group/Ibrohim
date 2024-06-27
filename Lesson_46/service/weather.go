package service

import (
	"database/sql"
	pb "weather-transport/genproto/weather"
	"weather-transport/storage/postgres"
)

type weatherService struct {
	pb.UnimplementedWeatherServiceServer
	Repo *postgres.WeatherRepo
}

func NewWeatherService(db *sql.DB) *weatherService {
	return &weatherService{Repo: postgres.NewWeatherRepo(db)}
}

