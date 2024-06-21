package model

type Enrollment struct {
	EnrollmentID   string `form:"enrollment_id" json:"enrollment_id"`
	UserID         string `form:"user_id" json:"user_id"`
	CourseID       string `form:"course_id" json:"course_id"`
	EnrollmentDate string `form:"enrollment_date" json:"enrollment_date"`
}

// for filtering enrollemnts
type EnrollmentFilter struct {
	EnrollmentDate string
	Limit          int
	Offset         int
}
