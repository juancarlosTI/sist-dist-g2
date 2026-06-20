package inbox

import (
	"context"
	"database/sql"
)

type InboxRepository struct {
	db *sql.DB
}

func NewInboxRepository(db *sql.DB) *InboxRepository {
	return &InboxRepository{db: db}
}

func (r *InboxRepository) JaProcessado(
	ctx context.Context,
	eventID string,
) (bool, error) {

	var exists bool

	err := r.db.QueryRowContext(
		ctx,
		`SELECT EXISTS(
			SELECT 1 FROM eventos_processados
			WHERE evento_id = $1
		)`,
		eventID,
	).Scan(&exists)

	return exists, err
}

func (r *InboxRepository) MarcarComoProcessado(
	ctx context.Context,
	eventID string,
) error {

	_, err := r.db.ExecContext(
		ctx,
		`INSERT INTO eventos_processados (evento_id)
		 VALUES($1)
		 ON CONFLICT DO NOTHING`,
		eventID,
	)

	return err
}
