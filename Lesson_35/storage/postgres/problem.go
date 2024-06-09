package postgres

import (
	"database/sql"
	"errors"
	"gorilla_pg/model"
	"strconv"
)

type ProblemRepo struct {
	DB *sql.DB
}

func NewProblemRepo(db *sql.DB) *ProblemRepo {
	return &ProblemRepo{DB: db}
}

func (p *ProblemRepo) GetProblem(filter model.Problem) ([]model.Problem, error) {
	query := "select * from problems where 1=1"
	var params []interface{}
	paramIndex := 1
	if filter.ID > 0 {
		query += " and id = $" + strconv.Itoa(paramIndex)
		params = append(params, filter.ID)
		paramIndex++
	}
	if filter.Title != "" {
		query += " and title = $" + strconv.Itoa(paramIndex)
		params = append(params, filter.Title)
		paramIndex++
	}
	if filter.Description != "" {
		query += " and description = $" + strconv.Itoa(paramIndex)
		params = append(params, filter.Description)
		paramIndex++
	}
	if filter.Difficulty != "" {
		query += " and difficulty = $" + strconv.Itoa(paramIndex)
		params = append(params, filter.Difficulty)
		paramIndex++
	}
	if filter.Acceptance > 0 {
		query += " and acceptance = $" + strconv.Itoa(paramIndex)
		params = append(params, filter.Acceptance)
		paramIndex++
	}

	rows, err := p.DB.Query(query, params...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var problems []model.Problem
	var pr model.Problem
	for rows.Next() {
		err := rows.Scan(&pr.ID, &pr.Title, &pr.Description, &pr.Difficulty, &pr.Acceptance)
		if err != nil {
			return nil, err
		}
		problems = append(problems, pr)
	}
	return problems, nil
}

func (p *ProblemRepo) CreateProblem(problem model.Problem) error {
	if problem.Title == "" || problem.Description == "" || problem.Difficulty == "" || problem.Acceptance <= 0 {
		return errors.New("cannot insert empty fields")
	}
	_, err := p.DB.Exec("insert into problems(title, description, difficulty, acceptance) values($1,$2,$3,$4)",
	problem.Title, problem.Description, problem.Difficulty, problem.Acceptance)
	if err != nil {
		return err
	}
	return nil
}

func (p *ProblemRepo) UpdateProblem(problem model.Problem) error {
	query := "update problems set"
	var params []interface{}
	paramIndex := 1
	if problem.Title != "" {
		query += " title = $" + strconv.Itoa(paramIndex)
		params = append(params, problem.Title)
		paramIndex++
	}
	if problem.Description != "" {
		query += ", description = $" + strconv.Itoa(paramIndex)
		params = append(params, problem.Description)
		paramIndex++
	}
	if problem.Difficulty != "" {
		query += ", difficulty = $" + strconv.Itoa(paramIndex)
		params = append(params, problem.Difficulty)
		paramIndex++
	}
	if problem.Acceptance > 0 {
		query += ", acceptance = $" + strconv.Itoa(paramIndex)
		params = append(params, problem.Acceptance)
		paramIndex++
	}
	if paramIndex == 1 {
		return errors.New("no fields provided for update")
	}
	query += " where id = $" + strconv.Itoa(paramIndex)
	params = append(params, problem.ID)

	_, err := p.DB.Exec(query, params...)
	if err != nil {
		return err
	}
	return nil
}

func (p *ProblemRepo) DeleteProblem(id int) error {
	_, err := p.DB.Exec("delete from problems where id = $1", id)
	if err != nil {
		return err
	}
	return nil
}