package audit

import (
	"errors"

	"github.com/google/uuid"
)

type AuditEvent struct {
	ID            string
	EventType     string
	UserID        string
	CorrelationID string
	Metadata      map[string]interface{}
}

func NewAuditEvent(
	eventType string,
	userID string,
	correlationID string,
	metadata map[string]interface{},
) (*AuditEvent, error) {

	if eventType == "" {
		return nil, errors.New("eventType is required")
	}

	if correlationID == "" {
		return nil, errors.New("correlationID is required")
	}

	if metadata == nil {
		metadata = make(map[string]interface{})
	}

	return &AuditEvent{
		ID:            uuid.NewString(),
		EventType:     eventType,
		UserID:        userID,
		CorrelationID: correlationID,
		Metadata:      metadata,
	}, nil
}
