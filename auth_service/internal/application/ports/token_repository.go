package ports

import (
	"github.com/juancarlosTI/monorepo-gestao-jur/backend/auth_service/internal/dominio/token_access"
)

type RefreshTokenRepository interface {
	Save(rt *token_access.RefreshToken) error
	FindByID(token string) (*token_access.RefreshToken, error)
	RevokeByID(token string) error
}
