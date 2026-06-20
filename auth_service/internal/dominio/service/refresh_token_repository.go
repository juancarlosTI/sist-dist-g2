package service

import (
	"context"

	"github.com/juancarlosTI/monorepo-gestao-jur/backend/auth_service/internal/dominio/token_access"
)

type RefreshTokenRepository interface {
	Save(ctx context.Context, rt *token_access.RefreshToken) error
	FindByHash(ctx context.Context, token string) (*token_access.RefreshToken, error)
	RevokeByHash(ctx context.Context, token string) error
}
