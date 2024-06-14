package handler

import (
	"language_learning_app/model"
	"net/http"
	"strconv"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// CRUD operations

// get a course
func (h *Handler) CourseGet(c *gin.Context) {
	// getting id from url
	id := c.Param("course_id")

	// implementation
	course, err := h.Course.Read(id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	// response
	c.JSON(http.StatusOK, gin.H{"course": *course})
}

// post a course
func (h *Handler) CoursePost(c *gin.Context) {
	// getting data from form body
	var course model.Course
	if err := c.ShouldBind(&course); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// assigning random uuid if course_id is not provided
	if course.CourseID == "" {
		randomUUID,_ := uuid.NewRandom()
		course.CourseID = randomUUID.String()
	}
	
	// implementation
	err := h.Course.Create(course)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// response
	c.JSON(http.StatusOK, "new course inserted to database")
}

// put a course
func (h *Handler) CoursePut(c *gin.Context) {
	// getting id from url
	id := c.Param("course_id")

	// getting data from form body
	var newData model.Course
	if err := c.ShouldBind(&newData); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// implementation
	err := h.Course.Update(id, newData)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// response
	c.JSON(http.StatusOK, "course with id " + id + " updated")
}

// delete a course
func (h *Handler) CourseDelete(c *gin.Context) {
	// getting id from url
	id := c.Param("course_id")

	// implementation
	err := h.Course.Delete(id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	// response
	c.JSON(http.StatusOK, "course with id " + id + " deleted")
}


// Additional methods

// get all courses with filtering
func (h *Handler) CoursesGet(c *gin.Context) {
	// constructing filter based on query parameters
	var filter model.CourseFilter
	filter.Title = c.Query("title")
	filter.Description = c.Query("description")
	filter.Limit, _ = strconv.Atoi(c.Query("limit"))
	filter.Offset, _ = strconv.Atoi(c.Query("offset"))

	// implementation
	courses, err := h.Course.GetAllCourses(filter)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// response
	c.JSON(http.StatusOK, gin.H{"all courses": courses})
}

// get lessons of a certain course
func (h *Handler) CourseLessons(c *gin.Context) {
	// getting id from url
	id := c.Param("course_id")

	// implementation
	lessons, err := h.Course.GetLessons(id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// response
	c.JSON(http.StatusOK, gin.H{"course_id": id, "lessons": lessons})
}

// get users enrolled on a certain course
func (h *Handler) CourseUsers(c *gin.Context) {
	// getting id from url
	id := c.Param("course_id")

	// implementation
	users, err := h.Course.GetUsers(id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// response
	c.JSON(http.StatusOK, gin.H{"course_id": id, "enrolled_users": users})
}

// get courses ordered by popularity between particular dates
func (h *Handler) CoursePopular(c *gin.Context) {
	// getting dates from query parameters
	start_date := c.Query("start_date")
	end_date   := c.Query("end_date")

	// implementation
	courses, err := h.Course.Popular(start_date, end_date)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// response
	c.JSON(http.StatusOK, gin.H{"time_period": []string{start_date, end_date}, "popular_courses": courses})
}