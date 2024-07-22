package postgres

import (
	pb "auth-service/genproto/user"
	"context"
	"database/sql"
	"log"
	"time"

	"github.com/lib/pq"
	"github.com/pkg/errors"
)

type UserRepo struct {
	DB *sql.DB
}

func NewUserRepo(db *sql.DB) *UserRepo {
	return &UserRepo{DB: db}
}

func (r *UserRepo) GetProfile(ctx context.Context) (*pb.Profile, error) {
	Id := ctx.Value("user_id").(string)
	query := `SELECT id, full_name, username, email, phone, image, role, created_at, updated_at
              FROM users WHERE id = $1 AND deleted_at = 0`
	row := r.DB.QueryRowContext(ctx, query, Id)

	res := &pb.Profile{}
	err := row.Scan(&res.Id, &res.FullName, &res.Username, &res.Email, &res.PhoneNumber, pq.Array(&res.Image), &res.Role, &res.CreatedAt, &res.UpdatedAt)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return res, nil
}

func (r *UserRepo) UpdateProfile(ctx context.Context, req *pb.NewData) (*pb.UpdateResp, error) {
	Id := ctx.Value("user_id").(string)
	res := &pb.UpdateResp{}
	err := r.DB.QueryRowContext(ctx, `UPDATE users SET
						full_name = $1,	username = $2, email = $3,
						phone = $4, image = $5, updated_at = $7
					WHERE id = $6 AND deleted_at = 0
					returning
						id, username, full_name, email, phone, image, updated_at
					`, req.FullName, req.Username, req.Email, req.PhoneNumber, pq.Array(req.Image), Id, time.Now()).
		Scan(&res.Id, &res.Username, &res.FullName, &res.Email, &res.PhoneNumber, pq.Array(&res.Image), &res.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (r *UserRepo) DeleteProfile(ctx context.Context) error {
	Id := ctx.Value("user_id").(string)
	result, err := r.DB.ExecContext(ctx, "UPDATE users SET deleted_at = EXTRACT(EPOCH FROM CURRENT_TIMESTAMP) WHERE deleted_at = 0 and id = $1", Id)
	if err != nil {
		return errors.New("Failure to database")
	}
	rows, err := result.RowsAffected()
	if err != nil {
		return errors.New("No rows affected")
	}
	if rows < 1 {
		return errors.New("user not found")
	}
	return nil
}

func (r *UserRepo) GetUserByEmail(ctx context.Context, email string) (string, string, string, error) {
	query := `
	SELECT
		id, username, password_hash
	FROM
		users
	WHERE
		email = $1 and deleted_at = 0`
	row := r.DB.QueryRowContext(ctx, query, email)

	var id, username, passwordHash string
	err := row.Scan(&id, &username, &passwordHash)
	if err != nil {
		return "", "", "", err
	}

	return id, username, passwordHash, nil
}

func (r *UserRepo) GetUserByID(ctx context.Context, id string) (string, string, string, error) {
	query := `
	SELECT
		username, email, password_hash
	FROM
		users
	WHERE
		id = $1 and deleted_at = 0`
	row := r.DB.QueryRowContext(ctx, query, id)

	var username, email, passwordHash string
	err := row.Scan(&username, &email, &passwordHash)
	if err != nil {
		return "", "", "", err
	}

	return username, email, passwordHash, nil
}
