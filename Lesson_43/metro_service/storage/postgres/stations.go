package postgres

import (
	"database/sql"
	"metro-service/model"
	"github.com/pkg/errors"
)

type StationRepo struct {
	DB *sql.DB
}

func NewStationRepo(db *sql.DB) *StationRepo {
	return &StationRepo{DB: db}
}

func (sr *StationRepo) Create(s *model.Station) error {
	if s.Name == "" {
		return errors.New("error: cannot insert empty fields")
	}

	query := "insert into stations "
	params := []interface{}{s.Name}
	if s.ID != "" {
		query += "(id, name) values($1, $2)"
		params = append([]interface{}{s.ID}, params...)
	} else {
		query += "(name) values($1)"
	}

	_, err := sr.DB.Exec(query, params...)
	if err != nil {
		return errors.Wrap(err, "failed to insert station into database")
	}

	return nil
}

func (sr *StationRepo) Read(stationID string) (*model.Station, error) {
	s := model.Station{ID: stationID}

	row := sr.DB.QueryRow("select name from cards where id = $1", stationID)
	err := row.Scan(&s.Name)
	if err != nil {
		return nil, errors.Wrap(err, "station not found")
	}

	return &s, nil
}

func (sr *StationRepo) Update(s *model.Station) error {
	if s.ID == "" && s.Name == "" {
		return errors.New("error: no fields provided for update")
	}
	
	query := "update stations set"
	filter := []string{"name = :name", " where id = :id"}
	params := map[string]interface{}{"name": s.Name, "id": s.ID}
	
	q, p := ReplaceUpdateParams(query, filter, params)
	_, err := sr.DB.Exec(q, p...)
	if err != nil {
		return errors.Wrap(err, "failed to update station")
	}

	return nil
}

func (sr *StationRepo) Delete(stationID string) error {
	_, err := sr.DB.Exec("delete from stations where id = $1", stationID)
	if err != nil {
		return errors.Wrap(err, "failed to delete station")
	}

	return nil
}

func (sr *StationRepo) FetchStations(f model.StationFilter) ([]model.Station, error) {
	query := "select id, name from stations where true"
	var filter []string
	var params = make(map[string]interface{})

	if f.Name != "" {
		filter = append(filter, "name = :name")
		params["name"] = f.Name
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
	rows, err := sr.DB.Query(q, p...)
	if err != nil {
		return nil, errors.Wrap(err, "failed to retrieve stations")
	}
	defer rows.Close()

	var stations []model.Station
	for rows.Next() {
		var s model.Station
		err = rows.Scan(&s.ID, &s.Name)
		if err != nil {
			return nil, errors.Wrap(err, "failed to read station")
		}
		stations = append(stations, s)
	}

	return stations, nil
}
