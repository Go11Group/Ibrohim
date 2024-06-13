package handler

import (
	"language_learning_app/storage/postgres"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	User *postgres.UserRepo
	Course *postgres.CourseRepo
	Lesson *postgres.LessonRepo
	Enrollment *postgres.EnrollmentRepo
}

func NewHandler(h Handler) *gin.Engine {
	r := gin.Default()
	llaRoutes := r.Group("/language_learning_app")
	{
		u := llaRoutes.Group("/users")
		u.GET("/get/:id", h.UserGet)
		u.POST("/post/", h.UserPost)
		u.PUT("/put/:id", h.UserPut)
		u.DELETE("/delete/:id", h.UserDelete)
		u.GET("/get-all", h.AllUsersGet)

		c := llaRoutes.Group("/courses")
		c.GET("/get/:id", h.CourseGet)
		c.POST("/post/", h.CoursePost)
		c.PUT("/put/:id", h.CoursePut)
		c.DELETE("/delete/:id", h.CourseDelete)
		c.GET("/get-all", h.AllCoursesGet)

		// l := llaRoutes.Group("/lessons")

		// e := llaRoutes.Group("/enrollments")
	}
	return r
}