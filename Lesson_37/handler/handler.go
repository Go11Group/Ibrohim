package handler

import (
	"les37/postgres"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	Person *postgres.PersonRepo
}

func NewHandler(h Handler) *gin.Engine {
	r := gin.Default()
	r.GET("/person_info/getAll", h.getAll)
	return r
}