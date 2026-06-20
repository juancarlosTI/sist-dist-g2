package postgres

import (
	"context"
	"database/sql"
	"time"

	dominio "github.com/juancarlosTI/monorepo-gestao-jur/backend/auth_service/internal/dominio/identidade"
)

type ExternalIdentityRepositorySQL struct {
	db *sql.DB
}

func NewExternalIdentityRepositorySQL(db *sql.DB) *ExternalIdentityRepositorySQL {
	return &ExternalIdentityRepositorySQL{db: db}
}

func (r *ExternalIdentityRepositorySQL) Salvar(ctx context.Context, ei *dominio.ExternalIdentity) error {
	query := `
    INSERT INTO external_identities (id, user_id, provider, provider_user_id, created_at)
    VALUES ($1, $2, $3, $4, $5)
    ON CONFLICT (provider, provider_user_id) DO UPDATE SET
        user_id = EXCLUDED.user_id
    `

	_, err := r.db.ExecContext(ctx, query,
		ei.ID(), ei.UserID(), ei.Provider(), ei.ProviderUserID(), ei.CreatedAt(),
	)

	return err
}

func (r *ExternalIdentityRepositorySQL) FindByProviderAndProviderID(ctx context.Context, provider string, providerUserID string) (*dominio.ExternalIdentity, error) {
	query := `
    SELECT id, user_id, provider, provider_user_id, created_at
    FROM external_identities
    WHERE provider = $1 AND provider_user_id = $2
    `

	row := r.db.QueryRowContext(ctx, query, provider, providerUserID)

	var id string
	var userID string
	var prov string
	var provUserID string
	var createdAt time.Time

	if err := row.Scan(&id, &userID, &prov, &provUserID, &createdAt); err != nil {
		return nil, err
	}

	return dominio.ReconstruirExternalIdentity(id, userID, prov, provUserID, createdAt)
}

func (r *ExternalIdentityRepositorySQL) FindByUserID(ctx context.Context, userID string) ([]*dominio.ExternalIdentity, error) {
	query := `
    SELECT id, user_id, provider, provider_user_id, created_at
    FROM external_identities
    WHERE user_id = $1
    `

	rows, err := r.db.QueryContext(ctx, query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var results []*dominio.ExternalIdentity
	for rows.Next() {
		var id string
		var uid string
		var prov string
		var provUserID string
		var createdAt time.Time
		if err := rows.Scan(&id, &uid, &prov, &provUserID, &createdAt); err != nil {
			return nil, err
		}
		ei, err := dominio.ReconstruirExternalIdentity(id, uid, prov, provUserID, createdAt)
		if err != nil {
			return nil, err
		}
		results = append(results, ei)
	}

	return results, nil
}

func (r *ExternalIdentityRepositorySQL) Delete(ctx context.Context, provider string, providerUserID string) error {
	query := `DELETE FROM external_identities WHERE provider = $1 AND provider_user_id = $2`
	_, err := r.db.ExecContext(ctx, query, provider, providerUserID)
	return err
}
