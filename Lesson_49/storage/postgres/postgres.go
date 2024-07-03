package postgres

import (
	"database/sql"
	"fmt"
	"podcast_service/config"
	_ "github.com/lib/pq"
)

func ConnectDB() (*sql.DB, error) {
	cfg := config.Load()

	conn := fmt.Sprintf("port = %s host=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.DB_PORT, cfg.DB_HOST, cfg.DB_USER, cfg.DB_PASSWORD, cfg.DB_NAME)

	db, err := sql.Open("postgres", conn)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}