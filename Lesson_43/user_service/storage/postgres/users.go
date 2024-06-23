package postgres

import (
	"database/sql"
	"user-service/model"
	"github.com/pkg/errors"
)

type UserRepo struct {
	DB *sql.DB
}

func NewUserRepo(db *sql.DB) *UserRepo {
	return &UserRepo{DB: db}
}

func (ur *UserRepo) Create(u *model.User) error {
	if u.Name == "" || u.Age < 1 || u.PhoneNumber == "" {
		return errors.New("error: cannot insert empty fields")
	}

	query := "insert into users "
	params := []interface{}{u.Name, u.Age, u.PhoneNumber}
	if u.ID != "" {
		query += "(id, name, age, phone_number) values($1, $2, $3, $4)"
		params = append([]interface{}{u.ID}, params...)
	} else {
		query += "(name, age, phone_number) values($1, $2, $3)"
	}

	_, err := ur.DB.Exec(query, params...)
	if err != nil {
		return errors.Wrap(err, "failed to insert user into database")
	}

	return nil
}

func (ur *UserRepo) Read(userID string) (*model.User, error) {
	u := model.User{ID: userID}

	row := ur.DB.QueryRow("select name, age, phone_number from users where id = $1", userID)
	err := row.Scan(&u.Name, &u.Age, &u.PhoneNumber)
	if err != nil {
		return nil, errors.Wrap(err, "user not found")
	}

	return &u, nil
}

func (ur *UserRepo) Update(u *model.User) error {
	if u.ID == "" && u.Name == "" && u.Age < 1 && u.PhoneNumber == "" {
		return errors.New("error: no fields provided for update")
	}

	query := "update users set"
	var filter []string
	var params = make(map[string]interface{})

	if u.Name != "" {
		filter = append(filter, "name = :name")
		params["name"] = u.Name
	}
	if u.Age > 0 {
		filter = append(filter, "age = :age")
		params["age"] = u.Age
	}
	if u.PhoneNumber != "" {
		filter = append(filter, "phone_number = :phone_number")
		params["phone_number"] = u.PhoneNumber
	}
	filter = append(filter, " where id = :id")
	params["id"] = u.ID
	
	q, p := ReplaceUpdateParams(query, filter, params)
	_, err := ur.DB.Exec(q, p...)
	if err != nil {
		return errors.Wrap(err, "failed to update user")
	}

	return nil
}

func (ur *UserRepo) Delete(userID string) error {
	_, err := ur.DB.Exec("delete from users where id = $1", userID)
	if err != nil {
		return errors.Wrap(err, "failed to delete user")
	}

	return nil
}

func (ur *UserRepo) FetchUsers(f model.UserFilter) ([]model.User, error) {
	query := "select id, name, age, phone_number from users where true"
	var filter []string
	var params = make(map[string]interface{})

	if f.Name != "" {
		filter = append(filter, "name = :name")
		params["name"] = f.Name
	}
	if f.AgeFrom > 0 {
		filter = append(filter, "age >= :ageFrom")
		params["ageFrom"] = f.AgeFrom
	}
	if f.AgeTo > 0 {
		filter = append(filter, "age <= :ageTo")
		params["ageTo"] = f.AgeTo
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
	rows, err := ur.DB.Query(q, p...)
	if err != nil {
		return nil, errors.Wrap(err, "failed to retrieve users")
	}
	defer rows.Close()

	var users []model.User
	for rows.Next() {
		var u model.User
		err = rows.Scan(&u.ID, &u.Name, &u.Age, &u.PhoneNumber)
		if err != nil {
			return nil, errors.Wrap(err, "failed to read user")
		}
		users = append(users, u)
	}

	return users, nil
}
