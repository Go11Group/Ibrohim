package handler

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

// MakeOrder godoc
// @Summary Makes order
// @Description Makes order
// @Tags order
// @Security ApiKeyAuth
// @Success 200 {string} string "Order made successfully"
// @Failure 400 {object} string "Invalid user id"
// @Failure 500 {object} string "Server error while processing request"
// @Router /user/order [post]
func (h *Handler) MakeOrder(c *gin.Context) {
	h.Log.Info("MakeOrder handler is invoked")

	id, ok := c.Get("user_id")
	if !ok {
		er := errors.New("user id not provided").Error()
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": er})
		h.Log.Error(er)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), h.ContextTimeout)
	defer cancel()

	err := h.KafkaProducer.Produce(ctx, h.OrderTopic, []byte(id.(string)))
	if err != nil {
		er := errors.Wrap(err, "error making order").Error()
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": er})
		h.Log.Error(er)
		return
	}

	c.JSON(http.StatusOK, "Order made successfully")
}
