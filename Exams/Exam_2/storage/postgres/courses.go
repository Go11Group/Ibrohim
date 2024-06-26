package postgres

import (
	"database/sql"
	"language_learning_app/model"
	"time"
	"github.com/pkg/errors"
)

type CourseRepo struct {
	DB *sql.DB
}

func NewCourseRepo(db *sql.DB) *CourseRepo {
	return &CourseRepo{DB: db}
}

// CRUD operations
func (c *CourseRepo) Create(newData model.Course) error {
	if newData.Title == "" || newData.Description == "" {
		return errors.New("error: cannot insert empty fields")
	}

	query := "insert into courses "
	params := []interface{}{newData.Title, newData.Description}
	if newData.CourseID != "" {
		query += "(course_id, title, description) values($1, $2, $3)"
		params = append([]interface{}{newData.CourseID}, params...)
	} else {
		query += "(title, description) values($1, $2)"
	}

	_, err := c.DB.Exec(query, params...)
	if err != nil {
		return errors.Wrap(err, "failed to insert course into database")
	}
	return nil
}

func (c *CourseRepo) Read(courseID string) (*model.Course, error) {
	isDel, err := IsDeleted(c.DB, "courses", "course_id", courseID)
	if err != nil {
		return nil, errors.Wrap(err, "course not found")
	}
	if isDel {
		return nil, errors.New("course deleted")
	}

	var course model.Course
	row := c.DB.QueryRow("select course_id, title, description from courses where course_id = $1", courseID)
	err = row.Scan(&course.CourseID, &course.Title, &course.Description)
	if err != nil {
		return nil, errors.Wrap(err, "course not found")
	}
	return &course, nil
}

func (c *CourseRepo) Update(courseID string, newData model.Course) error {
	isDel, err := IsDeleted(c.DB, "courses", "course_id", courseID)
	if err != nil {
		return errors.Wrap(err, "course not found")
	}
	if isDel {
		return errors.New("course deleted")
	}

	query := "update courses set"
	var filter []string
	var params = make(map[string]interface{})

	if newData.Title != "" {
		filter = append(filter, "title = :title")
		params["title"] = newData.Title
	}
	if newData.Description != "" {
		filter = append(filter, "description = :description")
		params["description"] = newData.Description
	}
	if len(filter) == 0 || len(params) == 0 {
		return errors.New("error: no fields provided for update")
	}

	filter = append(filter, "updated_at = :updated_at where course_id = :course_id")
	params["updated_at"] = time.Now()
	params["course_id"] = courseID

	q, p := ReplaceUpdateParams(query, filter, params)
	_, err = c.DB.Exec(q, p...)
	if err != nil {
		return errors.Wrap(err, "failed to update course")
	}
	return nil
}

func (c *CourseRepo) Delete(courseID string) error {
	isDel, err := IsDeleted(c.DB, "courses", "course_id", courseID)
	if err != nil {
		return errors.Wrap(err, "course not found")
	}
	if isDel {
		return errors.New("course already deleted")
	}
	
	// starting a transactoin to make sure all data related to a course is deleted
	tr, err := c.DB.Begin()
	if err != nil {
		return errors.Wrap(err, "failed to delete course")
	}
	defer func() {
		if err != nil {
			tr.Rollback() // rollback in case of errors
		} else {
			tr.Commit() // commit if no errors occurred
		}
	}()

	_, err = c.DB.Exec("update courses set deleted_at = date_part('epoch', current_timestamp)::INT where course_id = $1", courseID)
	if err != nil {
		return errors.Wrap(err, "failed to delete course info")
	}

	_, err = c.DB.Exec("update lessons set deleted_at = date_part('epoch', current_timestamp)::INT where course_id = $1", courseID)
	if err != nil {
		return errors.Wrap(err, "failed to delete course lessons")
	}

	_, err = c.DB.Exec("update enrollments set deleted_at = date_part('epoch', current_timestamp)::INT where course_id = $1", courseID)
	if err != nil {
		return errors.Wrap(err, "failed to delete course users")
	}

	return nil
}

