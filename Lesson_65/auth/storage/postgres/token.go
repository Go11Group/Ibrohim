package postgres

import (
	"auth-service/models"
	"context"
	"database/sql"

	"github.com/pkg/errors"
)

type TokenRepo struct {
	DB *sql.DB
}

func NewTokenRepo(db *sql.DB) *TokenRepo {
	return &TokenRepo{DB: db}
}

func (t *TokenRepo) Store(ctx context.Context, token *models.RefreshTokenDetails) error {
	query := `
	insert into
		refresh_tokens (user_id, token, expires_at)
	values
		($1, $2, $3)
	`

	_, err := t.DB.ExecContext(ctx, query, token.UserID, token.Token, token.Expiry)
	if err != nil {
		return errors.Wrap(err, "refresh token storage failure")
	}

	return nil
}

func (t *TokenRepo) Delete(ctx context.Context, UserId string) error {
	query := `update refresh_tokens set expires_at = now() where user_id = $1`
	_, err := t.DB.ExecContext(ctx, query, UserId)
	if err != nil {
		return errors.Wrap(err, "refresh token deletion failure")
	}

	return nil
}
