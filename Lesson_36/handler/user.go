package handler

import (
	"fmt"
	"gin_pg/model"
	"net/http"
	"github.com/gin-gonic/gin"
)

func (h * Handler) userGet(c *gin.Context) {
	id, err := GetID(c, "user")
	if id == 0 || err != nil {
		return
	}
	users, err := h.User.GetUser(model.User{ID: id})
	if err != nil {
		c.AbortWithError(http.StatusNotFound, fmt.Errorf("user not found"))
		fmt.Println("Error: ", err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"User": users})
}

func (h *Handler) userPost(c *gin.Context) {
	u := model.User{}
	if err := ReadUserBody(c, &u); err != nil {
		return
	}
	err := h.User.CreateUser(u)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, fmt.Errorf("failed to create user"))
		fmt.Println("Error: ", err)
		return
	}
	c.JSON(http.StatusOK, "New user inserted to database")
}

func (h *Handler) userPut(c *gin.Context) {
	id, err := GetID(c, "user")
	if id == 0 || err != nil {
		return
	}
	u := model.User{ID: id}
	if err := ReadUserBody(c, &u); err != nil {
		return
	}
	err = h.User.UpdateUser(u)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, fmt.Errorf("failed to update user"))
		fmt.Println("Error: ", err)
		return
	}
	c.JSON(http.StatusOK, fmt.Sprintf("User with ID %d updated", u.ID))
}

func (h *Handler) userDelete(c *gin.Context) {
	id, err := GetID(c, "user")
	if id == 0 || err != nil {
		return
	}
	err = h.User.DeleteUser(id)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, fmt.Errorf("failed to delete user"))
		fmt.Println("Error: ", err)
		return
	}
	c.JSON(http.StatusOK, fmt.Sprintf("User with ID %d deleted", id))
}