package main

import (
	"context"
	"fmt"
	"log"
	"time"
	tpb "weather-transport/genproto/transport"
	wpb "weather-transport/genproto/weather"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.NewClient(":50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to connect to the address: %v", err)
	}
	defer conn.Close()

	wc := wpb.NewWeatherServiceClient(conn)
	tc := tpb.NewTransportServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Hour)
	defer cancel()

	var n int
	fmt.Println("1 - weather | 2 - transport")
	fmt.Print(">> ")
	fmt.Scan(&n)

	op := 4
	var country, city, date, number, name string
	var transports int32
outerLoop:
	switch n {
	case 1:
		for {
			fmt.Println("1: current weather")
			fmt.Println("2: weather forecast")
			fmt.Println("3: weather condition")
			fmt.Println("0: exit")
			fmt.Print(">> ")
			fmt.Scan(&op)

			switch op {
			case 1:
				fmt.Print("Country: ")
				fmt.Scan(&country)
				fmt.Print("City: ")
				fmt.Scan(&city)
				GetCurrentWeather(wc, ctx, country, city)
			case 2:
				fmt.Print("Country: ")
				fmt.Scan(&country)
				fmt.Print("City: ")
				fmt.Scan(&city)
				fmt.Print("Forecast date: ")
				fmt.Scan(&date)
				GetWeatherForecast(wc, ctx, country, city, date)
			case 3:
				fmt.Print("Country: ")
				fmt.Scan(&country)
				fmt.Print("City: ")
				fmt.Scan(&city)
				ReportWeatherCondition(wc, ctx, country, city)
			case 0:
				break outerLoop
			default:
				fmt.Println("Invalid option")
			}
		}
	case 2:
		for {
			fmt.Println("1: bus schedule")
			fmt.Println("2: bus location")
			fmt.Println("3: traffic jam")
			fmt.Println("0: exit")
			fmt.Print(">> ")
			fmt.Scan(&op)

			switch op {
			case 1:
				fmt.Print("Bus number: ")
				fmt.Scan(&number)
				GetBusSchedule(tc, ctx, number)
			case 2:
				fmt.Print("Bus number: ")
				fmt.Scan(&number)
				TrackBusLocation(tc, ctx, number)
			case 3:
				fmt.Print("Route name: ")
				fmt.Scan(&name)
				fmt.Print("Number of transports: ")
				fmt.Scan(&transports)
				ReportTrafficJam(tc, ctx, name, transports)
			case 0:
				break outerLoop
			default:
				fmt.Println("Invalid option")
			}
		}
	}
}
