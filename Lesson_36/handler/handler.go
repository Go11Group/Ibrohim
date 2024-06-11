package handler

import (
	"fmt"
	"gin_pg/model"
	"gin_pg/storage/postgres"
	"net/http"
	"strconv"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	User *postgres.UserRepo
	Problem *postgres.ProblemRepo
	UserProblem *postgres.UserProblemRepo
}

func NewHandler(h Handler) *gin.Engine {
	r := gin.Default()
	ginRoutes := r.Group("/gin")
	{
		u := ginRoutes.Group("/user")
		u.GET("/:id", h.userGet)
		u.POST("", h.userPost)
		u.PUT("/:id", h.userPut)
		u.DELETE("/:id", h.userDelete)

		p := ginRoutes.Group("/problem")
		p.GET("/:id", h.problemGet)
		p.POST("", h.problemPost)
		p.PUT("/:id", h.problemPut)
		p.DELETE("/:id", h.problemDelete)

		ginRoutes.GET("/user-problems/:id", h.userProblemGet)
		ginRoutes.POST("/user-problems", h.userProblemPost)
		ginRoutes.PUT("/user-problems", h.userProblemPut)
		ginRoutes.DELETE("/user-problems", h.userProblemDelete)
	}
	return r
}

func GetID(c *gin.Context, name string) (int, error) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, fmt.Errorf("invalid %s ID", name))
		fmt.Println("Error: ", err)
		return 0, err
	}
	return id, nil
}

func ReadUserBody(c *gin.Context, u *model.User) error {
	u.Username = c.PostForm("username")
	u.Email = c.PostForm("email")
	u.Password = c.PostForm("password")
	return nil
}

func ReadProblemBody(c *gin.Context, p *model.Problem) error {
	p.Title = c.PostForm("title")
	p.Description = c.PostForm("description")
	p.Difficulty = c.DefaultPostForm("difficulty", "Easy")
	acceptance, err := strconv.ParseFloat(c.DefaultPostForm("acceptance", "1"), 32)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, fmt.Errorf("invalid acceptance rate"))
		fmt.Println("Error: ", err)
		return err
	}
	p.Acceptance = float32(acceptance)
	return nil
}