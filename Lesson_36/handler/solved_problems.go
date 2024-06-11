package handler

import (
	"fmt"
	"net/http"
	"strconv"
	"time"
	"github.com/gin-gonic/gin"
)

func QueryParameters(c *gin.Context) (int, int, *time.Time, error) {
	userID, err := strconv.Atoi(c.Query("user-id"))
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, fmt.Errorf("invalid user ID"))
		fmt.Println("Error: ", err)
		return 0, 0, nil, err
	}
	problemID, err := strconv.Atoi(c.Query("problem-id"))
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, fmt.Errorf("invalid problem ID"))
		fmt.Println("Error: ", err)
		return 0, 0, nil, err
	}
	solvedAt, err := time.Parse("02/01/2006 15:04:05", c.Query("solved-at"))
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, fmt.Errorf("invalid time of solution"))
		fmt.Println("Error: ", err)
		return 0, 0, nil, err
	}
	return userID, problemID, &solvedAt, nil
}

func (h * Handler) userProblemGet(c *gin.Context) {
	userID, err := GetID(c, "user")
	if userID == 0 || err != nil {
		return
	}
	problems, err := h.UserProblem.GetUserProblems(userID)
	if err != nil {
		c.AbortWithError(http.StatusNotFound, fmt.Errorf("problems not found for UserID %d", userID))
		fmt.Println("Error: ", err)
		return
	}
	c.JSON(http.StatusOK, gin.H{fmt.Sprintf("UserID %d solved problems", userID): problems})
}

func (h *Handler) userProblemPost(c *gin.Context) {
	userID, problemID, solvedAt, err := QueryParameters(c)
	if err != nil || userID == 0 || problemID == 0 || solvedAt == nil {
		return
	}
	err = h.UserProblem.AddProblemToUser(userID, problemID, *solvedAt)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, fmt.Errorf("failed to add problem to user"))
		fmt.Println("Error: ", err)
		return
	}
	c.JSON(http.StatusOK, fmt.Sprintf("UserID %d solved the %dth problem", userID, problemID))
}

func (h *Handler) userProblemPut(c *gin.Context) {
	userID, problemID, solvedAt, err := QueryParameters(c)
	if err != nil || userID == 0 || problemID == 0 || solvedAt == nil {
		return
	}
	err = h.UserProblem.UpdateTimeOfSolution(userID, problemID, *solvedAt)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, fmt.Errorf("failed to update time of solution"))
		fmt.Println("Error: ", err)
		return
	}
	c.JSON(http.StatusOK, fmt.Sprintf("The %dth problem time of solution for UserID %d updated", problemID, userID))
}

func (h *Handler) userProblemDelete(c *gin.Context) {
	userID, problemID, _, err := QueryParameters(c)
	if err != nil || userID == 0 || problemID == 0 {
		return
	}
	err = h.UserProblem.RemoveProblemFromUser(userID, problemID)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, fmt.Errorf("failed to delete problem from user"))
		fmt.Println("Error: ", err)
		return
	}
	c.JSON(http.StatusOK, fmt.Sprintf("The %dth problem removed from UserID %d", problemID, userID))
}