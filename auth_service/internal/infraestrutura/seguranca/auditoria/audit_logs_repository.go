package auditoria

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"

	"github.com/juancarlosTI/monorepo-gestao-jur/backend/auth_service/internal/dominio/audit"
)

type AuditRepositorySQL struct {
	db *sql.DB
}

func NewAuditRepository(db *sql.DB) *AuditRepositorySQL {
	return &AuditRepositorySQL{
		db: db,
	}
}

func (r *AuditRepositorySQL) Save(ctx context.Context, e *audit.AuditEvent) error {
	if e == nil {
		return fmt.Errorf("audit event is nil")
	}

	metaJSON, err := json.Marshal(e.Metadata)
	if err != nil {
		return fmt.Errorf("failed to marshal metadata: %w", err)
	}

	query := `
		INSERT INTO audit_logs (
			id,
			event_type,
			user_id,
			correlation_id,
			metadata
		)
		VALUES ($1, $2, $3, $4, $5)
	`

	_, err = r.db.ExecContext(ctx, query,
		e.ID,
		e.EventType,
		e.UserID,
		e.CorrelationID,
		metaJSON,
	)
	if err != nil {
		return fmt.Errorf("failed to insert audit log: %w", err)
	}

	return nil
}
