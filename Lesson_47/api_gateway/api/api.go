package api

import (
	"api-gateway/api/handler"
	pb "api-gateway/genproto"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

func NewRouter(conn *grpc.ClientConn) *gin.Engine {
	router := gin.Default()
	
	weather := pb.NewWeatherServiceClient(conn)
	transport := pb.NewTransportServiceClient(conn)

	h := handler.NewHandler(weather, transport)

	w := router.Group("weather")
	{
		w.GET("/current", h.GetCurrentWeather)
		w.GET("/forecast", h.GetWeatherForecast)
		w.GET("/condition", h.ReportWeatherCondition)
	}
	
	t := router.Group("transport")
	{
		t.GET("/schedule", h.GetBusSchedule)
		t.GET("/location", h.TrackBusLocation)
		t.GET("/traffic-jam", h.ReportTrafficJam)
	}

	return router
}
