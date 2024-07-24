package postgres

import (
	pb "auth-service/genproto/admin"
	"context"
	"database/sql"
	"strconv"

	"github.com/lib/pq"
	"github.com/pkg/errors"
)

type AdminRepo struct {
	DB *sql.DB
}

func NewAdminRepo(db *sql.DB) *AdminRepo {
	return &AdminRepo{DB: db}
}

func (r *AdminRepo) Add(ctx context.Context, d *pb.NewUser) (*pb.NewUserResp, error) {
	query := `
	insert into
		users (username, email, password_hash, full_name, role)
	values
		($1, $2, $3, $4, $5)
	returning
		id, username, email, full_name, role, created_at
	`

	var u pb.NewUserResp
	row := r.DB.QueryRowContext(ctx, query, d.Username, d.Email, d.Password, d.FullName, d.Role)
	err := row.Scan(&u.Id, &u.Username, &u.Email, &u.FullName, &u.Role, &u.CreatedAt)
	if err != nil {
		return nil, errors.Wrap(err, "user insertion failure")
	}

	return &u, nil
}

func (r *AdminRepo) Read(ctx context.Context, id *pb.ID) (*pb.UserInfo, error) {
	query := `
	select
		username, email, password_hash, full_name,
		phone, image, role, created_at, updated_at
	from
		users
	where
		deleted_at = 0 and id = $1
	`

	u := pb.UserInfo{Id: id.Id}
	err := r.DB.QueryRowContext(ctx, query, id.Id).Scan(&u.Username, &u.Email, &u.Password,
		&u.FullName, &u.PhoneNumber, pq.Array(&u.Image), &u.Role, &u.CreatedAt, &u.UpdatedAt,
	)
	if err != nil {
		return nil, errors.Wrap(err, "user reading failure")
	}

	return &u, nil
}

func (r *AdminRepo) Update(ctx context.Context, data *pb.NewData) (*pb.NewDataResp, error) {
	query := `
	update
		users
	set
		username = $1, email = $2, password_hash = $3, full_name = $4,
		phone = $5, image = $6, role = $7, updated_at = NOW()
	where
		id = $8 and deleted_at = 0
	returning
		id, username, email, password_hash, full_name, phone, image, role, updated_at
	`

	var d pb.NewDataResp
	row := r.DB.QueryRowContext(ctx, query, data.Username, data.Email, data.Password,
		data.FullName, data.PhoneNumber, pq.Array(data.Image), data.Role, data.Id)
	err := row.Scan(&d.Id, &d.Username, &d.Email, &d.Password, &d.FullName,
		&d.PhoneNumber, pq.Array(&d.Image), &d.Role, &d.UpdatedAt)
	if err != nil {
		return nil, errors.Wrap(err, "user update failure")
	}

	return &d, nil
}

func (r *AdminRepo) Delete(ctx context.Context, id *pb.ID) error {
	query := `
	update
		users
	set
		deleted_at = EXTRACT(EPOCH FROM NOW())
	where
		id = $1 and deleted_at = 0 and role <> 'admin'
	`
	rows, err := r.DB.ExecContext(ctx, query, id.Id)
	if err != nil {
		return errors.Wrap(err, "user deletion failure")
	}

	rowsAffected, err := rows.RowsAffected()
	if err != nil {
		return errors.Wrap(err, "rows affected failure")
	}
	if rowsAffected < 1 {
		return errors.New("user not found")
	}

	return nil
}

func (r *AdminRepo) FetchUsers(ctx context.Context, f *pb.Filter) (*pb.Users, error) {
	query := `
	select
		id, username, email, password_hash, full_name, role, ul.country, u.created_at, u.updated_at
	from
		users u
	join
		user_locations ul
	on
		ul.user_id = u.id
	where
		u.deleted_at = 0
	`

	if f.FullName != "" {
		query += " and full_name ILIKE '%" + f.FullName + "%'"
	}
	if f.Location != "" {
		query += " and id in (select user_id from user_locations where country ILIKE '%" + f.Location + "%')"
	}
	if f.Role != "" {
		query += " and role = '" + f.Role + "'"
	}
	if f.Limit > 0 {
		query += " limit " + strconv.Itoa(int(f.Limit))
	}
	if f.Page > 0 {
		offset := (int(f.Page) - 1) * int(f.Limit)
		query += " offset " + strconv.Itoa(offset)
	}

	rows, err := r.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, errors.Wrap(err, "users fetching failure")
	}
	defer rows.Close()

	var users []*pb.UserDetails
	for rows.Next() {
		var u pb.UserDetails
		err = rows.Scan(&u.Id, &u.Username, &u.Email, &u.Password, &u.FullName,
			&u.Role, &u.Country, &u.CreatedAt, &u.UpdatedAt)
		if err != nil {
			return nil, errors.Wrap(err, "user scan failure")
		}
		users = append(users, &u)
	}

	return &pb.Users{
		Users: users,
		Page:  f.Page,
		Limit: f.Limit,
	}, nil
}
