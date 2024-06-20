package postgres

import (
	"database/sql"
	_ "github.com/lib/pq"
)

func ConnectDB() (*sql.DB, error) {
	con := "postgres://postgres:root@localhost:5432/postgres?sslmode=disable"
	db, err := sql.Open("postgres", con)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}