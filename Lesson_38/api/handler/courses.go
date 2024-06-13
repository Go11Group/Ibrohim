package handler

import (
	"language_learning_app/model"
	"net/http"
	"strconv"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (h *Handler) CourseGet(c *gin.Context) {
	id := c.Param("id")
	course, err := h.Course.Read(id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"Course": *course})
}

func (h *Handler) CoursePost(c *gin.Context) {
	var course model.Course
	if err := c.ShouldBind(&course); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	randomUUID,_ := uuid.NewRandom()
	course.CourseID = randomUUID.String()
	err := h.Course.Create(course)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, "New course inserted to database")
}

func (h *Handler) CoursePut(c *gin.Context) {
	id := c.Param("id")
	var newData model.Course
	if err := c.ShouldBind(&newData); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := h.Course.Update(id, newData)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, "Course with ID " + id + " updated")
}

func (h *Handler) CourseDelete(c *gin.Context) {
	id := c.Param("id")
	err := h.Course.Delete(id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, "Course with ID " + id + " deleted")
}

func (h *Handler) AllCoursesGet(c *gin.Context) {
	var filter model.CourseFilter
	filter.Title = c.Query("title")
	filter.Description = c.Query("description")
	filter.Limit, _ = strconv.Atoi(c.Query("limit"))
	filter.Offset, _ = strconv.Atoi(c.Query("offset"))
	courses, err := h.Course.GetAllCourses(filter)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"All courses": courses})
}