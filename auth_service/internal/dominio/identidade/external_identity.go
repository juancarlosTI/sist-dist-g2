package identidade

import (
	"time"

	"github.com/google/uuid"
)

type ExternalIdentity struct {
	id             string
	userID         string
	provider       string
	providerUserID string
	createdAt      time.Time
}

func NovoExternalIdentity(userID string, provider string, providerUserID string) (*ExternalIdentity, error) {
	if provider == "" || providerUserID == "" || userID == "" {
		return nil, nil
	}

	id := uuid.NewString()

	return &ExternalIdentity{
		id:             id,
		userID:         userID,
		provider:       provider,
		providerUserID: providerUserID,
		createdAt:      time.Now().UTC(),
	}, nil
}

func ReconstruirExternalIdentity(idStr string, userIDStr string, provider string, providerUserID string, createdAt time.Time) (*ExternalIdentity, error) {
	return &ExternalIdentity{
		id:             idStr,
		userID:         userIDStr,
		provider:       provider,
		providerUserID: providerUserID,
		createdAt:      createdAt,
	}, nil
}

func (e *ExternalIdentity) ID() string             { return e.id }
func (e *ExternalIdentity) UserID() string         { return e.userID }
func (e *ExternalIdentity) Provider() string       { return e.provider }
func (e *ExternalIdentity) ProviderUserID() string { return e.providerUserID }
func (e *ExternalIdentity) CreatedAt() time.Time   { return e.createdAt }
