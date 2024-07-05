package api

import (
	"user_service/api/handler"

	"github.com/gin-gonic/gin"
)

func NewRouter(h *handler.Handler) *gin.Engine {
	router := gin.Default()

	users := router.Group("/users")
	users.POST("/register", h.Register)
	users.POST("/login", h.Login)

	return router
}
