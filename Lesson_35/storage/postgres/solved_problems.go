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
	tr, err := up.DB.Begin()
	if err != nil {
		return nil, err
	}
	defer func() {
		if err != nil {
			tr.Rollback()
		} else {
			tr.Commit()
		}
	}()

	query := `
	SELECT 	p.id, p.title, p.description, p.difficulty,
            sp.solved_at
    FROM users u
    JOIN solved_problems sp ON u.id = sp.user_id
    JOIN problems p ON sp.product_id = p.id
	WHERE u.id = $1`
	rows, err := up.DB.Query(query, userID)
	if err != nil {
		rows.Close()
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

func (up *UserProblemRepo) AddProblemToUser(userID, problemID int, time time.Time) error {
	tr, err := up.DB.Begin()
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			tr.Rollback()
		} else {
			tr.Commit()
		}
	}()

	query := `insert into solved_problems(user_id, problem_id, solved_at)
	values($1, $2, $3)
	on conflict (user_id, product_id)
	do update set solved_at = $3`
	_, err = up.DB.Exec(query, userID, problemID, time)
	if err != nil {
		return err
	}
	return nil
}

func (up *UserProblemRepo) UpdateTimeOfSolving(userID, problemID int, time time.Time) error {
	tr, err := up.DB.Begin()
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			tr.Rollback()
		} else {
			tr.Commit()
		}
	}()
	
	query := "update solved_problems set solved_at = $3 where user_id = $1 and problem_id = $2"
	_, err = up.DB.Exec(query, userID, problemID, time)
	if err != nil {
		return err
	}
	return nil
}

func (up *UserProblemRepo) RemoveProblemFromUser(userID, problemID int) error {
	tr, err := up.DB.Begin()
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			tr.Rollback()
		} else {
			tr.Commit()
		}
	}()
	
	_, err = up.DB.Exec("delete from solved_problems where user_id = $1 and problem_id = $2",
	userID, problemID)
	if err != nil {
		return err
	}
	return nil
}