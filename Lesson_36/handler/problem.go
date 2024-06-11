package handler

import (
	"fmt"
	"gin_pg/model"
	"net/http"
	"github.com/gin-gonic/gin"
)

func (h * Handler) problemGet(c *gin.Context) {
	id, err := GetID(c, "problem")
	if id == 0 || err != nil {
		return
	}
	problems, err := h.Problem.GetProblem(model.Problem{ID: id})
	if err != nil {
		c.AbortWithError(http.StatusNotFound, fmt.Errorf("problem not found"))
		fmt.Println("Error: ", err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"Problem": problems})
}

func (h *Handler) problemPost(c *gin.Context) {
	p := model.Problem{}
	if err := ReadProblemBody(c, &p); err != nil {
		return
	}
	err := h.Problem.CreateProblem(p)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, fmt.Errorf("failed to create problem"))
		fmt.Println("Error: ", err)
		return
	}
	c.JSON(http.StatusOK, "New problem inserted to database")
}

func (h *Handler) problemPut(c *gin.Context) {
	id, err := GetID(c, "problem")
	if id == 0 || err != nil {
		return
	}
	p := model.Problem{ID: id}
	if err := ReadProblemBody(c, &p); err != nil {
		return
	}
	err = h.Problem.UpdateProblem(p)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, fmt.Errorf("failed to update problem"))
		fmt.Println("Error: ", err)
		return
	}
	c.JSON(http.StatusOK, fmt.Sprintf("Problem with ID %d updated", p.ID))
}

func (h *Handler) problemDelete(c *gin.Context) {
	id, err := GetID(c, "problem")
	if id == 0 || err != nil {
		return
	}
	err = h.Problem.DeleteProblem(id)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, fmt.Errorf("failed to delete problem"))
		fmt.Println("Error: ", err)
		return
	}
	c.JSON(http.StatusOK, fmt.Sprintf("Problem with ID %d deleted", id))
}