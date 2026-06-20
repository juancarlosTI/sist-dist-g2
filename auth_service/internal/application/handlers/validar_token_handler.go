package handlers

import (
	"context"
	"fmt"

	app_services "github.com/juancarlosTI/monorepo-gestao-jur/backend/auth_service/internal/application/service"

	"github.com/juancarlosTI/monorepo-gestao-jur/backend/auth_service/internal/application/commands"
	services "github.com/juancarlosTI/monorepo-gestao-jur/backend/auth_service/internal/dominio/service"
)

type ValidateRefreshTokenHandler struct {
	accessTokenSvc     services.AccessTokenService
	refreshTokenAppSvc *app_services.RefreshTokenAppService
}

func NewValidateTokenHandler(accessTokenSvc services.AccessTokenService,
	refreshTokenAppSvc *app_services.RefreshTokenAppService) *ValidateRefreshTokenHandler {
	return &ValidateRefreshTokenHandler{
		accessTokenSvc:     accessTokenSvc,
		refreshTokenAppSvc: refreshTokenAppSvc,
	}
}

func (h *ValidateRefreshTokenHandler) Handle(ctx context.Context,
	cmd commands.RefreshTokenPayload) error {
	_, err := h.refreshTokenAppSvc.Validate(ctx, cmd.RefreshToken)
	if err != nil {
		return fmt.Errorf("invalid token")
	}

	return nil
}
