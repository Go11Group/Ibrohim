package api

import (
	"api-gateway/api/handler"
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Response(c *gin.Context, port string) {
	body, err := handler.Client(c, port)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	
	var respData map[string]interface{}
	err = json.Unmarshal(body, &respData)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "failed to parse response body"})
		return
	}

	c.JSON(http.StatusOK, respData)
}

func Routes() *http.Server {
	r := gin.Default()

	u := r.Group("/user-service")
	{
		u.GET("/:id", func(ctx *gin.Context) { Response(ctx, "1") })
		u.POST("", func(ctx *gin.Context) { Response(ctx, "1") })
		u.PUT("/:id", func(ctx *gin.Context) { Response(ctx, "1") })
		u.DELETE("/:id", func(ctx *gin.Context) { Response(ctx, "1") })
		u.GET("/all", func(ctx *gin.Context) { Response(ctx, "1") })
	}

	m := r.Group("/metro-service")
	c, s, t, tr := m.Group("/cards"), m.Group("/stations"), m.Group("/terminals"), m.Group("/transactions")
	{
		c.GET("/:id", func(ctx *gin.Context) { Response(ctx, "2") })
		c.POST("", func(ctx *gin.Context) { Response(ctx, "2") })
		c.PUT("/:id", func(ctx *gin.Context) { Response(ctx, "2") })
		c.DELETE("/:id", func(ctx *gin.Context) { Response(ctx, "2") })
		c.GET("/all", func(ctx *gin.Context) { Response(ctx, "2") })
	}
	{
		s.GET("/:id", func(ctx *gin.Context) { Response(ctx, "2") })
		s.POST("", func(ctx *gin.Context) { Response(ctx, "2") })
		s.PUT("/:id", func(ctx *gin.Context) { Response(ctx, "2") })
		s.DELETE("/:id", func(ctx *gin.Context) { Response(ctx, "2") })
		s.GET("/all", func(ctx *gin.Context) { Response(ctx, "2") })
	}
	{
		t.GET("/:id", func(ctx *gin.Context) { Response(ctx, "2") })
		t.POST("", func(ctx *gin.Context) { Response(ctx, "2") })
		t.PUT("/:id", func(ctx *gin.Context) { Response(ctx, "2") })
		t.DELETE("/:id", func(ctx *gin.Context) { Response(ctx, "2") })
		t.GET("/all", func(ctx *gin.Context) { Response(ctx, "2") })
	}
	{
		tr.GET("/:id", func(ctx *gin.Context) { Response(ctx, "2") })
		tr.POST("", func(ctx *gin.Context) { Response(ctx, "2") })
		tr.PUT("/:id", func(ctx *gin.Context) { Response(ctx, "2") })
		tr.DELETE("/:id", func(ctx *gin.Context) { Response(ctx, "2") })
		tr.GET("/all", func(ctx *gin.Context) { Response(ctx, "2") })
	}

	return &http.Server{Addr: ":8080", Handler: r}
}
