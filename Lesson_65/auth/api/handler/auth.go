package handler

import (
	"auth-service/api/tokens"
	pb "auth-service/genproto/admin"
	"auth-service/models"
	"context"
	"net/http"
	"time"

	"github.com/google/uuid"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
)

// Register godoc
// @Summary Registers user
// @Description Registers a new user
// @Tags user
// @Security ApiKeyAuth
// @Param user body admin.NewUser true "User data"
// @Success 200 {object} admin.NewUserResp
// @Failure 400 {object} string "Invalid data"
// @Failure 500 {object} string "Server error while processing request"
// @Router /register [post]
func (h *Handler) Register(c *gin.Context) {
	h.Log.Info("Register function is starting")
	var req *pb.NewUser
	if err := c.ShouldBindJSON(&req); err != nil {
		er := errors.Wrap(err, "invalid data provided").Error()
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": er})
		h.Log.Error(er)
		return
	}
	passByte, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		er := errors.Wrap(err, "invalid password").Error()
		c.AbortWithStatusJSON(http.StatusBadRequest,
			gin.H{"error": er},
		)
		h.Log.Error(er)
		return
	}
	req.Password = string(passByte)

	ctx, cancel := context.WithTimeout(c, time.Second*5)
	defer cancel()

	req.Role = "user"

	resp, err := h.RepoAdmin.Add(ctx, req)
	if err != nil {
		er := errors.Wrap(err, "error registering user").Error()
		c.AbortWithStatusJSON(http.StatusInternalServerError,
			gin.H{"error": er},
		)
		h.Log.Error(er)
		return
	}

	h.Log.Info("Register has successfully finished")
	c.JSON(http.StatusOK, gin.H{"data": resp})
}

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

	var req models.LoginRequest
	if err := c.ShouldBind(&req); err != nil {
		er := errors.Wrap(err, "invalid data").Error()
		c.AbortWithStatusJSON(http.StatusBadRequest,
			gin.H{"error": er},
		)
		h.Log.Error(er)
		return
	}

	ctx, cancel := context.WithTimeout(c, time.Second*5)
	defer cancel()

	id, username, passwordHash, err := h.RepoUser.GetUserByEmail(ctx, req.Email)
	if err != nil {
		er := errors.Wrap(err, "user not found").Error()
		c.AbortWithStatusJSON(http.StatusBadRequest,
			gin.H{"error": er},
		)
		h.Log.Error(er)
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(passwordHash), []byte(req.Password))
	if err != nil {
		er := errors.Wrap(err, "invalid password").Error()
		c.AbortWithStatusJSON(http.StatusBadRequest,
			gin.H{"error": er},
		)
		h.Log.Error(er)
		return
	}

	accessToken, err := tokens.GenerateAccessToken(id, username, req.Email)
	if err != nil {
		er := errors.Wrap(err, "error generating access token").Error()
		c.AbortWithStatusJSON(http.StatusInternalServerError,
			gin.H{"error": er},
		)
		h.Log.Error(er)
		return
	}

	refreshToken, err := tokens.GenerateRefreshToken(id)
	if err != nil {
		er := errors.Wrap(err, "error generating refresh token").Error()
		c.AbortWithStatusJSON(http.StatusInternalServerError,
			gin.H{"error": er},
		)
		h.Log.Error(er)
		return
	}

	exp, err := tokens.GetRefreshTokenExpiry(refreshToken)
	if err != nil {
		er := errors.Wrap(err, "error getting refresh token expiry").Error()
		c.AbortWithStatusJSON(http.StatusInternalServerError,
			gin.H{"error": er},
		)
		h.Log.Error(er)
		return
	}

	err = h.RepoToken.Store(ctx, &models.RefreshTokenDetails{
		UserID: id,
		Token:  refreshToken,
		Expiry: exp,
	})
	if err != nil {
		er := errors.Wrap(err, "error storing refresh token").Error()
		c.AbortWithStatusJSON(http.StatusInternalServerError,
			gin.H{"error": er},
		)
		h.Log.Error(er)
		return
	}

	h.Log.Info("Login has successfully finished")
	c.JSON(http.StatusOK, gin.H{"Tokens": models.Tokens{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}})
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

	var t models.RefreshToken
	if err := c.ShouldBind(&t); err != nil {
		er := errors.Wrap(err, "invalid data").Error()
		c.AbortWithStatusJSON(http.StatusBadRequest,
			gin.H{"error": er},
		)
		h.Log.Error(er)
		return
	}

	ctx, cancel := context.WithTimeout(c, time.Second*5)
	defer cancel()

	valid, err := tokens.ValidateRefreshToken(t.Token)
	if !valid || err != nil {
		er := errors.Wrap(err, "invalid refresh token").Error()
		c.AbortWithStatusJSON(http.StatusBadRequest,
			gin.H{"error": er},
		)
		h.Log.Error(er)
		return
	}

	id, err := tokens.GetUserIdFromRefreshToken(t.Token)
	if err != nil {
		er := errors.Wrap(err, "error getting user id").Error()
		c.AbortWithStatusJSON(http.StatusBadRequest,
			gin.H{"error": er},
		)
		h.Log.Error(er)
		return
	}

	username, email, _, err := h.RepoUser.GetUserByID(ctx, id)
	if err != nil {
		er := errors.Wrap(err, "user not found").Error()
		c.AbortWithStatusJSON(http.StatusBadRequest,
			gin.H{"error": er},
		)
		h.Log.Error(er)
		return
	}

	accessToken, err := tokens.GenerateAccessToken(id, username, email)
	if err != nil {
		er := errors.Wrap(err, "error generating access token").Error()
		c.AbortWithStatusJSON(http.StatusInternalServerError,
			gin.H{"error": er},
		)
		h.Log.Error(er)
		return
	}

	h.Log.Info("Refresh has successfully finished")
	c.JSON(http.StatusOK, gin.H{"Tokens": models.Tokens{
		AccessToken:  accessToken,
		RefreshToken: t.Token,
	}})
}

// Logout godoc
// @Summary Logouts user
// @Description Logouts user by ID
// @Tags user
// @Security ApiKeyAuth
// @Param user_id path string true "User ID"
// @Success 200 {string} string "User logged out successfully"
// @Failure 400 {object} string "Invalid user id"
// @Failure 500 {object} string "Server error while processing request"
// @Router /logout [post]
func (h *Handler) Logout(c *gin.Context) {
	h.Log.Info("Logout function is starting")

	UserId := c.Param("user_id")
	_, err := uuid.Parse(UserId)
	if err != nil {
		er := errors.Wrap(err, "invalid user id").Error()
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": er})
		h.Log.Error(er)
		return
	}

	ctx, cancel := context.WithTimeout(c, time.Second*5)
	defer cancel()

	err = h.RepoToken.Delete(ctx, UserId)
	if err != nil {
		er := errors.Wrap(err, "error deleting refresh token").Error()
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": er})
		h.Log.Error(er)
		return
	}

	h.Log.Info("Logout has successfully finished")
	c.JSON(http.StatusOK, "User logged out")

}
