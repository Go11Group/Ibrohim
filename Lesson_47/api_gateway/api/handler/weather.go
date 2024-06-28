package handler

import (
	pb "api-gateway/genproto"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetCurrentWeather(c *gin.Context) {
	req := &pb.Place{Country: c.Query("country"), City: c.Query("city")}

	resp, err := h.Weather.GetCurrentWeather(c, req)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}

func (h *Handler) GetWeatherForecast(c *gin.Context) {
	req := &pb.Forecast{
		Place: &pb.Place{Country: c.Query("country"), City: c.Query("city")},
		Date:  c.Query("date"),
	}

	resp, err := h.Weather.GetWeatherForecast(c, req)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}

func (h *Handler) ReportWeatherCondition(c *gin.Context) {
	req := &pb.Place{Country: c.Query("country"), City: c.Query("city")}

	resp, err := h.Weather.ReportWeatherCondition(c, req)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}
