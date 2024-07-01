package postgres

import (
	"database/sql"
	"fmt"
	"time"
	pb "user_service/genproto/user"
)

type UserRepo struct {
	Db *sql.DB
}

func NewUserRepo(db *sql.DB) *UserRepo {
	return &UserRepo{Db: db}
}

func (u *UserRepo) GetUserById(userId string) (*pb.User, error) {
	query := `select id, username, email, password_hash, 
	created_at, updated_at from users where id = $1`

	user := pb.User{Id: userId}
	err := u.Db.QueryRow(query, userId).Scan(&user.Username,
		&user.Email, &user.Password, &user.CreatedAt,
		&user.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (u *UserRepo) UpdateUserProfile(profile *pb.Profile) error {
	query := `
	update
	    user_profiles 
	set
    	full_name = $1,
    	bio = $2,
   	 	role = $3,
    	location = $4,
    	avatar_image = $5,
    	website = $6,
	    updated_at = now(),
	WHERE 
	    id = $7;
`
	res, err := u.Db.Exec(query, profile.FullName, profile.Bio,
		profile.Location, profile.AvatarImage, profile.Website)
	affectedRows, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if affectedRows == 0 {
		return fmt.Errorf("user %s not found", profile.UserId)
	}

	return nil
}

func (u *UserRepo) UpdateUser(user *pb.User) error {
	query := "update users set "
	params := []interface{}{}

	if user.Username != "" {
		query += fmt.Sprintf("username = $%d, ", len(params)+1)
		params = append(params, user.Username)
	}
	if user.Email != "" {
		query += fmt.Sprintf("email = $%d, ", len(params)+1)
		params = append(params, user.Email)
	}
	if user.Password != "" {
		query += fmt.Sprintf("password_hash = $%d, ", len(params)+1)
	}
	query += fmt.Sprintf("updated_at = $%d ", len(params)+1)
	params = append(params, time.Now())
	query += fmt.Sprintf("where id = $%d", len(params)+1)
	params = append(params, user.Id)

	_, err := u.Db.Exec(query, params...)
	if err != nil {
		return err
	}

	return nil
}

func (u *UserRepo) DeleteUser(id string) error {
	query := `
	update
	    users (u *UserRepo) 
	set
	    deleted_at = now(),
	WHERE 
	    id = $1;
`
	res, err := u.Db.Exec(query, id)
	affectedRows, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if affectedRows == 0 {
		return fmt.Errorf("user %s not found", id)
	}

	return nil
}

func (u *UserRepo) GetUserProfile(id string) (*pb.Profile, error) {
	query := `select full_name, bio, role, location, avatar_image, website
	from user_profiles where user_id = $1`

	profile := &pb.Profile{UserId: id}
	row := u.Db.QueryRow(query, id)
	err := row.Scan(&profile.FullName, &profile.Bio, &profile.Role, &profile.Location, &profile.AvatarImage, &profile.Website)
	if err != nil {
		return nil, err
	}

	return profile, nil
}