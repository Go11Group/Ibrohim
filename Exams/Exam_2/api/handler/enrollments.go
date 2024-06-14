package handler

import (
	"language_learning_app/model"
	"net/http"
	"strconv"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// CRUD operations

// get an enrollment
func (h *Handler) EnrollmentGet(c *gin.Context) {
	// getting id from url
	id := c.Param("enrollment_id")

	// implementation
	enrollment, err := h.Enrollment.Read(id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	// response
	c.JSON(http.StatusOK, gin.H{"enrollment": *enrollment})
}

// post an enrollment
func (h *Handler) EnrollmentPost(c *gin.Context) {
	// getting data from form body
	var enrollment model.Enrollment
	if err := c.ShouldBind(&enrollment); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// assigning random uuid if enrollment_id is not provided
	if enrollment.EnrollmentID == "" {
		randomUUID,_ := uuid.NewRandom()
		enrollment.EnrollmentID = randomUUID.String()
	}

	// implementation
	err := h.Enrollment.EnrollUserOnCourse(enrollment)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// response
	c.JSON(http.StatusOK, "user enrolled in new course")
}

// put an enrollment
func (h *Handler) EnrollmentPut(c *gin.Context) {
	// getting id from url
	id := c.Param("enrollment_id")
	
	// getting data from form body
	var newData model.Enrollment
	if err := c.ShouldBind(&newData); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	// implementation
	err := h.Enrollment.Update(id, newData)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// response
	c.JSON(http.StatusOK, "enrollment with id " + id + " updated")
}

// delete an enrollment
func (h *Handler) EnrollmentDelete(c *gin.Context) {
	// getting id from url
	id := c.Param("enrollment_id")

	// implementation
	err := h.Enrollment.Delete(id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	// response
	c.JSON(http.StatusOK, "enrollment with id " + id + " deleted")
}


// Additional methods

// get all enrollments with filtering
func (h *Handler) EnrollmentsGet(c *gin.Context) {
	// constructing filter based on query parameters
	var filter model.EnrollmentFilter
	filter.EnrollmentDate = c.Query("enrollment_date")
	filter.Limit, _ = strconv.Atoi(c.Query("limit"))
	filter.Offset, _ = strconv.Atoi(c.Query("offset"))
	
	// implementation
	enrollments, err := h.Enrollment.GetAllEnrollments(filter)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	
	// response
	c.JSON(http.StatusOK, gin.H{"all enrollments": enrollments})
}