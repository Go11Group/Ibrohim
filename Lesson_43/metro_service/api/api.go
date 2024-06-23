package api

import (
	"database/sql"
	"metro-service/api/handler"
	"net/http"
	"github.com/gin-gonic/gin"
)

func Routes(db *sql.DB) *http.Server {
	h := handler.NewHandler(db)

	r := gin.Default()
	mux := r.Group("/metro-service")

	c := mux.Group("/cards")
	{
		c.GET("/:id", h.GetUser)
		c.POST("", h.PostUser)
		c.PUT("/:id", h.PutUser)
		c.DELETE("/:id", h.DeleteUser)
		c.GET("/all", h.GetAllUsers)
	}
	s := mux.Group("/stations")
	{
		s.GET("/:id", h.GetStation)
		s.POST("", h.PostStation)
		s.PUT("/:id", h.PutStation)
		s.DELETE("/:id", h.DeleteStation)
		s.GET("/all", h.GetAllStations)
	}
	t := mux.Group("/terminals")
	{
		t.GET("/:id", h.GetTerminal)
		t.POST("", h.PostTerminal)
		t.PUT("/:id", h.PutTerminal)
		t.DELETE("/:id", h.DeleteTerminal)
		t.GET("/all", h.GetAllTerminals)
	}
	tr := mux.Group("/transactions")
	{
		tr.GET("/:id", h.GetTransaction)
		tr.POST("", h.PostTransaction)
		tr.PUT("/:id", h.PutTransaction)
		tr.DELETE("/:id", h.DeleteTransaction)
		tr.GET("/all", h.GetAllTransactions)
	}

	return &http.Server{Addr: ":8082", Handler: r}
}
