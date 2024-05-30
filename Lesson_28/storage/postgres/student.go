package postgres

import (
	"database/sql"
	"Stu_Cou/models"
)

type StudentRepo struct {
	DB *sql.DB
}

func NewStudentRepo(db *sql.DB) *StudentRepo {
	return &StudentRepo{DB: db}
}

func (s *StudentRepo) GetAllStudents() ([]model.Student, error) {
	rows, err := s.DB.Query(`select * from student`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var sts []model.Student
	var stu model.Student
	for rows.Next() {
		err = rows.Scan(&stu.ID, &stu.Name, &stu.Age)
		if err != nil {
			return nil, err
		}
		sts = append(sts, stu)
	}

	return sts, nil
}

func (s *StudentRepo) GetByID(id string) (*model.Student, error) {
	var stu model.Student
	
	err := s.DB.QueryRow(`select name, age from student where id = $1`, id).
		Scan(&stu.Name, &stu.Age)
	if err != nil {
		return nil, err
	}
	stu.ID = id
	return &stu, nil
}

func (s *StudentRepo) Create(stu model.Student) error {
	_, err := s.DB.Exec("insert into student(id, name, age) values ($1, $2, $3)", stu.ID, stu.Name, stu.Age)
	if err != nil {
		return err
	}
	return nil
}

func (s *StudentRepo) Update(id, name string, age int) error {
	_, err := s.DB.Exec("update student set name = $2, age = $3 where id = $1", id, name, age)
	if err != nil {
		return err
	}
	return nil
}

func (s *StudentRepo) Delete(id string) error {
	_, err := s.DB.Exec("delete from student where id = $1", id)
	if err != nil {
		return err
	}
	return nil
}