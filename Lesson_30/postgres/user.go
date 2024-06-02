package postgres

import (
	model "User_Product/models"
	"database/sql"
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
	if filter.ID != 0 {
		query += " and id = $1"
		params = append(params, filter.ID)
	}
	if filter.Username != "" {
		query += " and username = $2"
		params = append(params, filter.Username)
	}
	if filter.Email != "" {
		query += " and email = $3"
		params = append(params, filter.Email)
	}
	if filter.Password != "" {
		query += " and password = $4"
		params = append(params, filter.Password)
	}

	rows, err := u.DB.Query(query, params...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []model.User
	var user model.User
	for rows.Next() {
		err = rows.Scan(&user.ID, &user.Username, &user.Email, &user.Password)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

func (u *UserRepo) CreateUser(user model.User) error {
	_, err := u.DB.Exec("insert into users(username, email, password) values($1, $2, $3)",
	user.Username, user.Email, user.Password)
	if err != nil {
		return err
	}
	return nil
}

func (u *UserRepo) UpdateUser(user model.User) error {
	query := "update users set"
	var params []interface{}
	if user.Username != "" {
		query += " username = $1"
		params = append(params, user.Username)
	}
	if user.Email != "" {
		query += " and email = $2"
		params = append(params, user.Email)
	}
	if user.Password != "" {
		query += " and password = $3"
		params = append(params, user.Password)
	}
	query += " where id = $4"
	params = append(params, user.ID)
	
	_, err := u.DB.Exec(query, params...)
	if err != nil {
		return err
	}
	return nil
}

func (u *UserRepo) DeleteUser(id int) error {
	_, err := u.DB.Exec("delete from users where id = $1", id)
	if err != nil {
		return err
	}
	return nil
}