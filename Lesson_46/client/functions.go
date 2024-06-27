package main

import (
	"context"
	"fmt"
	"log"
	tpb "weather-transport/genproto/transport"
	wpb "weather-transport/genproto/weather"
)

func GetCurrentWeather(c wpb.WeatherServiceClient, ctx context.Context, country, city string) {
	req := &wpb.Place{Country: country, City: city}
	resp, err := c.GetCurrentWeather(ctx, req)
	if err != nil {
		log.Fatalf("could not get: %v", err)
	}
	fmt.Println(resp)
}

func GetWeatherForecast(c wpb.WeatherServiceClient, ctx context.Context, country, city, date string) {
	req := &wpb.Forecast{
		Place: &wpb.Place{Country: country, City: city},
		Date:  date,
	}
	resp, err := c.GetWeatherForecast(ctx, req)
	if err != nil {
		log.Fatalf("could not get: %v", err)
	}
	fmt.Println(resp)
}

func ReportWeatherCondition(c wpb.WeatherServiceClient, ctx context.Context, country, city string) {
	req := &wpb.Place{Country: country, City: city}
	resp, err := c.ReportWeatherCondition(ctx, req)
	if err != nil {
		log.Fatalf("could not report: %v", err)
	}
	fmt.Println(resp)
}

func GetBusSchedule(c tpb.TransportServiceClient, ctx context.Context, number string) {
	req := &tpb.Number{Number: number}
	resp, err := c.GetBusSchedule(ctx, req)
	if err != nil {
		log.Fatalf("could not get: %v", err)
	}
	fmt.Println(resp)
}

func TrackBusLocation(c tpb.TransportServiceClient, ctx context.Context, number string) {
	req := &tpb.Number{Number: number}
	resp, err := c.TrackBusLocation(ctx, req)
	if err != nil {
		log.Fatalf("could not track: %v", err)
	}
	fmt.Println(resp)
}

func ReportTrafficJam(c tpb.TransportServiceClient, ctx context.Context, name string, transports int32) {
	req := &tpb.Route{Name: name, Transports: transports}
	resp, err := c.ReportTrafficJam(ctx, req)
	if err != nil {
		log.Fatalf("could not report: %v", err)
	}
	fmt.Println(resp)
}
