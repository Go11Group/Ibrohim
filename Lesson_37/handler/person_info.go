package handler

import (
	"les37/model"
	"net/http"
	"github.com/gin-gonic/gin"
)

func (h *Handler) getAll(c *gin.Context) {
	var f model.Filter
	if err := c.ShouldBindQuery(&f); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error in ShouldBindQuery": err.Error()})
        return
	}
	people, err := h.Person.GetAll(f)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error in GetAll method": err.Error()})
        return
	}
	c.JSON(http.StatusOK, gin.H{"People":people})
}