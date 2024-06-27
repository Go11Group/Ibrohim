package postgres

import (
	"database/sql"

	_ "github.com/lib/pq"
	"github.com/pkg/errors"
)

func ConnectDB() (*sql.DB, error) {
	con := "postgres://postgres:root@localhost:5432/postgres?sslmode=disable"
	db, err := sql.Open("postgres", con)
	if err != nil {
		return nil, errors.Wrap(err, "failed to connect to database")
	}

	if err := db.Ping(); err != nil {
		db.Close()
		return nil, errors.Wrap(err, "connection to database is not alive")
	}

	return db, nil
}
