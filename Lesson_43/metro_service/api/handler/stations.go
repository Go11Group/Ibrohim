package handler

import (
	"metro-service/model"
	"net/http"
	"strconv"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (h *Handler) GetStation(c *gin.Context) {
	id := c.Param("id")
	station, err := h.Station.Read(id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"Station": *station})
}

func (h *Handler) PostStation(c *gin.Context) {
	var station model.Station
	if err := c.ShouldBind(&station); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if station.ID == "" {
		randomUUID, _ := uuid.NewRandom()
		station.ID = randomUUID.String()
	}

	err := h.Station.Create(&station)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, "New user inserted to database")
}

func (h *Handler) PutStation(c *gin.Context) {
	id := c.Param("id")
	newData := model.Station{ID: id}
	if err := c.ShouldBind(&newData); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	err := h.Station.Update(&newData)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{id: "Station updated successfully"})
}

func (h *Handler) DeleteStation(c *gin.Context) {
	id := c.Param("id")
	err := h.Station.Delete(id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{id: "Station deleted successfully"})
}

func (h *Handler) GetAllStations(c *gin.Context) {
	var filter model.StationFilter
	filter.Name = c.Query("name")
	filter.Limit, _ = strconv.Atoi(c.Query("limit"))
	filter.Offset, _ = strconv.Atoi(c.Query("offset"))

	stations, err := h.Station.FetchStations(filter)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"All stations": stations})
}
