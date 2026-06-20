package service

import (
	"context"

	"github.com/juancarlosTI/monorepo-gestao-jur/backend/auth_service/internal/dominio/service"
	"github.com/juancarlosTI/monorepo-gestao-jur/backend/auth_service/internal/dominio/token_access"
)

type RefreshTokenAppService struct {
	token_hash token_access.TokenHasher
	service    service.RefreshTokenService
}

func NewRefreshTokenAppService(
	refreshSvc service.RefreshTokenService,
	tokenHasher token_access.TokenHasher,
) *RefreshTokenAppService {
	return &RefreshTokenAppService{
		service:    refreshSvc,
		token_hash: tokenHasher,
	}
}

func (s *RefreshTokenAppService) Create(ctx context.Context,
	user_id string) (string, error) {

	refresh_token, err := s.service.Create(ctx, user_id)
	if err != nil {
		return "", err
	}

	return refresh_token, nil
}

func (s *RefreshTokenAppService) Rotate(ctx context.Context,
	rt string) (string, error) {

	rotate_token, err := s.service.Rotate(ctx, rt)
	if err != nil {
		return "", err
	}

	return rotate_token, nil
}

func (s *RefreshTokenAppService) Revoke(
	ctx context.Context,
	rt string,
) error {

	hash := s.token_hash.Hash(rt)

	return s.service.Revoke(ctx, hash)
}

func (s *RefreshTokenAppService) Validate(
	ctx context.Context,
	value string,
) (*token_access.RefreshToken, error) {

	// 1. Hash do token (NUNCA usar token puro no banco)
	hash := s.token_hash.Hash(value)

	// 2. Validar no domínio
	rt, err := s.service.Validate(ctx, hash)
	if err != nil {
		return nil, err
	}

	// 3. Retornar resposta (não precisa devolver token de novo)
	return rt, nil
}
