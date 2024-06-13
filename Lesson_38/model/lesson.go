package model

type Lesson struct {
	LessonID   string `form:"lesson_id"`
    CourseID   string `form:"course_id"`
    Title      string `form:"title"`
    Content    string `form:"content"`
}

type LessonFilter struct {
    CourseID string
    Title string
    Content string
    Limit int
    Offset int
}