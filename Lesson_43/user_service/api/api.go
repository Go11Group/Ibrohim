package api

import (
	"database/sql"
	"net/http"
	"user-service/api/handler"
	"github.com/gin-gonic/gin"
)

func Routes(db *sql.DB) *http.Server {
	h := handler.NewHandler(db)

	r := gin.Default()
	mux := r.Group("/user-service")
	{
		mux.GET("/:id", h.Get)
		mux.POST("", h.Post)
		mux.PUT("/:id", h.Put)
		mux.DELETE("/:id", h.Delete)
		mux.GET("/all", h.GetAll)
	}

	return &http.Server{Addr: ":8081", Handler: r}
}
