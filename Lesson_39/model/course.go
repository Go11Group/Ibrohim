package model

type Course struct {
	CourseID    string `form:"course_id" json:"course_id"`
	Title       string `form:"title" json:"title"`
	Description string `form:"description" json:"description"`
}

// for filtering courses
type CourseFilter struct {
	Title       string
	Description string
	Limit       int
	Offset      int
}

// for json responses
type CourseResponse struct {
	CourseID    string `json:"course_id"`
	Title       string `json:"title"`
	Enrollments int    `json:"enrollments_count"`
}
