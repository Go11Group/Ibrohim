package postgres

import (
	"auth-service/models"
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

func (l *LocationRepo) Add(ctx context.Context, newLoc *models.NewLocation) (*models.LocationDetails, error) {
	query := `
	insert into
		user_locations (user_id, address, city, state, country, postal_code)
	values
		($1, $2, $3, $4, $5, $6)
	returning
		location_id, user_id, address, city, state, country, postal_code, created_at
	`

	var loc models.LocationDetails
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

func (l *LocationRepo) Read(ctx context.Context, id string) (*models.LocationDetails, error) {
	query := `
	select
		location_id, user_id, address, city, state, country, postal_code, created_at
	from
		user_locations
	where
		deleted_at = 0 and user_id = $1
	`

	loc := models.LocationDetails{Id: id}
	err := l.DB.QueryRowContext(ctx, query, id).Scan(&loc.Id, &loc.UserId, &loc.Address,
		&loc.City, &loc.State, &loc.Country, &loc.PostalCode, &loc.CreatedAt,
	)
	if err != nil {
		return nil, errors.Wrap(err, "location reading failure")
	}

	return &loc, nil
}

func (l *LocationRepo) Update(ctx context.Context, newLoc *models.NewLocation) (*models.UpdateResp, error) {
	query := `
	UPDATE user_locations SET
		address = $1, city = $2, state = $3, country = $4, postal_code = $5, updated_at = CURRENT_TIMESTAMP
	WHERE user_id = $6 AND deleted_at = 0
	RETURNING location_id, user_id, address, city, state, country, postal_code, updated_at
	`

	var loc models.UpdateResp
	err := l.DB.QueryRowContext(ctx, query, newLoc.Address, newLoc.City,
		newLoc.State, newLoc.Country, newLoc.PostalCode, newLoc.UserId,
	).Scan(&loc.Id, &loc.UserId, &loc.Address, &loc.City, &loc.State,
		&loc.Country, &loc.PostalCode, &loc.UpdatedAt,
	)
	if err != nil {
		return nil, errors.Wrap(err, "location update failure")
	}

	return &loc, nil
}

func (l *LocationRepo) Delete(ctx context.Context, id string) error {
	query := `
	UPDATE user_locations SET deleted_at = EXTRACT(EPOCH FROM CURRENT_TIMESTAMP)
	WHERE user_id = $1 AND deleted_at = 0
	`

	result, err := l.DB.ExecContext(ctx, query, id)
	if err != nil {
		return errors.Wrap(err, "location deletion failure")
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return errors.Wrap(err, "failed to get rows affected")
	}
	if rows < 1 {
		return errors.New("location not found or already deleted")
	}

	return nil
}
