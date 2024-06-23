package postgres

import (
	"database/sql"
	"strconv"
	"strings"
	_ "github.com/lib/pq"
	"github.com/pkg/errors"
)

func ConnectDB() (*sql.DB, error) {
	con := "postgres://postgres:root@localhost:5432/metro_system?sslmode=disable"
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

func replaceParams(query string, params map[string]interface{}) (string, []interface{}) {
	i := 1
	var args []interface{}

	for k, v := range params {
		if k != "" && strings.Contains(query, ":"+k) {
			query = strings.ReplaceAll(query, ":"+k, "$"+strconv.Itoa(i))
			args = append(args, v)
			i++
		}
	}

	return query, args
}

// a function to construct query with fields for reading
func ReplaceReadParams(query string, filter []string, params map[string]interface{}) (string, []interface{}) {
	for _, f := range filter {
		if strings.Contains(f, "LIMIT") || strings.Contains(f, "OFFSET") { // checking for limit and offset
			query += " " + f
			continue
		}
		query += " and " + f
	}

	return replaceParams(query, params)
}

// a function to construct query with fields for updating
func ReplaceUpdateParams(query string, filter []string, params map[string]interface{}) (string, []interface{}) {
	for i, f := range filter {
		if i == 0 || i == len(filter)-1 {
			query += " " + f
			continue
		}
		query += ", " + f
	}

	return replaceParams(query, params)
}
