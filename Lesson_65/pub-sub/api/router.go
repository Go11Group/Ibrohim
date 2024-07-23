package api

import (
	"stocks-management/api/handler"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

func NewRouter(rclient *redis.Client) *gin.Engine {
	h := handler.NewHandler(rclient)
	router := gin.Default()

	router.GET("/stock/:company", h.GetStockPrice)

	return router
}
