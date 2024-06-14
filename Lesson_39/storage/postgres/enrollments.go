package postgres

import (
	"database/sql"
	"fmt"
	"language_learning_app/model"
	"time"

	"github.com/pkg/errors"
)

type EnrollmentRepo struct {
	DB *sql.DB
}

func NewEnrollmentRepo(db *sql.DB) *EnrollmentRepo {
	return &EnrollmentRepo{DB: db}
}

// CRUD operations
func (e *EnrollmentRepo) EnrollUserOnCourse(newData model.Enrollment) error {
	if newData.UserID == "" || newData.CourseID == "" || newData.EnrollmentDate == "" {
		return errors.New("error: cannot insert empty fields")
	}

	query := "insert into enrollments "
	params := []interface{}{newData.UserID, newData.CourseID, newData.EnrollmentDate}
	if newData.EnrollmentID != "" {
		query += "(enrollment_id, user_id, course_id, enrollment_date) values($1, $2, $3, $4)"
		params = append([]interface{}{newData.EnrollmentID}, params...)
	} else {
		query += "(user_id, course_id, enrollment_date) values($1, $2, $3)"
	}

	_, err := e.DB.Exec(query, params...)
	if err != nil {
		return errors.Wrap(err, "failed to enroll user on course")
	}
	return nil
}

func (e *EnrollmentRepo) Read(enrollmentID string) (*model.Enrollment, error) {
	isDel, err := IsDeleted(e.DB, "enrollments", "enrollment_id", enrollmentID)
	if err != nil {
		return nil, errors.Wrap(err, "enrollment not found")
	}
	if isDel {
		return nil, errors.New("enrollment deleted")
	}

	var enrollment model.Enrollment
	row := e.DB.QueryRow("select enrollment_id, user_id, course_id, enrollment_date from enrollments where enrollment_id = $1", enrollmentID)
	err = row.Scan(&enrollment.EnrollmentID, &enrollment.UserID, &enrollment.CourseID, &enrollment.EnrollmentDate)
	if err != nil {
		return nil, errors.Wrap(err, "enrollment not found")
	}
	return &enrollment, nil
}

func (e *EnrollmentRepo) Update(enrollmentID string, newData model.Enrollment) error {
	isDel, err := IsDeleted(e.DB, "enrollments", "enrollment_id", enrollmentID)
	if err != nil {
		return errors.Wrap(err, "enrollment not found")
	}
	if isDel {
		return errors.New("enrollment deleted")
	}

	query := "update enrollments set"
	var filter []string
	var params = make(map[string]interface{})
	
	if newData.UserID != "" {
		filter = append(filter, "user_id = :user_id")
		params["user_id"] = newData.UserID
	}
	if newData.CourseID != "" {
		filter = append(filter, "course_id = :course_id")
		params["course_id"] = newData.CourseID
	}
	if newData.EnrollmentDate != "" {
		filter = append(filter, "enrollment_date = :enrollment_date")
		params["enrollment_date"] = newData.EnrollmentDate
	}
	if len(filter) == 0 || len(params) == 0 {
		return errors.New("error: no fields provided for update")
	}
	
	filter = append(filter, "updated_at = :updated_at where enrollment_id = :enrollment_id")
	params["updated_at"] = time.Now()
	params["enrollment_id"] = enrollmentID

	q, p := ReplaceUpdateParams(query, filter, params)
	_, err = e.DB.Exec(q, p...)
	if err != nil {
		return errors.Wrap(err, "failed to update enrollment")
	}
	return nil
}

func (e *EnrollmentRepo) Delete(enrollmentID string) error {
	isDel, err := IsDeleted(e.DB, "enrollments", "enrollment_id", enrollmentID)
	if err != nil {
		return errors.Wrap(err, "enrollment not found")
	}
	if isDel {
		return errors.New("enrollment already deleted")
	}
	
	_, err = e.DB.Exec("update enrollments set deleted_at = date_part('epoch', current_timestamp)::INT where enrollment_id = $1", enrollmentID)
	if err != nil {
		return errors.Wrap(err, "failed to delete enrollment")
	}
	return nil
}

// Additional methods
func (e *EnrollmentRepo) GetAllEnrollments(f model.EnrollmentFilter) ([]model.Enrollment, error) {
	query := "select enrollment_id, user_id, course_id, enrollment_date from enrollments"
	var filter []string
	var params = make(map[string]interface{})
	
	if f.EnrollmentDate != "" {
		filter = append(filter, "date_trunc('day', enrollment_date) = :enrollment_date")
		params["enrollment_date"] = f.EnrollmentDate
	}
	if f.Limit > 0 {
		filter = append(filter, "LIMIT :limit")
		params["limit"] = f.Limit
	}
	if f.Offset > 0 {
		filter = append(filter, "OFFSET :offset")
		params["offset"] = f.Offset
	}
	query += " where deleted_at = 0"
	q, p := ReplaceReadParams(query, filter, params)

	fmt.Println(q, p)
	rows, err := e.DB.Query(q, p...)
	if err != nil {
		return nil, errors.Wrap(err, "failed to retrieve enrollments")
	}
	defer rows.Close()

	var enrollments []model.Enrollment
	for rows.Next() {
		var enrollment model.Enrollment
		err = rows.Scan(&enrollment.EnrollmentID, &enrollment.UserID, &enrollment.CourseID, &enrollment.EnrollmentDate)
		if err != nil {
			return nil, errors.Wrap(err, "failed to read enrollment")
		}
		enrollments = append(enrollments, enrollment)
	}
	return enrollments, nil
}