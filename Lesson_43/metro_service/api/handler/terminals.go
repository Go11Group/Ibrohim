package handler

import (
	"metro-service/model"
	"net/http"
	"strconv"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (h *Handler) GetTerminal(c *gin.Context) {
	id := c.Param("id")
	terminal, err := h.Terminal.Read(id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"Terminal": *terminal})
}

func (h *Handler) PostTerminal(c *gin.Context) {
	var terminal model.Terminal
	if err := c.ShouldBind(&terminal); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if terminal.ID == "" {
		randomUUID, _ := uuid.NewRandom()
		terminal.ID = randomUUID.String()
	}

	err := h.Terminal.Create(&terminal)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, "New terminal inserted to database")
}

func (h *Handler) PutTerminal(c *gin.Context) {
	id := c.Param("id")
	newData := model.Terminal{ID: id}
	if err := c.ShouldBind(&newData); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	err := h.Terminal.Update(&newData)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{id: "Terminal updated successfully"})
}

func (h *Handler) DeleteTerminal(c *gin.Context) {
	id := c.Param("id")
	err := h.Terminal.Delete(id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{id: "Terminal deleted successfully"})
}

func (h *Handler) GetAllTerminals(c *gin.Context) {
	var filter model.TerminalFilter
	filter.StationID = c.Query("station_id")
	filter.Limit, _ = strconv.Atoi(c.Query("limit"))
	filter.Offset, _ = strconv.Atoi(c.Query("offset"))

	terminals, err := h.Terminal.FetchTerminals(filter)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"All terminals": terminals})
}
