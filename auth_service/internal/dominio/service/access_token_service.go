package service

import (
	"context"

	"github.com/juancarlosTI/monorepo-gestao-jur/backend/auth_service/internal/dominio/token_access"
)

type AccessTokenService interface {
	Generate(at *token_access.AccessToken) (string, error)
	Validate(ctx context.Context, token string) (*token_access.AccessToken, error)
}
