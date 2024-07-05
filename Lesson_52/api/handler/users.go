package handler

import (
	"log"
	"net/http"
	"time"
	"user_service/api/token"

	"github.com/gin-gonic/gin"
	// pbU "user_service/genproto/user"
	"encoding/json"
	pbAu "user_service/genproto/authentication"

	"golang.org/x/crypto/bcrypt"
)

func (h *Handler) Register(ctx *gin.Context) {
	req := pbAu.RegisterRequest{}

	err := json.NewDecoder(ctx.Request.Body).Decode(&req)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, err)
		log.Println(err)
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, err)
		log.Println(err)
		return
	}
	req.Password = string(hashedPassword)

	err = h.UserRepo.Register(&req)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, err)
		log.Println(err)
		return
	}

	ctx.JSON(http.StatusCreated, nil)
}

func (h *Handler) Login(ctx *gin.Context) {
	req := pbAu.LoginRequest{}

	if err := json.NewDecoder(ctx.Request.Body).Decode(&req); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, err)
		log.Println(err)
		return
	}

	user, err := h.UserRepo.GetUserByEmail(req.Email)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, err)
		log.Println(err)
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, err)
		log.Println(err)
		return
	}

	tokens := token.GenerateJWT(&pbAu.UserToken{
		Id:       user.Id,
		Username: user.Username,
		Email:    req.Email,
	})

	h.UserRepo.StoreRefreshToken(&pbAu.TokenRequest{
		UserId:    user.Id,
		Token:     tokens.RefreshToken,
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
	})

	ctx.JSON(http.StatusOK, tokens)
}

func (h *Handler) RefreshToken(ctx *gin.Context) {
	rft := ctx.PostForm("refresh_token")
	
	_, err := h.UserRepo.ValidateRefreshToken(rft)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, err)
		log.Println(err)
		return
	}

	newToken, err := token.GenerateAccessToken(rft)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, err)
		log.Println(err)
		return
	}

	ctx.JSON(http.StatusOK, newToken)
}

func (h *Handler) logout(c *gin.Context) {
	refreshToken := c.PostForm("refresh_token")
	if refreshToken == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Refresh token is required",
		})
		return
	}

	_, err := h.UserRepo.ValidateRefreshToken(refreshToken)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Invalid refresh token",
		})
		return
	}

	err = h.UserRepo.DeleteRefreshToken(refreshToken)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to logout",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Successfully logged out",
	})
}
