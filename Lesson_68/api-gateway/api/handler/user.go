package handler

import (
	pb "api-gateway/genproto/user"
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

// GetProfile godoc
// @Summary Gets user
// @Description Retrieves user profile by ID
// @Tags user
// @Security ApiKeyAuth
// @Param id path string true "User ID"
// @Success 200 {object} user.Profile
// @Failure 400 {object} string "Invalid user id"
// @Failure 500 {object} string "Server error while processing request"
// @Router /user/{id} [get]
func (h *Handler) GetProfile(c *gin.Context) {
	h.Log.Info("GetProfile handler is starting")

	id, ok := c.Get("user_id")
	if !ok {
		er := errors.New("user id not provided").Error()
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": er})
		h.Log.Error(er)
		return
	}

	ctxv := context.WithValue(c, h.UserIDKey, id)
	ctx, cancel := context.WithTimeout(ctxv, h.ContextTimeout)
	defer cancel()

	pr, err := h.User.GetProfile(ctx, &pb.Void{})
	if err != nil {
		er := errors.Wrap(err, "error getting user").Error()
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": er})
		h.Log.Error(er)
		return
	}

	h.Log.Info("GetProfile handler is completed")
	c.JSON(http.StatusOK, gin.H{"user": pr})
}

// UpdateProfile godoc
// @Summary Updates user
// @Description Updates user info by ID
// @Tags user
// @Security ApiKeyAuth
// @Param id path string true "User ID"
// @Param user body user.NewData true "New user data"
// @Success 200 {object} user.UpdateResp
// @Failure 400 {object} string "Invalid user id or data"
// @Failure 500 {object} string "Server error while processing request"
// @Router /user/{id} [put]
func (h *Handler) UpdateProfile(c *gin.Context) {
	h.Log.Info("UpdateProfile handler is starting")

	id, ok := c.Get("user_id")
	if !ok {
		er := errors.New("user id not provided").Error()
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": er})
		h.Log.Error(er)
		return
	}

	var data pb.NewData
	err := c.ShouldBind(&data)
	if err != nil {
		er := errors.Wrap(err, "invalid user data").Error()
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": er})
		h.Log.Error(er)
		return
	}

	ctxv := context.WithValue(c, h.UserIDKey, id)
	ctx, cancel := context.WithTimeout(ctxv, h.ContextTimeout)
	defer cancel()

	resp, err := h.User.UpdateProfile(ctx, &data)
	if err != nil {
		er := errors.Wrap(err, "error updating user").Error()
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": er})
		h.Log.Error(er)
		return
	}
	h.Log.Info("UpdateProfile handler is completed")
	c.JSON(http.StatusOK, gin.H{"user": resp})
}

// DeleteProfile godoc
// @Summary Deletes user
// @Description Deletes user info by ID
// @Tags user
// @Security ApiKeyAuth
// @Param id path string true "User ID"
// @Success 200 {string} string "User deleted successfully"
// @Failure 400 {object} string "Invalid user id"
// @Failure 500 {object} string "Server error while processing request"
// @Router /user/{id} [delete]
func (h *Handler) DeleteProfile(c *gin.Context) {
	h.Log.Info("DeleteProfile handler is starting")

	id, ok := c.Get("user_id")
	if !ok {
		er := errors.New("user id not provided").Error()
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": er})
		h.Log.Error(er)
		return
	}

	ctxv := context.WithValue(c, h.UserIDKey, id)
	ctx, cancel := context.WithTimeout(ctxv, h.ContextTimeout)
	defer cancel()

	_, err := h.User.DeleteProfile(ctx, &pb.Void{})
	if err != nil {
		er := errors.Wrap(err, "error deleting user").Error()
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": er})
		h.Log.Error(er)
		return
	}
	h.Log.Info("DeleteProfile handler is completed")

	c.JSON(http.StatusOK, "user deleted successfully")
}