// Additional methods
func (c *CourseRepo) GetAllCourses(f model.CourseFilter) ([]model.Course, error) {
	query := "select course_id, title, description from courses"
	var filter []string
	var params = make(map[string]interface{})

	if f.Title != "" {
		filter = append(filter, "title = :title")
		params["title"] = f.Title
	}
	if f.Description != "" {
		filter = append(filter, "description = :description")
		params["description"] = f.Description
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

	rows, err := c.DB.Query(q, p...)
	if err != nil {
		return nil, errors.Wrap(err, "failed to retrieve courses")
	}
	defer rows.Close()
	
	var courses []model.Course
	for rows.Next() {
		var course model.Course
		err = rows.Scan(&course.CourseID, &course.Title, &course.Description)
		if err != nil {
			return nil, errors.Wrap(err, "failed to read course")
		}
		courses = append(courses, course)
	}
	return courses, nil
}

func (c *CourseRepo) GetLessons(courseID string) ([]model.LessonResponse, error) {
	isDel, err := IsDeleted(c.DB, "courses", "course_id", courseID)
	if err != nil {
		return nil, errors.Wrap(err, "course not found")
	}
	if isDel {
		return nil, errors.New("course deleted")
	}

	query := `
	select l.lesson_id, l.title, l.content
	from lessons l
	join courses c
	on l.course_id = c.course_id
	where l.deleted_at = 0 and c.course_id = $1
	group by l.lesson_id;`
	
	rows, err := c.DB.Query(query, courseID)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to retrieve lessons of course_id %s", courseID)
	}
	defer rows.Close()

	var lessons []model.LessonResponse
	for rows.Next() {
		var lesson model.LessonResponse
		err := rows.Scan(&lesson.LessonID, &lesson.Title, &lesson.Content)
		if err != nil {
			return nil, errors.Wrapf(err, "failed to read lesson of course_id %s", courseID)
		}
		lessons = append(lessons, lesson)
	}
	return lessons, nil
}

func (c *CourseRepo) GetUsers(courseID string) ([]model.UserResponse, error) {
	isDel, err := IsDeleted(c.DB, "courses", "course_id", courseID)
	if err != nil {
		return nil, errors.Wrap(err, "course not found")
	}
	if isDel {
		return nil, errors.New("course deleted")
	}

	query := `
	select u.user_id, u.name, u.email
	from users u
	join enrollments e
	on u.user_id = e.user_id
	where u.deleted_at = 0 and e.deleted_at = 0 and e.course_id = $1
	group by u.user_id;`

	rows, err := c.DB.Query(query, courseID)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to retrieve users enrolled in course_id %s", courseID)
	}
	defer rows.Close()

	var users []model.UserResponse
	for rows.Next() {
		var user model.UserResponse
		err := rows.Scan(&user.UserID, &user.Name, &user.Email)
		if err != nil {
			return nil, errors.Wrapf(err, "failed to read user enrolled in course_id %s", courseID)
		}
		users = append(users, user)
	}
	return users, nil
}

func (c *CourseRepo) Popular(start_date, end_date string) ([]model.CourseResponse, error) {
	if start_date == "" || end_date == "" {
		return nil, errors.New("invalid dates")
	}

	query := `
	select c.course_id, c.title, count(e.user_id) as enrollments_count
	from courses c
	join enrollments e
	on c.course_id = e.course_id
	where c.deleted_at = 0 and e.deleted_at = 0 and e.enrollment_date between $1 and $2
	group by c.course_id
	order by enrollments_count desc;`

	rows, err := c.DB.Query(query, start_date, end_date)
	if err != nil {
		return nil, errors.Wrap(err, "failed to retrieve popular courses")
	}
	defer rows.Close()

	var courses []model.CourseResponse
	for rows.Next() {
		var course model.CourseResponse
		err := rows.Scan(&course.CourseID, &course.Title, &course.Enrollments)
		if err != nil {
			return nil, errors.Wrap(err, "failed to read popular course")
		}
		courses = append(courses, course)
	}
	return courses, nil
}