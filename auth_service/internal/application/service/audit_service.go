package service

import (
	"context"

	"github.com/juancarlosTI/monorepo-gestao-jur/backend/auth_service/internal/dominio/audit"
)

type AuditService struct {
	repo audit.AuditRepository
}

func NewAuditService(repo audit.AuditRepository) *AuditService {
	return &AuditService{
		repo: repo,
	}
}

func (s *AuditService) Log(
	ctx context.Context,
	eventType string,
	userID string,
	correlationID string,
	metadata map[string]interface{},
) error {

	event, err := audit.NewAuditEvent(
		eventType,
		userID,
		correlationID,
		metadata,
	)
	if err != nil {
		return err
	}

	return s.repo.Save(ctx, event)
}
