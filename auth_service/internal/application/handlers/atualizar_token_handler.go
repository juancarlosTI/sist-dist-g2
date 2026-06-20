package handlers

import (
	"context"

	"github.com/juancarlosTI/monorepo-gestao-jur/backend/auth_service/internal/application/auth"
	"github.com/juancarlosTI/monorepo-gestao-jur/backend/auth_service/internal/application/commands"
	"github.com/juancarlosTI/monorepo-gestao-jur/backend/auth_service/internal/application/dto"
	"github.com/juancarlosTI/monorepo-gestao-jur/backend/auth_service/internal/application/mappers"
	app_services "github.com/juancarlosTI/monorepo-gestao-jur/backend/auth_service/internal/application/service"
	"github.com/juancarlosTI/monorepo-gestao-jur/backend/auth_service/internal/dominio/identidade"
	services "github.com/juancarlosTI/monorepo-gestao-jur/backend/auth_service/internal/dominio/service"
	"github.com/juancarlosTI/monorepo-gestao-jur/backend/auth_service/internal/dominio/token_access"
	"github.com/juancarlosTI/monorepo-gestao-jur/backend/shared/types"
)

type UpdateRefreshTokenHandler struct {
	accessSvc     services.AccessTokenService
	refreshAppSvc *app_services.RefreshTokenAppService
	userRepo      identidade.UserRepository
}

func NewUpdateRefreshTokenHandler(
	accessSvc services.AccessTokenService,
	refreshAppSvc *app_services.RefreshTokenAppService,
	userRepo identidade.UserRepository,
) *UpdateRefreshTokenHandler {
	return &UpdateRefreshTokenHandler{
		accessSvc:     accessSvc,
		refreshAppSvc: refreshAppSvc,
		userRepo:      userRepo,
	}
}

func (h *UpdateRefreshTokenHandler) Handle(
	ctx context.Context,
	authCtx auth.AuthContext,
	cmd commands.RefreshTokenPayload,
) (*dto.TokenResponseDTO, error) {

	rt, err := h.refreshAppSvc.Validate(ctx, cmd.RefreshToken)
	if err != nil {
		return nil, err
	}

	user, err := h.userRepo.FindByID(ctx, rt.UserID)
	if err != nil {
		return nil, err
	}

	newRefreshRaw, err := h.refreshAppSvc.Rotate(ctx, rt.Value)
	if err != nil {
		return nil, err
	}

	// Busca userID para preencher Roles
	map_roles := mappers.StringToRole(user.Role())

	// 🔥 usa userID do refresh token
	access := &token_access.AccessToken{
		UserID: rt.UserID,
		Roles: types.Role{
			Tipo: map_roles,
		},
		Autor: types.Autor{
			Tipo: authCtx.Autor.Tipo,
			ID:   rt.UserID,
		},
		Origem: types.Origem{
			Canal:   authCtx.Origem.Canal,
			Sistema: authCtx.Origem.Sistema,
		},
	}

	at, err := h.accessSvc.Generate(access)
	if err != nil {
		return nil, err
	}

	return &dto.TokenResponseDTO{
		AccessToken:  at,
		RefreshToken: newRefreshRaw,
	}, nil
}
