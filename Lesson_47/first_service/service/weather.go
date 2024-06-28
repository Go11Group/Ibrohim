package service

import (
	"context"
	"database/sql"
	pb "weather-transport/genproto/weather"
	"weather-transport/storage/postgres"

	"github.com/pkg/errors"
)

type weatherService struct {
	pb.UnimplementedWeatherServiceServer
	Repo *postgres.WeatherRepo
}

func NewWeatherService(db *sql.DB) *weatherService {
	return &weatherService{Repo: postgres.NewWeatherRepo(db)}
}

func (w *weatherService) GetCurrentWeather(ctx context.Context, p *pb.Place) (*pb.Weather, error) {
	if p.Country == "" || p.City == "" {
		return nil, errors.New("empty fields provided")
	}

	resp, err := w.Repo.GetCurrentWeather(p)
	if err != nil {
		return nil, errors.Wrap(err, "error getting current weather")
	}

	return resp, nil
}

func (w *weatherService) GetWeatherForecast(ctx context.Context, f *pb.Forecast) (*pb.Weather, error) {
	if f.Date == "" || f.Place.City == "" {
		return nil, errors.New("empty fields provided")
	}

	resp, err := w.Repo.GetWeatherForecast(f)
	if err != nil {
		return nil, errors.Wrap(err, "error getting weather forecast")
	}

	return resp, nil
}

func (w *weatherService) ReportWeatherCondition(ctx context.Context, p *pb.Place) (*pb.WeatherType, error) {
	if p.Country == "" || p.City == "" {
		return nil, errors.New("empty fields provided")
	}

	resp, err := w.Repo.ReportWeatherCondition(p)
	if err != nil {
		return nil, errors.Wrap(err, "error reporting weather condition")
	}

	return resp, nil
}
