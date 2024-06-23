package handler

import (
	"metro-service/model"
	"net/http"
	"strconv"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (h *Handler) GetUser(c *gin.Context) {
	id := c.Param("id")
	card, err := h.Card.Read(id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"card": *card})
}

func (h *Handler) PostUser(c *gin.Context) {
	var card model.Card
	if err := c.ShouldBind(&card); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if card.ID == "" {
		randomUUID, _ := uuid.NewRandom()
		card.ID = randomUUID.String()
	}

	err := h.Card.Create(&card)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, "new card inserted to database")
}

func (h *Handler) PutUser(c *gin.Context) {
	id := c.Param("id")
	newData := model.Card{ID: id}
	if err := c.ShouldBind(&newData); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.Card.Update(&newData)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{id: "Card updated successfully"})
}

func (h *Handler) DeleteUser(c *gin.Context) {
	id := c.Param("id")
	err := h.Card.Delete(id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{id: "Card deleted successfully"})
}

func (h *Handler) GetAllUsers(c *gin.Context) {
	var filter model.CardFilter
	filter.UserID = c.Query("user_id")
	filter.Limit, _ = strconv.Atoi(c.Query("limit"))
	filter.Offset, _ = strconv.Atoi(c.Query("offset"))

	cards, err := h.Card.FetchCards(filter)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"all cards": cards})
}
