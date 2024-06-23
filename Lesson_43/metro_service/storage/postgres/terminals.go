package postgres

import (
	"database/sql"
	"metro-service/model"
	"github.com/pkg/errors"
)

type TerminalRepo struct {
	DB *sql.DB
}

func NewTerminalRepo(db *sql.DB) *TerminalRepo {
	return &TerminalRepo{DB: db}
}

func (tr *TerminalRepo) Create(t *model.Terminal) error {
	if t.StationID == "" {
		return errors.New("error: cannot insert empty fields")
	}

	query := "insert into terminals "
	params := []interface{}{t.StationID}
	if t.ID != "" {
		query += "(id, station_id) values($1, $2)"
		params = append([]interface{}{t.ID}, params...)
	} else {
		query += "(station_id) values($1)"
	}

	_, err := tr.DB.Exec(query, params...)
	if err != nil {
		return errors.Wrap(err, "failed to insert terminal into database")
	}

	return nil
}

func (tr *TerminalRepo) Read(terminalID string) (*model.Terminal, error) {
	t := model.Terminal{ID: terminalID}

	row := tr.DB.QueryRow("select station_id from terminals where id = $1", terminalID)
	err := row.Scan(&t.StationID)
	if err != nil {
		return nil, errors.Wrap(err, "terminal not found")
	}

	return &t, nil
}

func (tr *TerminalRepo) Update(t *model.Terminal) error {
	if t.ID == "" && t.StationID == "" {
		return errors.New("error: no fields provided for update")
	}

	query := "update terminals set"
	filter := []string{"station_id = :station_id", " where id = :id"}
	params := map[string]interface{}{"station_id": t.StationID, "id": t.ID}

	q, p := ReplaceUpdateParams(query, filter, params)
	_, err := tr.DB.Exec(q, p...)
	if err != nil {
		return errors.Wrap(err, "failed to update terminal")
	}

	return nil
}

func (tr *TerminalRepo) Delete(terminalID string) error {
	_, err := tr.DB.Exec("delete from terminals where id = $1", terminalID)
	if err != nil {
		return errors.Wrap(err, "failed to delete terminal")
	}

	return nil
}

func (tr *TerminalRepo) FetchTerminals(f model.TerminalFilter) ([]model.Terminal, error) {
	query := "select id, station_id from terminals where true"
	var filter []string
	var params = make(map[string]interface{})

	if f.StationID != "" {
		filter = append(filter, "station_id = :station_id")
		params["station_id"] = f.StationID
	}
	if f.Limit > 0 {
		filter = append(filter, "LIMIT :limit")
		params["limit"] = f.Limit
	}
	if f.Offset > 0 {
		filter = append(filter, "OFFSET :offset")
		params["offset"] = f.Offset
	}

	q, p := ReplaceReadParams(query, filter, params)
	rows, err := tr.DB.Query(q, p...)
	if err != nil {
		return nil, errors.Wrap(err, "failed to retrieve terminals")
	}
	defer rows.Close()

	var terminals []model.Terminal
	for rows.Next() {
		var t model.Terminal
		err = rows.Scan(&t.ID, &t.StationID)
		if err != nil {
			return nil, errors.Wrap(err, "failed to read terminal")
		}
		terminals = append(terminals, t)
	}

	return terminals, nil
}
