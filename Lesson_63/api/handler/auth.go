package handler

import (
	"auth-service/api/models"
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

// Login godoc
// @Summary Logs user in
// @Description Logs user in
// @Tags auth
// @Param data body models.LoginRequest true "User credentials"
// @Success 200 {object} models.Tokens
// @Failure 400 {object} string "Invalid data"
// @Failure 500 {object} string "Server error while processing request"
// @Router /login [post]
func (h *Handler) Login(c *gin.Context) {
	h.Log.Info("Login function is starting")

	var data models.LoginRequest
	if err := c.ShouldBind(&data); err != nil {
		er := errors.Wrap(err, "invalid data").Error()
		c.AbortWithStatusJSON(http.StatusBadRequest,
			gin.H{"error": er},
		)
		h.Log.Error(er)
		return
	}

	ctx, cancel := context.WithTimeout(c, time.Second*5)
	defer cancel()

	tokens, err := h.Auth.Login(ctx, &data)
	if err != nil {
		er := errors.Wrap(err, "error logging in user").Error()
		c.AbortWithStatusJSON(http.StatusBadRequest,
			gin.H{"error": er},
		)
		h.Log.Error(er)
		return
	}

	h.Log.Info("Login has successfully finished")
	c.JSON(http.StatusOK, gin.H{"Tokens": tokens})
}

// Refresh godoc
// @Summary Refreshes refresh token
// @Description Refreshes refresh token
// @Tags auth
// @Param data body models.RefreshToken true "Refresh token"
// @Success 200 {object} models.Tokens
// @Failure 400 {object} string "Invalid data"
// @Failure 500 {object} string "Server error while processing request"
// @Router /refresh-token [post]
func (h Handler) Refresh(c *gin.Context) {
	h.Log.Info("Refresh function is starting")

	var data models.RefreshToken
	if err := c.ShouldBind(&data); err != nil {
		er := errors.Wrap(err, "invalid data").Error()
		c.AbortWithStatusJSON(http.StatusBadRequest,
			gin.H{"error": er},
		)
		h.Log.Error(er)
		return
	}

	ctx, cancel := context.WithTimeout(c, time.Second*5)
	defer cancel()

	tokens, err := h.Auth.RefreshToken(ctx, &data)
	if err != nil {
		er := errors.Wrap(err, "error refreshing token").Error()
		c.AbortWithStatusJSON(http.StatusBadRequest,
			gin.H{"error": er},
		)
		h.Log.Error(er)
		return
	}

	h.Log.Info("Refresh has successfully finished")
	c.JSON(http.StatusOK, gin.H{"Tokens": tokens})
}
