package handler

import (
	"language_learning_app/model"
	"net/http"
	"strconv"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// CRUD operations

// get a user
func (h *Handler) UserGet(c *gin.Context) {
	// getting id from url
	id := c.Param("user_id")

	// implementation
	user, err := h.User.Read(id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	
	// response
	c.JSON(http.StatusOK, gin.H{"user": *user})
}

// post a user
func (h *Handler) UserPost(c *gin.Context) {
	// getting data from form body
	var user model.User
	if err := c.ShouldBind(&user); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// assigning random uuid if user_id is not provided
	if user.UserID == "" {
		randomUUID,_ := uuid.NewRandom()
		user.UserID = randomUUID.String()
	}
	
	// implementation
	err := h.User.Create(user)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	
	// response
	c.JSON(http.StatusOK, "new user inserted to database")
}

// put a user
func (h *Handler) UserPut(c *gin.Context) {
	// getting id from url
	id := c.Param("user_id")

	// getting data from form body
	var newData model.User
	if err := c.ShouldBind(&newData); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// implementation
	err := h.User.Update(id, newData)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// response
	c.JSON(http.StatusOK, "user with id " + id + " updated")
}

// delete a user
func (h *Handler) UserDelete(c *gin.Context) {
	// getting id from url
	id := c.Param("user_id")

	// implementation
	err := h.User.Delete(id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	// response
	c.JSON(http.StatusOK, "user with id " + id + " deleted")
}


// Additional methods

// get all users with filtering
func (h *Handler) UsersGet(c *gin.Context) {
	// constructing filter based on query parameters
	var filter model.UserFilter
	filter.Name = c.Query("name")
	filter.AgeFrom, _ = strconv.Atoi(c.Query("age_from"))
	filter.AgeTo, _ = strconv.Atoi(c.Query("age_to"))
	filter.Limit, _ = strconv.Atoi(c.Query("limit"))
	filter.Offset, _ = strconv.Atoi(c.Query("offset"))

	// implementation
	users, err := h.User.GetAllUsers(filter)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// response
	c.JSON(http.StatusOK, gin.H{"all users": users})
}

// get courses a certain user is enrolled on
func (h *Handler) UserCoursesGet(c *gin.Context) {
	// getting id from url
	id := c.Param("user_id")

	// implementation
	courses, err := h.User.GetCourses(id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// response
	c.JSON(http.StatusOK, gin.H{"user_id": id, "courses": courses})
}

// search a user basen on name, email and age
func (h *Handler) UserSearch(c *gin.Context) {
	// getting name, email and age from query parameters
	name  := c.Query("name")
	email := c.Query("email")
	ageFrom,_  := strconv.Atoi(c.Query("age_from"))
	ageTo, _ := strconv.Atoi(c.Query("age_to"))

	// implementation
	users, err := h.User.Search(name, email, ageFrom, ageTo)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// response
	c.JSON(http.StatusOK, gin.H{"results": users})
}