package handler

import (
	pb "api-gateway/genproto/product"
	"context"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

// PostProduct godoc
// @Summary Creates a product
// @Description Adds new product
// @Tags product
// @Security ApiKeyAuth
// @Param product body product.NewProduct true "Product data"
// @Success 200 {object} product.InsertResp
// @Failure 400 {object} string "Invalid data provided"
// @Failure 500 {object} string "Server error while processing request"
// @Router /admin/product [post]
func (h *Handler) PostProduct(c *gin.Context) {
	h.Log.Info("PostProduct handler is invoked")

	var req pb.NewProduct
	if err := c.ShouldBindJSON(&req); err != nil {
		er := errors.Wrap(err, "invalid data provided").Error()
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": er})
		h.Log.Error(er)
		return
	}

	ctx, cancel := context.WithTimeout(c, h.ContextTimeout)
	defer cancel()

	resp, err := h.Product.CreateProduct(ctx, &req)
	if err != nil {
		er := errors.Wrap(err, "error creating product").Error()
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": er})
		h.Log.Error(er)
		return
	}

	h.Log.Info("PostProduct handler is completed")
	c.JSON(http.StatusOK, gin.H{"New product": resp})
}

// GetProduct godoc
// @Summary Gets product
// @Description Retrieves product
// @Tags product
// @Security ApiKeyAuth
// @Param id path string true "Product ID"
// @Success 200 {object} product.ProductInfo
// @Failure 400 {object} string "Invalid data provided"
// @Failure 500 {object} string "Server error while processing request"
// @Router /admin/product/{id} [get]
func (h *Handler) GetProduct(c *gin.Context) {
	h.Log.Info("GetProduct handler is invoked")

	id := c.Param("id")
	if id == "" {
		er := errors.New("product id not provided").Error()
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": er})
		h.Log.Error(er)
		return
	}

	ctx, cancel := context.WithTimeout(c, h.ContextTimeout)
	defer cancel()

	resp, err := h.Product.GetProductById(ctx, &pb.Id{Id: id})
	if err != nil {
		er := errors.Wrap(err, "error getting product").Error()
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": er})
		h.Log.Error(er)
		return
	}

	h.Log.Info("GetProduct handler is completed")
	c.JSON(http.StatusOK, gin.H{"Product": resp})
}

// PutProduct godoc
// @Summary Updates product
// @Description Updates product info
// @Tags product
// @Security ApiKeyAuth
// @Param id path string true "Product ID"
// @Param product body product.NewDataNoId true "Product data"
// @Success 200 {object} product.UpdateResp
// @Failure 400 {object} string "Invalid data provided"
// @Failure 500 {object} string "Server error while processing request"
// @Router /admin/product/{id} [put]
func (h *Handler) PutProduct(c *gin.Context) {
	h.Log.Info("PutProduct handler is invoked")

	id := c.Param("id")
	if id == "" {
		er := errors.New("product id not provided").Error()
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": er})
		h.Log.Error(er)
		return
	}

	var req pb.NewDataNoId
	if err := c.ShouldBindJSON(&req); err != nil {
		er := errors.Wrap(err, "invalid data provided").Error()
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": er})
		h.Log.Error(er)
		return
	}

	ctx, cancel := context.WithTimeout(c, h.ContextTimeout)
	defer cancel()

	resp, err := h.Product.UpdateProduct(ctx, &pb.NewData{
		Id:          id,
		Name:        req.Name,
		Description: req.Description,
		Category:    req.Category,
		Price:       req.Price,
		Stock:       req.Stock,
		Discount:    req.Discount,
	})
	if err != nil {
		er := errors.Wrap(err, "error updating product").Error()
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": er})
		h.Log.Error(er)
		return
	}

	h.Log.Info("PutProduct handler is completed")
	c.JSON(http.StatusOK, gin.H{"Updated product": resp})
}

// DeleteProduct godoc
// @Summary Deletes product
// @Description Deletes product
// @Tags product
// @Security ApiKeyAuth
// @Param id path string true "Product ID"
// @Success 200 {object} string "Product deleted successfully"
// @Failure 400 {object} string "Invalid data provided"
// @Failure 500 {object} string "Server error while processing request"
// @Router /admin/product/{id} [delete]
func (h *Handler) DeleteProduct(c *gin.Context) {
	h.Log.Info("DeleteProduct handler is invoked")

	id := c.Param("id")
	if id == "" {
		er := errors.New("product id not provided").Error()
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": er})
		h.Log.Error(er)
		return
	}

	ctx, cancel := context.WithTimeout(c, h.ContextTimeout)
	defer cancel()

	_, err := h.Product.DeleteProduct(ctx, &pb.Id{Id: id})
	if err != nil {
		er := errors.Wrap(err, "error deleting product").Error()
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": er})
		h.Log.Error(er)
		return
	}

	h.Log.Info("DeleteProduct handler is completed")
	c.JSON(http.StatusOK, "Product deleted successfully")
}

// FetchProducts godoc
// @Summary Gets products
// @Description Retrieves products
// @Tags product
// @Security ApiKeyAuth
// @Param name query string false "Name"
// @Param category query string false "Category"
// @Param comment_count query int false "Comment count"
// @Param rating query float32 false "Rating"
// @Param sort_by query string false "Sort by"
// @Param page query int false "Page"
// @Param limit query int false "Limit"
// @Success 200 {object} product.Products
// @Failure 400 {object} string "Invalid pagination parameters"
// @Failure 500 {object} string "Server error while processing request"
// @Router /admin/product/all [get]
func (h *Handler) FetchProducts(c *gin.Context) {
	h.Log.Info("GetProducts handler is invoked")

	name := c.Query("name")
	category := c.Query("category")
	comment := c.Query("comment_count")
	ratingStr := c.Query("rating")
	sortBy := c.Query("sort_by")
	pageStr := c.Query("page")
	limitStr := c.Query("limit")

	commentCount, err := parseIntQueryParam(comment)
	if err != nil {
		er := errors.Wrap(err, "invalid comments count parameter").Error()
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": er})
		h.Log.Error(er)
		return
	}

	rating, err := parseFloatQueryParam(ratingStr)
	if err != nil {
		er := errors.Wrap(err, "invalid rating parameter").Error()
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": er})
		h.Log.Error(er)
		return
	}

	page, err := parseIntQueryParam(pageStr)
	if err != nil {
		er := errors.Wrap(err, "invalid pagination parameter").Error()
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": er})
		h.Log.Error(er)
		return
	}

	limit, err := parseIntQueryParam(limitStr)
	if err != nil {
		er := errors.Wrap(err, "invalid pagination parameter").Error()
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": er})
		h.Log.Error(er)
		return
	}

	filter := pb.Filter{
		Name:         name,
		Category:     category,
		CommentCount: commentCount,
		Rating:       rating,
		Page:         page,
		Limit:        limit,
	}

	if sortBy != "" {
		switch sortBy {
		case "most_purchased":
			filter.MostPurchased = true
		case "most_commented":
			filter.MostCommented = true
		case "most_recent":
			filter.MostRecent = true
		case "cheapest":
			filter.Cheapest = true
		case "most_expensive":
			filter.MostExpensive = true
		case "discount":
			filter.Discount = true
		default:
			er := errors.New("invalid sort by parameter").Error()
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": er})
			h.Log.Error(er)
			return
		}
	}

	ctx, cancel := context.WithTimeout(c, h.ContextTimeout)
	defer cancel()

	resp, err := h.Product.FetchProducts(ctx, &filter)
	if err != nil {
		er := errors.Wrap(err, "error getting products").Error()
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": er})
		h.Log.Error(er)
		return
	}

	h.Log.Info("GetProducts handler is completed")
	c.JSON(http.StatusOK, gin.H{"Products": resp})
}

func parseIntQueryParam(queryParam string) (int32, error) {
	if queryParam == "" {
		return 0, nil
	}
	value, err := strconv.Atoi(queryParam)
	if err != nil || value < 1 {
		return 0, errors.New("invalid integer parameter")
	}
	return int32(value), nil
}

func parseFloatQueryParam(queryParam string) (float32, error) {
	if queryParam == "" {
		return 0, nil
	}
	value, err := strconv.ParseFloat(queryParam, 32)
	if err != nil || value < 0 || value > 5 {
		return 0, errors.New("invalid float parameter")
	}
	return float32(value), nil
}
