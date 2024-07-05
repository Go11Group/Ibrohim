package postgres

import (
	"database/sql"
	"fmt"
	"time"
	pbAu "user_service/genproto/authentication"
	pb "user_service/genproto/user"
)

type UserRepo struct {
	Db *sql.DB
}

func NewUserRepo(db *sql.DB) *UserRepo {
	return &UserRepo{Db: db}
}

func (u *UserRepo) Register(user *pbAu.RegisterRequest) error {
	query := `
	insert into users(
		username, email, password
	) values(
	 	$1, $2, $3
	)`
	_, err := u.Db.Exec(query, user.Username, user.Email, user.Password)
	if err != nil {
		return err
	}
	return nil
}

func (u *UserRepo) StoreRefreshToken(token *pbAu.TokenRequest) error {
	tx, err := u.Db.Begin()
	if err != nil {
		return err
	}
	defer tx.Commit()

	query := `
	insert into refresh_tokens(user_id, token, expires_at)
	values ($1, $2, $3)
	`
	_, err = tx.Exec(query, token.UserId, token.Token, token.ExpiresAt)
	return err
}

func (u *UserRepo) ValidateRefreshToken(token string) (string, error) {
	query := `
	select
		user_id, expires_at
	from
		refresh_tokens
	where
		token = $1
	`

	userId := ""
	var expiresAt int64
	err := u.Db.QueryRow(query, token).Scan(&userId, &expiresAt)
	if err != nil {
		return "", err
	}

	if time.Now().After(time.Unix(expiresAt, 0)) {
		return "", fmt.Errorf("refresh token expired")
	}
	return userId, nil
}

func (u *UserRepo) DeleteRefreshToken(token string) error {
	_, err := u.Db.Exec(`	
		DELETE 
			FROM 
				refresh_tokens 
		WHERE 
			token = $1
		`, token)
	return err
}

func (u *UserRepo) GetUserByEmail(email string) (*pbAu.LoginResponse, error) {
	query := `
	select
		id,
		username,
		password_hash,
	from
		users
	where
		deleted_at is null
		and email = $1
	`

	user := pbAu.LoginResponse{}
	err := u.Db.QueryRow(query, email).Scan(&user.Id, &user.Username,
		&user.Password)

	return &user, err
}

func (u *UserRepo) GetUserById(userId string) (*pb.User, error) {
	query := `
	select
		username,
		email,
		password_hash,
		created_at,
		updated_at
	from
		users
	where
		deleted_at is null
		and id = $1`

	user := pb.User{Id: userId}
	updatedAt := sql.NullString{}
	err := u.Db.QueryRow(query, userId).Scan(
		&user.Username, &user.Email, &user.Password, &user.CreatedAt, &updatedAt)
	if err != nil {
		return nil, err
	}
	user.UpdatedAt = updatedAt.String

	return &user, nil
}

func (u *UserRepo) UpdateUser(user *pb.User) error {
	tx, err := u.Db.Begin()
	if err != nil {
		return err
	}
	defer tx.Commit()

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
	query += fmt.Sprintf("where deleted_at is null and id = $%d", len(params)+1)
	params = append(params, user.Id)

	_, err = tx.Exec(query, params...)
	if err != nil {
		return err
	}

	return nil
}

func (u *UserRepo) DeleteUser(id string) error {
	tx, err := u.Db.Begin()
	if err != nil {
		return err
	}
	defer tx.Commit()
	query := `
	update
	    users
	set
	    deleted_at = now()
	where 
		deleted_at is null
		and id = $1;
`

	res, err := tx.Exec(query, id)
	if err != nil {
		return err
	}
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
	query := `
	select
		full_name,
		bio,
		role,
		location,
		avatar_image,
		website
	from
		user_profiles
	where
		user_id = $1`

	profile := &pb.Profile{UserId: id}
	website := sql.NullString{}
	row := u.Db.QueryRow(query, id)
	err := row.Scan(&profile.FullName, &profile.Bio, &profile.Role, &profile.Location, &profile.AvatarImage, &website)
	if err != nil {
		return nil, err
	}
	profile.Website = website.String

	return profile, nil
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
    	website = $6
	where
		user_id = $7;
`
	res, err := u.Db.Exec(query, profile.FullName, profile.Bio, profile.Role, profile.Location,
		profile.AvatarImage, profile.Website, profile.UserId)
	if err != nil {
		return err
	}
	affectedRows, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if affectedRows == 0 {
		return fmt.Errorf("user %s not found", profile.UserId)
	}

	return nil
}
