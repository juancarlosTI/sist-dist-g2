package service

import (
	"context"
	"errors"
	"time"

	"github.com/juancarlosTI/monorepo-gestao-jur/backend/auth_service/internal/dominio/token_access"
)

type RefreshTokenService interface {
	Create(ctx context.Context, userID string) (string, error)
	Rotate(ctx context.Context, rt string) (string, error)
	Revoke(ctx context.Context, rt string) error
	Validate(ctx context.Context, rt string) (*token_access.RefreshToken, error)
}

type refreshTokenService struct {
	repo      RefreshTokenRepository
	hasher    token_access.TokenHasher
	generator token_access.TokenGenerator
}

func NewRefreshTokenService(
	repo RefreshTokenRepository,
	hasher token_access.TokenHasher,
	generator token_access.TokenGenerator,
) *refreshTokenService {
	return &refreshTokenService{
		repo:      repo,
		hasher:    hasher,
		generator: generator,
	}
}

func (s *refreshTokenService) Create(ctx context.Context,
	userID string) (string, error) {

	value, err := s.generator.Generate()
	if err != nil {
		return "", err
	}

	hash := s.hasher.Hash(value)

	rt := &token_access.RefreshToken{
		Value:     hash,
		UserID:    userID,
		ExpiresAt: time.Now().Add(7 * 24 * time.Hour),
		Revoked:   false,
		CreatedAt: time.Now(),
	}

	if err := s.repo.Save(ctx, rt); err != nil {
		return "", err
	}

	return value, nil
}

func (s *refreshTokenService) Validate(ctx context.Context,
	value string) (*token_access.RefreshToken, error) {

	hash := s.hasher.Hash(value)

	rt, err := s.repo.FindByHash(ctx, hash)
	if err != nil {
		return nil, err
	}

	if rt.IsRevoked() {
		return nil, errors.New("invalid token")
	}

	if rt.IsExpired(time.Now()) {
		return nil, errors.New("expired token")
	}

	return rt, nil
}

func (s *refreshTokenService) Rotate(ctx context.Context,
	rt string) (string, error) {

	// 🔥 Revoga o antigo
	if err := s.repo.RevokeByHash(ctx, rt); err != nil {
		return "", err
	}

	// 🔥 Gera novo
	return s.Create(ctx, "")
}

func (s *refreshTokenService) Revoke(ctx context.Context, value string) error {

	hash := s.hasher.Hash(value)

	return s.repo.RevokeByHash(ctx, hash)
}
