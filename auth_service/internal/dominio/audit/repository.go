package audit

import "context"

type AuditRepository interface {
	Save(ctx context.Context, event *AuditEvent) error
}
