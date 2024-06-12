package postgres

import (
	"database/sql"
	"les37/model"
	"strconv"
	"strings"
)

type PersonRepo struct {
	DB *sql.DB
}

func NewPersonRepo(db *sql.DB) *PersonRepo {
	return &PersonRepo{DB: db}
}

func (p *PersonRepo) GetAll(f model.Filter) ([]model.Person, error) {
	var (
		params = make(map[string]interface{})
		arr    []interface{}
		limit  string
		offset string
	)
	query := "select id, first_name, last_name, age, nationality, field, city from person_info"
	filter := " where true"
	if f.ID > 0 {
		filter += " and id = :id"
		params["id"] = f.ID
	}
	if f.First_name != "" {
		filter += " and first_name = :first_name"
		params["first_name"] = f.First_name
	}
	if f.Last_name != "" {
		filter += " and last_name = :last_name"
		params["last_name"] = f.Last_name
	}
	if f.Age > 0 {
		filter += " and age = :age"
		params["age"] = f.Age
	}
	if f.Nationality != "" {
		filter += " and nationality = :nationality"
		params["nationality"] = f.Nationality
	}
	if f.Field != "" {
		filter += " and field = :field"
		params["field"] = f.Field
	}
	if f.City != "" {
		filter += " and city = :city"
		params["city"] = f.City
	}
	if f.Limit > 0 {
		params["limit"] = f.Limit
		limit = " LIMIT :limit"
	}
	if f.Offset > 0 {
		params["offset"] = f.Offset
		offset = " OFFSET :offset"
	}
	query = query + filter + limit + offset
	query, arr = ReplaceQueryParams(query, params)
	rows, err := p.DB.Query(query, arr...)
	if err != nil {
		return nil, err
	}
	var people []model.Person
	for rows.Next() {
		var p model.Person
		err := rows.Scan(&p.ID, &p.First_name, &p.Last_name, &p.Age, &p.Nationality, &p.Field, &p.City)
		if err != nil {
			return nil, err
		}
		people = append(people, p)
	}
	return people, nil
}

func ReplaceQueryParams(namedQuery string, params map[string]interface{}) (string, []interface{}) {
	index := 1
	var args []interface{}
	for k,v := range params {
		if k != "" && strings.Contains(namedQuery, ":"+k) {
			namedQuery = strings.ReplaceAll(namedQuery, ":"+k, "$" + strconv.Itoa(index))
			args = append(args, v)
			index++
		}
	}
	return namedQuery, args
}