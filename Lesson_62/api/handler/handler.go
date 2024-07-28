package handler

import (
	"redis-crud/models"
	rds "redis-crud/redis"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/redis/go-redis/v9"
)

type Handler struct {
	repo *rds.PersonRepo
}

func NewHandler(r *redis.Client) *Handler {
	return &Handler{repo: rds.NewPersonRepo(r)}
}

func (h *Handler) Add(c *gin.Context) {
	var data models.PersonInfo
	if err := c.ShouldBindJSON(&data); err != nil {
		er := errors.Wrap(err, "error binding json").Error()
		c.JSON(400, gin.H{"error": er})
		return
	}

	res, err := h.repo.Add(c, &data)
	if err != nil {
		er := errors.Wrap(err, "error adding person").Error()
		c.JSON(500, gin.H{"error": er})
		return
	}

	c.JSON(200, res)
}

func (h *Handler) Read(c *gin.Context) {
	id := c.Param("id")
	res, err := h.repo.Read(c, id)
	if err != nil {
		er := errors.Wrap(err, "error reading person").Error()
		c.JSON(500, gin.H{"error": er})
		return
	}

	c.JSON(200, res)
}
