package service

import (
	"context"

	"github.com/juancarlosTI/monorepo-gestao-jur/backend/auth_service/internal/dominio/audit"
)

type AuditService interface {
	Log(ctx context.Context, event audit.AuditEvent) error
}
