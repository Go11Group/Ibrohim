package postgres

import (
	pb "auth-service/genproto/location"
	"context"
	"database/sql"

	"github.com/pkg/errors"
)

type LocationRepo struct {
	DB *sql.DB
}

func NewLocationRepo(db *sql.DB) *LocationRepo {
	return &LocationRepo{DB: db}
}

func (l *LocationRepo) Add(ctx context.Context, newLoc *pb.NewLocation) (*pb.NewLocationResp, error) {
	query := `
	insert into
		user_locations (user_id, address, city, state, country, postal_code)
	values
		($1, $2, $3, $4, $5, $6)
	returning
		location_id, user_id, address, city, state, country, postal_code, created_at
	`

	var loc pb.NewLocationResp
	row := l.DB.QueryRowContext(ctx, query, newLoc.UserId, newLoc.Address,
		newLoc.City, newLoc.State, newLoc.Country, newLoc.PostalCode,
	)
	err := row.Scan(&loc.Id, &loc.UserId, &loc.Address, &loc.City,
		&loc.State, &loc.Country, &loc.PostalCode, &loc.CreatedAt,
	)
	if err != nil {
		return nil, errors.Wrap(err, "location insertion failure")
	}

	return &loc, nil
}

func (l *LocationRepo) Read(ctx context.Context, id *pb.ID) (*pb.LocationDetails, error) {
	query := `
	select
		location_id, user_id, address, city, state, country, postal_code, created_at
	from
		user_locations
	where
		deleted_at = 0 and location_id = $1
	`

	loc := pb.LocationDetails{Id: id.Id}
	err := l.DB.QueryRowContext(ctx, query, id.Id).Scan(&loc.Id, &loc.UserId, &loc.Address,
		&loc.City, &loc.State, &loc.Country, &loc.PostalCode, &loc.CreatedAt,
	)
	if err != nil {
		return nil, errors.Wrap(err, "location reading failure")
	}

	return &loc, nil
}
