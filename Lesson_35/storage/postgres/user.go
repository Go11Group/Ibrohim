package postgres

import (
	"database/sql"
	"errors"
	"gorilla_pg/model"
	"strconv"
)

type UserRepo struct {
	DB *sql.DB
}

func NewUserRepo(db *sql.DB) *UserRepo {
	return &UserRepo{DB: db}
}

func (u *UserRepo) GetUser(filter model.User) ([]model.User, error) {
	query := "select * from users where 1=1"
	var params []interface{}
	if filter.ID > 0 {
		query += " and id = $" + strconv.Itoa(len(params) + 1)
		params = append(params, filter.ID)
	}
	if filter.Username != "" {
		query += " and username = $" + strconv.Itoa(len(params) + 1)
		params = append(params, filter.Username)
	}
	if filter.Email != "" {
		query += " and email = $" + strconv.Itoa(len(params) + 1)
		params = append(params, filter.Email)
	}
	if filter.Password != "" {
		query += " and password = $" + strconv.Itoa(len(params) + 1)
		params = append(params, filter.Password)
	}

	rows, err := u.DB.Query(query, params...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []model.User
	for rows.Next() {
		var user model.User
		err = rows.Scan(&user.ID, &user.Username, &user.Email, &user.Password)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

func (u *UserRepo) CreateUser(user model.User) error {
	if user.Username == "" || user.Email == "" || user.Password == "" {
		return errors.New("cannot insert empty fields")
	}
	_, err := u.DB.Exec("insert into users(username, email, password) values($1, $2, $3)",
	user.Username, user.Email, user.Password)
	return err
}

func (u *UserRepo) UpdateUser(user model.User) error {
	query := "update users set"
	var params []interface{}
	if user.Username != "" {
		query += " username = $" + strconv.Itoa(len(params) + 1)
		params = append(params, user.Username)
	}
	if user.Email != "" {
		query += ", email = $" + strconv.Itoa(len(params) + 1)
		params = append(params, user.Email)
	}
	if user.Password != "" {
		query += ", password = $" + strconv.Itoa(len(params) + 1)
		params = append(params, user.Password)
	}
	if len(params) == 0 {
		return errors.New("no fields provided for update")
	}
	query += " where id = $" + strconv.Itoa(len(params) + 1)
	params = append(params, user.ID)
	
	_, err := u.DB.Exec(query, params...)
	return err
}

func (u *UserRepo) DeleteUser(id int) error {
	_, err := u.DB.Exec("delete from users where id = $1", id)
	return err
}