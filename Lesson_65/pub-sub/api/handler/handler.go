package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"

	rdb "stocks-management/redis"
)

type Handler struct {
	redisClient *redis.Client
}

func NewHandler(redisClient *redis.Client) *Handler {
	return &Handler{
		redisClient: redisClient,
	}
}

func (h *Handler) GetStockPrice(c *gin.Context) {
	stockName := c.Param("company")

	if stockName == "" {
		c.AbortWithStatusJSON(400, gin.H{"error": "company name is required"})
		return
	}

	stockPrice, err := rdb.GetStockPrice(c, stockName)
	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{"error getting stock price": err.Error()})
		return
	}

	c.JSON(200, gin.H{stockName: stockPrice})
}
