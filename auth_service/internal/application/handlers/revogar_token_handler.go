package handlers

import (
	"context"

	"github.com/juancarlosTI/monorepo-gestao-jur/backend/auth_service/internal/application/auth"
	"github.com/juancarlosTI/monorepo-gestao-jur/backend/auth_service/internal/application/commands"
	app_services "github.com/juancarlosTI/monorepo-gestao-jur/backend/auth_service/internal/application/service"
)

type RevokeTokenHandler struct {
	// userRepo ports.UserRepository
	refreshTokenAppSvc *app_services.RefreshTokenAppService
}

func NewRevokeTokenHandler(
	// userRepo ports.UserRepository,
	refreshTokenAppSvc *app_services.RefreshTokenAppService,
) *RevokeTokenHandler {
	return &RevokeTokenHandler{
		// userRepo: userRepo,
		refreshTokenAppSvc: refreshTokenAppSvc,
	}
}

func (h *RevokeTokenHandler) Handle(ctx context.Context,
	authCtx auth.AuthContext,
	cmd commands.RefreshTokenPayload) error {

	if err := h.refreshTokenAppSvc.Revoke(ctx, cmd.RefreshToken); err != nil {
		return err
	}

	return nil

}
