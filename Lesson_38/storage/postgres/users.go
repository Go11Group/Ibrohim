package postgres

import (
	"database/sql"
	"language_learning_app/model"
	"time"
	"github.com/pkg/errors"
)

type UserRepo struct {
	DB *sql.DB
}

func NewUserRepo(db *sql.DB) *UserRepo {
	return &UserRepo{DB: db}
}

// CRUD operations
func (u *UserRepo) Create(newData model.User) error {
	if newData.Name == "" || newData.Birthday == "" || newData.Email == "" || newData.Password == "" {
		return errors.New("error: cannot insert empty fields")
	}
	query := "insert into users "
	params := []interface{}{newData.Name, newData.Birthday, newData.Email, newData.Password}
	if newData.UserID != "" {
		query += "(user_id, name, birthday, email, password) values($1, $2, $3, $4, $5)"
		params = append([]interface{}{newData.UserID}, params...)
	} else {
		query += "(name, birthday, email, password) values($1, $2, $3, $4)"
	}
	_, err := u.DB.Exec(query, params...)
	if err != nil {
		return errors.Wrap(err, "failed to insert user into database")
	}
	return nil
}

func (u *UserRepo) Read(userID string) (*model.User, error) {
	isDel, err := IsDeleted(u.DB, "users", "user_id", userID)
	if err != nil {
		return nil, errors.Wrap(err, "user not found")
	}
	if isDel {
		return nil, errors.New("user deleted")
	}
	var user model.User
	row := u.DB.QueryRow("select user_id, name, birthday, email, password from users where user_id = $1", userID)
	err = row.Scan(&user.UserID, &user.Name, &user.Birthday, &user.Email, &user.Password)
	if err != nil {
		return nil, errors.Wrap(err, "user not found")
	}
	return &user, nil
}

func (u *UserRepo) Update(userID string, newData model.User) error {
	isDel, err := IsDeleted(u.DB, "users", "user_id", userID)
	if err != nil {
		return errors.Wrap(err, "user not found")
	}
	if isDel {
		return errors.New("user deleted")
	}
	query := "update users set"
	var filter []string
	var params = make(map[string]interface{})
	if newData.Name != "" {
		filter = append(filter, "name = :name")
		params["name"] = newData.Name
	}
	if newData.Birthday != "" {
		filter = append(filter, "birthday = :birthday")
		params["birthday"] = newData.Birthday
	}
	if newData.Email != "" {
		filter = append(filter, "email = :email")
		params["email"] = newData.Email
	}
	if newData.Password != "" {
		filter = append(filter, "password = :password")
		params["password"] = newData.Password
	}
	if len(filter) == 0 || len(params) == 0 {
		return errors.New("error: no fields provided for update")
	}
	filter = append(filter, "updated_at = :updated_at where user_id = :user_id and deleted_at = 0")
	params["updated_at"] = time.Now()
	params["user_id"] = userID
	q, p := ReplaceUpdateParams(query, filter, params)
	_, err = u.DB.Exec(q, p...)
	if err != nil {
		return errors.Wrap(err, "failed to update user")
	}
	return nil
}

func (u *UserRepo) Delete(userID string) error {
	isDel, err := IsDeleted(u.DB, "users", "user_id", userID)
	if err != nil {
		return errors.Wrap(err, "user not found")
	}
	if isDel {
		return errors.New("user already deleted")
	}
	_, err = u.DB.Exec("update users set deleted_at = date_part('epoch', current_timestamp)::INT where user_id = $1", userID)
	if err != nil {
		return errors.Wrap(err, "failed to delete user")
	}
	return nil
}


// Additional methods
func (u *UserRepo) GetAllUsers(f model.UserFilter) ([]model.User, error) {
	query := "select user_id, name, birthday, email, password from users"
	var filter []string
	var params = make(map[string]interface{})
	if f.Name != "" {
		filter = append(filter, "name = :name")
		params["name"] = f.Name
	}
	if f.AgeFrom > 0 {
		yearTo := time.Now().AddDate(-f.AgeFrom, 0, 0)
		filter = append(filter, "birthday <= :yearTo")
		params["yearTo"] = yearTo
	}
	if f.AgeTo > 0 {
		yearFrom := time.Now().AddDate(-(f.AgeTo+1), 0, 0)
		filter = append(filter, "birthday >= :yearFrom")
		params["yearFrom"] = yearFrom
	}
	if f.Limit > 0 {
		filter = append(filter, "LIMIT :limit")
		params["limit"] = f.Limit
	}
	if f.Offset > 0 {
		filter = append(filter, "OFFSET :offset")
		params["offset"] = f.Offset
	}
	query += " where deleted_at = 0"
	q, p := ReplaceReadParams(query, filter, params)
	rows, err := u.DB.Query(q, p...)
	if err != nil {
		return nil, errors.Wrap(err, "failed to retrieve users")
	}
	var users []model.User
	for rows.Next() {
		var user model.User
		err = rows.Scan(&user.UserID, &user.Name, &user.Birthday, &user.Email, &user.Password)
		if err != nil {
			return nil, errors.Wrap(err, "failed to read user")
		}
		users = append(users, user)
	}
	return users, nil
}