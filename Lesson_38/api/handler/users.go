package handler

import (
	"language_learning_app/model"
	"net/http"
	"strconv"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (h *Handler) UserGet(c *gin.Context) {
	id := c.Param("id")
	user, err := h.User.Read(id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"User": *user})
}

func (h *Handler) UserPost(c *gin.Context) {
	var user model.User
	if err := c.ShouldBind(&user); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	randomUUID,_ := uuid.NewRandom()
	user.UserID = randomUUID.String()
	err := h.User.Create(user)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, "New user inserted to database")
}

func (h *Handler) UserPut(c *gin.Context) {
	id := c.Param("id")
	var newData model.User
	if err := c.ShouldBind(&newData); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := h.User.Update(id, newData)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, "User with ID " + id + " updated")
}

func (h *Handler) UserDelete(c *gin.Context) {
	id := c.Param("id")
	err := h.User.Delete(id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, "User with ID " + id + " deleted")
}

func (h *Handler) AllUsersGet(c *gin.Context) {
	var filter model.UserFilter
	filter.Name = c.Query("name")
	filter.AgeFrom, _ = strconv.Atoi(c.Query("age_from"))
	filter.AgeTo, _ = strconv.Atoi(c.Query("age_to"))
	filter.Limit, _ = strconv.Atoi(c.Query("limit"))
	filter.Offset, _ = strconv.Atoi(c.Query("offset"))
	users, err := h.User.GetAllUsers(filter)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"All users": users})
}