package model

type Enrollment struct {
	EnrollmentID    string `form:"enrollment_id"`
    UserID          string `form:"user_id"`
    CourseID        string `form:"course_id"`
    EnrollmentDate  string `form:"enrollment_date"`
}

type EnrollmentFilter struct {
	UserID string
	CourseID string
	EnrollmentDate string
	Limit int
    Offset int
}