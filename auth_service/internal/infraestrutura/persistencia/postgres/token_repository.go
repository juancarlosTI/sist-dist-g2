package postgres

import (
	"context"

	"github.com/juancarlosTI/monorepo-gestao-jur/backend/auth_service/internal/dominio/token_access"
)

type RefreshTokenRepositorySQL struct {
	db DBTX
}

func NewRefreshTokenRepository(db DBTX) *RefreshTokenRepositorySQL {
	return &RefreshTokenRepositorySQL{
		db: db,
	}
}

func (r *RefreshTokenRepositorySQL) Save(ctx context.Context,
	rt *token_access.RefreshToken) error {

	query := `
	INSERT INTO refresh_tokens (value_hash, user_id, expires_at, revoked)
	VALUES ($1, $2, $3, $4)
	`

	_, err := r.db.ExecContext(ctx, query,
		rt.Value,
		rt.UserID,
		rt.ExpiresAt,
		rt.Revoked,
	)

	return err
}

func (r *RefreshTokenRepositorySQL) FindByHash(ctx context.Context, hash string) (*token_access.RefreshToken, error) {

	query := `
	SELECT value_hash, user_id, expires_at, revoked
	FROM refresh_tokens
	WHERE value_hash = $1
	AND expires_at > NOW()
	AND revoked = FALSE
	`

	row := r.db.QueryRowContext(ctx, query, hash)

	rt := &token_access.RefreshToken{}

	err := row.Scan(
		&rt.Value,
		&rt.UserID,
		&rt.ExpiresAt,
		&rt.Revoked,
	)

	if err != nil {
		return nil, err
	}

	return rt, nil
}

func (r *RefreshTokenRepositorySQL) RevokeByHash(ctx context.Context, hash string) error {

	query := `
	UPDATE refresh_tokens
	SET revoked = true, revoked_at = NOW()
	WHERE value_hash = $1
	`

	_, err := r.db.ExecContext(ctx, query, hash)
	return err
}

// func (r *RefreshTokenRepositorySQL) DeleteByHash(ctx context.Context, hash string) error {

// 	query := `
// 	DELETE FROM refresh_tokens
// 	WHERE value_hash = $1
// 	`

// 	_, err := r.db.ExecContext(ctx, query, hash)
// 	return err
// }
