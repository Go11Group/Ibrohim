package handler

import (
	"database/sql"
	"language_learning_app/storage/postgres"
	"github.com/gin-gonic/gin"
)

// a struct to hold database methods
type Handler struct {
	User       *postgres.UserRepo
	Course     *postgres.CourseRepo
	Lesson     *postgres.LessonRepo
	Enrollment *postgres.EnrollmentRepo
}

// a function returning a handler with 4 repos
func NewHandler(db *sql.DB) *Handler {
	return &Handler{
		User:       postgres.NewUserRepo(db),
		Course:     postgres.NewCourseRepo(db),
		Lesson:     postgres.NewLessonRepo(db),
		Enrollment: postgres.NewEnrollmentRepo(db),
	}
}

// a function returning a route
func NewRoute(h Handler) *gin.Engine {
	r := gin.Default()
	llaRoutes := r.Group("/language_learning_app")
	{
		u := llaRoutes.Group("/users")
		{
			u.GET("/:user_id", h.UserGet)
			u.POST("", h.UserPost)
			u.PUT("/:user_id", h.UserPut)
			u.DELETE("/:user_id", h.UserDelete)

			u.GET("/get-all", h.UsersGet)
			u.GET("/:user_id/courses", h.UserCoursesGet)
			u.GET("/search", h.UserSearch)
		}

		c := llaRoutes.Group("/courses")
		{
			c.GET("/:course_id", h.CourseGet)
			c.POST("", h.CoursePost)
			c.PUT("/:course_id", h.CoursePut)
			c.DELETE("/:course_id", h.CourseDelete)

			c.GET("/get-all", h.CoursesGet)
			c.GET("/:course_id/lessons", h.CourseLessons)
			c.GET("/:course_id/enrollments", h.CourseUsers)
			c.GET("/popular", h.CoursePopular)
		}

		l := llaRoutes.Group("/lessons")
		{
			l.GET("/:lesson_id", h.LessonGet)
			l.POST("", h.LessonPost)
			l.PUT("/:lesson_id", h.LessonPut)
			l.DELETE("/:lesson_id", h.LessonDelete)

			l.GET("/get-all", h.LessonsGet)
		}

		e := llaRoutes.Group("/enrollments")
		{
			e.GET("/:enrollment_id", h.EnrollmentGet)
			e.POST("", h.EnrollmentPost)
			e.PUT("/:enrollment_id", h.EnrollmentPut)
			e.DELETE("/:enrollment_id", h.EnrollmentDelete)

			e.GET("/get-all", h.EnrollmentsGet)
		}
	}
	return r
}
