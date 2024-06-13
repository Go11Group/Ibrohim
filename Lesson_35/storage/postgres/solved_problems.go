package postgres

import (
	"database/sql"
	"gorilla_pg/model"
	"time"
)

type UserProblemRepo struct {
	DB *sql.DB
}

func NewUserProblemRepo(db *sql.DB) *UserProblemRepo {
	return &UserProblemRepo{DB: db}
}

func (up *UserProblemRepo) GetUserProblems(userID int) ([]model.Problem, error) {
	tr, err := BeginTransaction(up)
	if err != nil {
		return nil, err
	}
	defer CloseTransaction(tr, err)
	query := `
	SELECT 	p.id, p.title, p.description, p.difficulty
	FROM users u
	JOIN solved_problems sp ON u.id = sp.user_id
	JOIN problems p ON sp.problem_id = p.id
	WHERE u.id = $1`
	rows, err := up.DB.Query(query, userID)
	if err != nil {
		rows.Close()
		return nil, err
	}
	defer rows.Close()

	var problems []model.Problem
	for rows.Next() {
		var pr model.Problem
		err := rows.Scan(&pr.ID, &pr.Title, &pr.Description, &pr.Difficulty)
		if err != nil {
            return nil, err
        }
		problems = append(problems, pr)
	}
	return problems, nil
}

func (up *UserProblemRepo) AddProblemToUser(userID, problemID int, time time.Time) error {
	tr, err := BeginTransaction(up)
	if err != nil {
		return err
	}
	defer CloseTransaction(tr, err)
	query := `
	INSERT INTO solved_problems (user_id, problem_id, solved_at)
	VALUES ($1, $2, $3)
	ON CONFLICT (user_id, problem_id)
	DO UPDATE SET solved_at = $3`
	_, err = up.DB.Exec(query, userID, problemID, time)
	return err
}

func (up *UserProblemRepo) UpdateTimeOfSolution(userID, problemID int, time time.Time) error {
	tr, err := BeginTransaction(up)
	if err != nil {
		return err
	}
	defer CloseTransaction(tr, err)
	query := "update solved_problems set solved_at = $3 where user_id = $1 and problem_id = $2"
	_, err = up.DB.Exec(query, userID, problemID, time)
	return err
}

func (up *UserProblemRepo) RemoveProblemFromUser(userID, problemID int) error {
	tr, err := BeginTransaction(up)
	if err != nil {
		return err
	}
	defer CloseTransaction(tr, err)
	_, err = up.DB.Exec("delete from solved_problems where user_id = $1 and problem_id = $2",
	userID, problemID)
	return err
}