package handler

import (
	"net/http"
	"strconv"
	"user-service/model"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (h *Handler) Get(c *gin.Context) {
	id := c.Param("id")
	user, err := h.User.Read(id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"User": *user})
}

func (h *Handler) Post(c *gin.Context) {
	var user model.User
	if err := c.ShouldBind(&user); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if user.ID == "" {
		randomUUID, _ := uuid.NewRandom()
		user.ID = randomUUID.String()
	}

	err := h.User.Create(&user)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, "New user inserted to database")
}

func (h *Handler) Put(c *gin.Context) {
	id := c.Param("id")
	newData := model.User{ID: id}
	if err := c.ShouldBind(&newData); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	err := h.User.Update(&newData)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{id: "User updated successfully"})
}

func (h *Handler) Delete(c *gin.Context) {
	id := c.Param("id")
	err := h.User.Delete(id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{id: "User deleted successfully"})
}

func (h *Handler) GetAll(c *gin.Context) {
	var filter model.UserFilter
	filter.Name = c.Query("name")
	filter.AgeFrom, _ = strconv.Atoi(c.Query("age_from"))
	filter.AgeTo, _ = strconv.Atoi(c.Query("age_to"))
	filter.Limit, _ = strconv.Atoi(c.Query("limit"))
	filter.Offset, _ = strconv.Atoi(c.Query("offset"))

	users, err := h.User.FetchUsers(filter)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"All users": users})
}
