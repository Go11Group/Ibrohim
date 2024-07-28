package api

import (
	"redis-crud/api/handler"
	"redis-crud/api/middleware"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

func NewRouter(rdb *redis.Client) *gin.Engine {
	h := handler.NewHandler(rdb)
	r := gin.Default()

	r.Use(middleware.Check)

	r.POST("/person", h.Add)
	r.GET("/person/:id", h.Read)

	return r
}
