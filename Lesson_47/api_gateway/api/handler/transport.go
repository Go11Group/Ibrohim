package handler

import (
	pb "api-gateway/genproto"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

func (h *Handler) GetBusSchedule(c *gin.Context) {
	req := &pb.Number{Number: c.Query("bus_number")}

	resp, err := h.Transport.GetBusSchedule(c, req)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}

func (h *Handler) TrackBusLocation(c *gin.Context) {
	req := &pb.Number{Number: c.Query("bus_number")}

	resp, err := h.Transport.TrackBusLocation(c, req)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}

func (h *Handler) ReportTrafficJam(c *gin.Context) {
	t, err := strconv.Atoi(c.Query("transports"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": errors.Wrap(err, "invalid number of transports")})
		return
	}

	req := &pb.Route{Name: c.Query("name"), Transports: int32(t)}

	resp, err := h.Transport.ReportTrafficJam(c, req)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}
