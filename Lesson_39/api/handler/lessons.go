package handler

import (
	"language_learning_app/model"
	"net/http"
	"strconv"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// CRUD operations

// get a lesson
func (h *Handler) LessonGet(c *gin.Context) {
	// getting id from url
	id := c.Param("lesson_id")

	// implementation
	lesson, err := h.Lesson.Read(id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	// response
	c.JSON(http.StatusOK, gin.H{"lesson": *lesson})
}

// post a lesson
func (h *Handler) LessonPost(c *gin.Context) {
	// getting data from form body
	var lesson model.Lesson
	if err := c.ShouldBind(&lesson); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// assigning random uuid if lesson_id is not provided
	if lesson.LessonID == "" {
		randomUUID,_ := uuid.NewRandom()
		lesson.LessonID = randomUUID.String()
	}

	// implementation
	err := h.Lesson.Create(lesson)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// response
	c.JSON(http.StatusOK, "new lesson added to course with id "+lesson.CourseID)
}

// put a lesson
func (h *Handler) LessonPut(c *gin.Context) {
	// getting id from url
	id := c.Param("lesson_id")

	// getting data from form body
	var newData model.Lesson
	if err := c.ShouldBind(&newData); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// implementation
	err := h.Lesson.Update(id, newData)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// response
	c.JSON(http.StatusOK, "lesson with id " + id + " updated")
}

// delete a lesson
func (h *Handler) LessonDelete(c *gin.Context) {
	// getting id from url
	id := c.Param("lesson_id")

	// implementation
	err := h.Lesson.Delete(id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	// response
	c.JSON(http.StatusOK, "lesson with id " + id + " deleted")
}


// Additional methods

// get all lessons with filtering
func (h *Handler) LessonsGet(c *gin.Context) {
	// constructing filter based on query parameters
	var filter model.LessonFilter
	filter.Title = c.Query("title")
	filter.Content = c.Query("content")
	filter.Limit, _ = strconv.Atoi(c.Query("limit"))
	filter.Offset, _ = strconv.Atoi(c.Query("offset"))
	
	// implementation
	lessons, err := h.Lesson.GetAllLessons(filter)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// response
	c.JSON(http.StatusOK, gin.H{"all lessons": lessons})
}
