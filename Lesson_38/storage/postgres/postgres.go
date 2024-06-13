package postgres

import (
	"database/sql"
	"fmt"
	"strconv"
	"strings"
	_ "github.com/lib/pq"
	"github.com/pkg/errors"
)

const (
	host = "localhost"
	port = 5432
	user = "postgres"
	dbname = "language_learning_app"
	password = "root"
)
// a function to connect to the database
func ConnectDB() (*sql.DB, error) {
	con := fmt.Sprintf("host = %s port = %d user = %s dbname = %s password = %s sslmode=disable",
	host, port, user, dbname, password)
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

// a function to check if a record is marked as deleted in the database using the table, column and id
func IsDeleted(db *sql.DB, tb, col string, id string) (bool, error) {
	var deleted_at int
	query := fmt.Sprintf("select deleted_at from %s where %s = $1", tb, col)
	err := db.QueryRow(query, id).Scan(&deleted_at)
	if err != nil {
		if err == sql.ErrNoRows { // checking the absence of rows
			return false, nil
		}
		return false, err // checking any errors
	}
	return deleted_at != 0, nil // returning the state (deleted = any int | not deleted = 0)
}

// a function to replace query params (:columnName) with placeholders ($1, $2...) and construct a slice with values
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
		if i == 0 {
			query += " " + f
			continue
		}
		query += ", " + f
	}
	return replaceParams(query, params)
}