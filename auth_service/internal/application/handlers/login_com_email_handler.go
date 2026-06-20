package handlers

import (
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"

	"context"
	"fmt"

	"github.com/juancarlosTI/monorepo-gestao-jur/backend/auth_service/interface/http/observability"
	"github.com/juancarlosTI/monorepo-gestao-jur/backend/auth_service/internal/application/auth"
	"github.com/juancarlosTI/monorepo-gestao-jur/backend/auth_service/internal/application/commands"
	"github.com/juancarlosTI/monorepo-gestao-jur/backend/auth_service/internal/application/dto"
	"github.com/juancarlosTI/monorepo-gestao-jur/backend/auth_service/internal/application/mappers"
	app_services "github.com/juancarlosTI/monorepo-gestao-jur/backend/auth_service/internal/application/service"
	audit_service "github.com/juancarlosTI/monorepo-gestao-jur/backend/auth_service/internal/application/service"
	"github.com/juancarlosTI/monorepo-gestao-jur/backend/auth_service/internal/dominio/identidade"
	services "github.com/juancarlosTI/monorepo-gestao-jur/backend/auth_service/internal/dominio/service"
	"github.com/juancarlosTI/monorepo-gestao-jur/backend/auth_service/internal/dominio/token_access"
	shared_types "github.com/juancarlosTI/monorepo-gestao-jur/backend/shared/types"
)

type LoginComEmailHandler struct {
	userRepo           identidade.UserRepository
	passwordHashSvc    services.PasswordService
	accessTokenSvc     services.AccessTokenService
	refreshTokenAppSvc *app_services.RefreshTokenAppService
	auditSvc           *audit_service.AuditService
}

func NewLoginComEmailHandler(
	userRepo identidade.UserRepository,
	hashSvc services.PasswordService,
	accessTokenSvc services.AccessTokenService,
	refreshTokenAppSvc *app_services.RefreshTokenAppService,
	auditSvc *audit_service.AuditService,
) *LoginComEmailHandler {
	return &LoginComEmailHandler{
		userRepo:           userRepo,
		passwordHashSvc:    hashSvc,
		accessTokenSvc:     accessTokenSvc,
		refreshTokenAppSvc: refreshTokenAppSvc,
		auditSvc:           auditSvc,
	}
}

var tracer = otel.Tracer("auth-service")

func (h *LoginComEmailHandler) Handle(
	ctx context.Context,
	authCtx auth.AuthContext,
	cmd commands.LoginComEmailCommand) (*dto.TokenResponseDTO, error) {

	ctx, span := tracer.Start(ctx, "LoginComEmailHandler.Handle")
	defer span.End()

	span.SetAttributes(
		attribute.String("auth.email", cmd.Email),
		attribute.String("auth.flow", "login"),
	)

	span.SetAttributes(
		attribute.String("auth.email_hash", observability.Hash(cmd.Email)),
		attribute.String("auth.email_mask", observability.MaskEmail(cmd.Email)),
	)

	// Verificar email

	user, err := h.userRepo.FindByEmail(ctx, cmd.Email)
	if err != nil {
		span.RecordError(err)
		return nil, err
	}

	if user == nil {
		// span.RecordError(user)
		// ✅ TRACE (evento técnico)
		span.AddEvent("auth.user_not_found")
		span.SetAttributes(attribute.String("auth.result", "user_not_found"))

		h.auditSvc.Log(
			ctx,
			"login_failed",
			"",
			"", // correlationID (ou pegar do ctx depois)
			map[string]interface{}{
				"reason": "user_not_found",
				"email":  cmd.Email,
			},
		)
		return nil, fmt.Errorf("invalid credentials")
	}

	span.SetAttributes(attribute.String("auth.user_id", user.ID()))

	// Regra de dominio
	if !user.CanLoginWithPassword() {
		fmt.Println("Entrou na regra de dominio")
		span.SetAttributes(attribute.String("auth.reason", "password_login_not_allowed"))
		return nil, fmt.Errorf("invalid credentials")
	}

	if err := h.passwordHashSvc.Compare(cmd.Password, user.Password()); err != nil {
		span.RecordError(err)
		span.AddEvent("auth.invalid_password")
		span.SetAttributes(attribute.String("auth.result", "invalid_password"))

		h.auditSvc.Log(
			ctx,
			"login_failed",
			user.ID(),
			"", // correlationID (ou pegar do ctx depois)
			map[string]interface{}{
				"reason": "invalid_password",
			},
		)
		return nil, fmt.Errorf("invalid credentials")
	}

	map_roles := mappers.StringToRole(user.Role())

	access := &token_access.AccessToken{
		UserID: user.ID(),
		Roles: shared_types.Role{
			Tipo: map_roles,
		},
		Autor: shared_types.Autor{
			Tipo: authCtx.Autor.Tipo,
			ID:   user.ID(),
		},
		Origem: shared_types.Origem{
			Canal:   authCtx.Origem.Canal,
			Sistema: authCtx.Origem.Sistema,
		},
	}

	at, err := h.accessTokenSvc.Generate(access)
	if err != nil {
		span.RecordError(err)
		return nil, err
	}

	refreshTokenRaw, err := h.refreshTokenAppSvc.Create(ctx, user.ID())
	if err != nil {
		span.RecordError(err)
		return nil, err
	}

	span.SetAttributes(
		attribute.String("auth.result", "success"),
	)

	return &dto.TokenResponseDTO{
		AccessToken:  at,
		RefreshToken: refreshTokenRaw,
	}, nil
}
