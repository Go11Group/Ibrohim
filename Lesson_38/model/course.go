package model

type Course struct {
	CourseID    string `form:"course_id"`
    Title       string `form:"title"`
    Description string `form:"description"`
}

type CourseFilter struct {
    Title string
    Description string
    Limit int
    Offset int
}