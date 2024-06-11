package postgres

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

const (
	host = "localhost"
	port = 5432
	user = "postgres"
	dbname = "user_problem"
	password = "root"
)

func ConnectDB() (*sql.DB, error) {
	con := fmt.Sprintf("host = %s port = %d user = %s dbname = %s password = %s sslmode=disable",
	host, port, user, dbname, password)
	db, err := sql.Open("postgres", con)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return db, nil
}

func BeginTransaction(up *UserProblemRepo) (*sql.Tx, error) {
	tr, err := up.DB.Begin()
	if err != nil {
		return nil, err
	}
	return tr, nil
}

func CloseTransaction(tr *sql.Tx, err error) {
	if err != nil {
		tr.Rollback()
	} else {
		tr.Commit()
	}
}