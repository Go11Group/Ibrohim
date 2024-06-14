package model

type Lesson struct {
	LessonID string `form:"lesson_id" json:"lesson_id"`
	CourseID string `form:"course_id" json:"course_id"`
	Title    string `form:"title" json:"title"`
	Content  string `form:"content" json:"content"`
}

// for filtering lessons
type LessonFilter struct {
	Title   string
	Content string
	Limit   int
	Offset  int
}

// for json responses
type LessonResponse struct {
	LessonID string `json:"lesson_id"`
	Title    string `json:"title"`
	Content  string `json:"content"`
}
