package postgres

import (
	"database/sql"
	"Stu_Cou/models"
)

type CourseRepo struct {
	DB *sql.DB
}

func NewCourseRepo(db *sql.DB) *CourseRepo {
	return &CourseRepo{DB: db}
}

func (c *CourseRepo) GetAllCourses() ([]model.Course, error) {
	rows, err := c.DB.Query("select * from course")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var cs []model.Course
	var cou model.Course
	for rows.Next() {
		err = rows.Scan(&cou.ID, &cou.Name)
		if err != nil {
			return nil, err
		}
		cs = append(cs, cou)
	}

	return cs, nil
}

func (c *CourseRepo) GetByID(id string) (*model.Course, error) {
	var cou model.Course

	err := c.DB.QueryRow("select name from course where id = $1",id).Scan(&cou.Name)
	if err != nil {
		return nil, err
	}
	cou.ID = id
	return &cou, nil
}

func (c *CourseRepo) Create(cou *model.Course) error {
	_, err := c.DB.Exec("insert into course(id, name) values($1, $2)", cou.ID, cou.Name) 
	if err != nil {
		return err
	}
	return nil
}

func (c *CourseRepo) Update(id, name string) error {
	_, err := c.DB.Exec("update course set name = $2 where id = $1", id, name)
	if err != nil {
		return err
	}
	return nil
}

func (c *CourseRepo) Delete(id string) error {
	_, err := c.DB.Exec("delete from course where id = $1", id)
	if err != nil {
		return err
	}
	return nil
}