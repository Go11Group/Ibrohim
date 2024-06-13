package postgres

import (
	"database/sql"
	"language_learning_app/model"
	"time"
	"github.com/pkg/errors"
)

type LessonRepo struct {
	DB *sql.DB
}

func NewLessonRepo(db *sql.DB) *LessonRepo {
	return &LessonRepo{DB: db}
}

// CRUD operations
func (l *LessonRepo) Create(newData model.Lesson) error {
	if newData.CourseID == "" || newData.Title == "" || newData.Content == "" {
		return errors.New("error: cannot insert empty fields")
	}

	query := "insert into lessons "
	params := []interface{}{newData.CourseID, newData.Title, newData.Content}
	if newData.LessonID != "" {
		query += "(lesson_id, course_id, title, content) values($1, $2, $3, $4)"
		params = append([]interface{}{newData.LessonID}, params...)
	} else {
		query += "(course_id, title, content) values($1, $2, $3)"
	}

	_, err := l.DB.Exec(query, params...)
	if err != nil {
		return errors.Wrap(err, "failed to insert lesson into database")
	}
	return nil
}

func (l *LessonRepo) Read(lessonID string) (*model.Lesson, error) {
	isDel, err := IsDeleted(l.DB, "lessons", "lesson_id", lessonID)
	if err != nil {
		return nil, errors.Wrap(err, "lesson not found")
	}
	if isDel {
		return nil, errors.New("lesson deleted")
	}

	var lesson model.Lesson
	row := l.DB.QueryRow("select lesson_id, course_id, title, content from lessons where lesson_id = $1", lessonID)
	err = row.Scan(&lesson.LessonID, &lesson.CourseID, &lesson.Title, &lesson.Content)
	if err != nil {
		return nil, errors.Wrap(err, "lesson not found")
	}
	return &lesson, nil
}

func (l *LessonRepo) Update(lessonID string, newData model.Lesson) error {
	isDel, err := IsDeleted(l.DB, "lessons", "lesson_id", lessonID)
	if err != nil {
		return errors.Wrap(err, "lesson not found")
	}
	if isDel {
		return errors.New("lesson deleted")
	}
	if !isDel {
		return errors.New("lesson not found")
	}

	query := "update lessons set"
	var filter []string
	var params = make(map[string]interface{})

	if newData.CourseID != "" {
		filter = append(filter, "course_id = :course_id")
		params["course_id"] = newData.CourseID
	}
	if newData.Title != "" {
		filter = append(filter, "title = :title")
		params["title"] = newData.Title
	}
	if newData.Content != "" {
		filter = append(filter, "content = :content")
		params["content"] = newData.Content
	}
	if len(filter) == 0 || len(params) == 0 {
		return errors.New("error: no fields provided for update")
	}

	filter = append(filter, "updated_at = :updated_at where lesson_id = :lesson_id and deleted_at = 0")
	params["updated_at"] = time.Now()
	params["lesson_id"] = lessonID

	q, p := ReplaceUpdateParams(query, filter, params)
	_, err = l.DB.Exec(q, p...)
	if err != nil {
		return errors.Wrap(err, "failed to update lesson")
	}
	return nil
}

func (l *LessonRepo) Delete(lessonID string) error {
	isDel, err := IsDeleted(l.DB, "lessons", "lesson_id", lessonID)
	if err != nil {
		return errors.Wrap(err, "lesson not found")
	}
	if isDel {
		return errors.New("lesson already deleted")
	}
	if !isDel {
		return errors.New("lesson not found")
	}
	
	_, err = l.DB.Exec("update lessons set deleted_at = date_part('epoch', current_timestamp)::INT where lesson_id = $1", lessonID)
	if err != nil {
		return errors.Wrap(err, "failed to delete lesson")
	}
	return nil
}

// Additional methods
func (l *LessonRepo) GetAllLessons(f model.LessonFilter) ([]model.Lesson, error) {
	query := "select lesson_id, course_id, title, content from lessons"
	var filter []string
	var params = make(map[string]interface{})
	
	if f.Title != "" {
		filter = append(filter, "title = :title")
		params["title"] = f.Title
	}
	if f.Content != "" {
		filter = append(filter, "content = :content")
		params["content"] = f.Content
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

	rows, err := l.DB.Query(q, p...)
	if err != nil {
		return nil, errors.Wrap(err, "failed to retrieve lessons")
	}
	defer rows.Close()

	var lessons []model.Lesson
	for rows.Next() {
		var lesson model.Lesson
		err = rows.Scan(&lesson.LessonID, &lesson.CourseID, &lesson.Title, &lesson.Content)
		if err != nil {
			return nil, errors.Wrap(err, "failed to read lesson")
		}
		lessons = append(lessons, lesson)
	}
	return lessons, nil
}