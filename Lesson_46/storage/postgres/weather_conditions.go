package postgres

import (
	"database/sql"
	pb "weather-transport/genproto/weather"

	"github.com/pkg/errors"
)

type WeatherRepo struct {
	DB *sql.DB
}

func NewWeatherRepo(db *sql.DB) *WeatherRepo {
	return &WeatherRepo{DB: db}
}

func (wr *WeatherRepo) GetCurrentWeather(p *pb.Place) (*pb.Weather, error) {
	w := &pb.Weather{}

	query := "select temperature, humidity, wind_speed from weather_conditions where city = $1 and date = now()"
	err := wr.DB.QueryRow(query, p.City).Scan(&w.Temperature, &w.Humidity, &w.WindSpeed)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.Wrap(err, "city not found")
		}
		return nil, errors.Wrap(err, "failed to get current weather conditions")
	}

	return w, nil
}

func (wr *WeatherRepo) GetWeatherForecast(f *pb.Forecast) (*pb.Weather, error) {
	w := &pb.Weather{}

	query := "select temperature, humidity, wind_speed from weather_conditions where city = $1 and date = $2"
	err := wr.DB.QueryRow(query, f.Place.City, f.Date).Scan(&w.Temperature, &w.Humidity, &w.WindSpeed)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.Wrap(err, "forecast not found")
		}
		return nil, errors.Wrap(err, "failed to get forecast")
	}

	return w, nil
}

func (wr *WeatherRepo) ReportWeatherCondition(p *pb.Place) (*pb.WeatherType, error) {
	w := &pb.WeatherType{}
	query := "select weather_type from weather_conditions where city = $1 and date = now()"
	err := wr.DB.QueryRow(query, p.City).Scan(&w.Type)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.Wrap(err, "weather condition not found")
		}
		return nil, errors.Wrap(err, "failed to query weather condition")
	}

	return w, nil
}
