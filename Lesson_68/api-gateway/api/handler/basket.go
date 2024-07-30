package handler

import (
	pb "api-gateway/genproto/basket"
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

// AddToBasket godoc
// @Summary Adds product
// @Description Adds product to basket
// @Tags basket
// @Security ApiKeyAuth
// @Param product body basket.NewProduct true "Product data"
// @Success 200 {object} string "Product added to basket successfully"
// @Failure 400 {object} string "Invalid data provided"
// @Failure 500 {object} string "Server error while processing request"
// @Router /user/basket [post]
func (h *Handler) AddToBasket(c *gin.Context) {
	h.Log.Info("AddToBasket handler is invoked")

	var req pb.NewProduct
	if err := c.ShouldBindJSON(&req); err != nil {
		er := errors.Wrap(err, "invalid data provided").Error()
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": er})
		h.Log.Error(er)
		return
	}

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

	_, err := h.Basket.AddProduct(ctx, &req)
	if err != nil {
		er := errors.Wrap(err, "error adding product to basket").Error()
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": er})
		h.Log.Error(er)
		return
	}

	h.Log.Info("AddToBasket is completed")
	c.JSON(http.StatusOK, "Product added to basket successfully")
}

// GetBasket godoc
// @Summary Gets basket
// @Description Retrieves basket
// @Tags basket
// @Security ApiKeyAuth
// @Success 200 {object} basket.Products
// @Failure 400 {object} string "Invalid user id"
// @Failure 500 {object} string "Server error while processing request"
// @Router /user/basket [get]
func (h *Handler) GetBasket(c *gin.Context) {
	h.Log.Info("GetBasket handler is invoked")

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

	resp, err := h.Basket.GetProducts(ctx, &pb.Void{})
	if err != nil {
		er := errors.Wrap(err, "error getting basket").Error()
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": er})
		h.Log.Error(er)
		return
	}

	h.Log.Info("GetBasket handler is completed")
	c.JSON(http.StatusOK, gin.H{"basket": resp})
}

// UpdateBasket godoc
// @Summary Updates basket
// @Description Updates quantity of product in basket
// @Tags basket
// @Security ApiKeyAuth
// @Param product body basket.Quantity true "Product data"
// @Success 200 {object} string "Basket updated successfully"
// @Failure 400 {object} string "Invalid data provided"
// @Failure 500 {object} string "Server error while processing request"
// @Router /user/basket [put]
func (h *Handler) UpdateBasket(c *gin.Context) {
	h.Log.Info("UpdateBasket handler is invoked")

	var req pb.Quantity
	if err := c.ShouldBindJSON(&req); err != nil {
		er := errors.Wrap(err, "invalid data provided").Error()
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": er})
		h.Log.Error(er)
		return
	}

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

	_, err := h.Basket.UpdateProduct(ctx, &req)
	if err != nil {
		er := errors.Wrap(err, "error updating basket").Error()
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": er})
		h.Log.Error(er)
		return
	}

	h.Log.Info("UpdateBasket handler is completed")
	c.JSON(http.StatusOK, "Basket updated successfully")
}

// RemoveFromBasket godoc
// @Summary Removes product from basket
// @Description Removes product from basket
// @Tags basket
// @Security ApiKeyAuth
// @Param product_id body basket.Id true "Product ID"
// @Success 200 {object} string "Product removed from basket successfully"
// @Failure 400 {object} string "Invalid data provided"
// @Failure 500 {object} string "Server error while processing request"
// @Router /user/basket [delete]
func (h *Handler) RemoveFromBasket(c *gin.Context) {
	h.Log.Info("RemoveFromBasket handler is invoked")

	var req pb.Id
	if err := c.ShouldBindJSON(&req); err != nil {
		er := errors.Wrap(err, "invalid data provided").Error()
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": er})
		h.Log.Error(er)
		return
	}

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

	_, err := h.Basket.RemoveProduct(ctx, &req)
	if err != nil {
		er := errors.Wrap(err, "error removing product from basket").Error()
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": er})
		h.Log.Error(er)
		return
	}

	h.Log.Info("RemoveFromBasket handler is completed")
	c.JSON(http.StatusOK, "Product removed from basket successfully")
}
