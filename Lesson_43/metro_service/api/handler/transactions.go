package handler

import (
	"metro-service/model"
	"net/http"
	"strconv"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (h *Handler) GetTransaction(c *gin.Context) {
	id := c.Param("id")
	transaction, err := h.Transaction.Read(id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"Transaction": *transaction})
}

func (h *Handler) PostTransaction(c *gin.Context) {
	var transaction model.Transaction
	if err := c.ShouldBind(&transaction); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if transaction.ID == "" {
		randomUUID, _ := uuid.NewRandom()
		transaction.ID = randomUUID.String()
	}

	err := h.Transaction.Create(&transaction)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, "New transaction inserted to database")
}

func (h *Handler) PutTransaction(c *gin.Context) {
	id := c.Param("id")
	newData := model.Transaction{ID: id}
	if err := c.ShouldBind(&newData); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	err := h.Transaction.Update(&newData)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{id: "Transaction updated successfully"})
}

func (h *Handler) DeleteTransaction(c *gin.Context) {
	id := c.Param("id")
	err := h.Transaction.Delete(id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{id: "Transaction deleted successfully"})
}

func (h *Handler) GetAllTransactions(c *gin.Context) {
	var filter model.TransactionFilter
	filter.CardID = c.Query("card_id")
	filter.Amount, _ = strconv.Atoi(c.Query("amount"))
	filter.TerminalID = c.Query("terminal_id")
	filter.Type = c.Query("type")
	filter.Limit, _ = strconv.Atoi(c.Query("limit"))
	filter.Offset, _ = strconv.Atoi(c.Query("offset"))

	transactions, err := h.Transaction.FetchTransactions(filter)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"All transactions": transactions})
}
