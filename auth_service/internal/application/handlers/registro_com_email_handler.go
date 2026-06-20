package handlers

import (
	"context"
	"fmt"
	"strings"

	"github.com/juancarlosTI/monorepo-gestao-jur/backend/auth_service/internal/application/auth"
	"github.com/juancarlosTI/monorepo-gestao-jur/backend/auth_service/internal/application/commands"
	"github.com/juancarlosTI/monorepo-gestao-jur/backend/auth_service/internal/application/mappers"
	app_services "github.com/juancarlosTI/monorepo-gestao-jur/backend/auth_service/internal/application/service"
	"github.com/juancarlosTI/monorepo-gestao-jur/backend/auth_service/internal/dominio/identidade"
	services "github.com/juancarlosTI/monorepo-gestao-jur/backend/auth_service/internal/dominio/service"
	"go.opentelemetry.io/otel/attribute"
)

type RegistroComEmailHandler struct {
	userRepo           identidade.UserRepository
	hashSvc            services.PasswordService
	accessTokenSvc     services.AccessTokenService
	refreshTokenAppSvc *app_services.RefreshTokenAppService
}

func NewRegistroComEmailHandler(
	userRepo identidade.UserRepository,
	hashSvc services.PasswordService,
	accessTokenSvc services.AccessTokenService,
	refreshTokenAppSvc *app_services.RefreshTokenAppService,
) *RegistroComEmailHandler {
	return &RegistroComEmailHandler{
		userRepo:           userRepo,
		hashSvc:            hashSvc,
		accessTokenSvc:     accessTokenSvc,
		refreshTokenAppSvc: refreshTokenAppSvc,
	}
}

func (h *RegistroComEmailHandler) Handle(
	ctx context.Context,
	reqCtx auth.RequestContext,
	cmd commands.RegistroComEmailCommand) error {

	ctx, span := tracer.Start(ctx, "RegistroComEmailHandler.Handle")
	defer span.End()

	email_user := strings.ToLower(strings.TrimSpace(cmd.Email))
	span.SetAttributes(
		attribute.String("auth.email", email_user),
		attribute.String("auth.flow", "register"),
	)

	if !strings.Contains(email_user, "@") {
		err := fmt.Errorf("invalid email")
		span.RecordError(err)
		span.SetAttributes(attribute.String("auth.result", "invalid_email"))
		return err
	}

	user, err := h.userRepo.FindByEmail(ctx, email_user)
	if err != nil {
		span.RecordError(err)
		return err
	}

	if user != nil {
		err := fmt.Errorf("user already exists")
		span.RecordError(err)
		span.SetAttributes(attribute.String("auth.result", "user_exists"))
		return err
	}

	fmt.Println("User: ", user)

	eventCtx := mappers.EventContextFromRequest(reqCtx, cmd.AccountType)

	u, err := identidade.RegistrarUsuario(cmd.Nome, email_user, cmd.Password, h.hashSvc, eventCtx)
	if err != nil {
		span.RecordError(err)
		return err
	}

	span.SetAttributes(attribute.String("auth.user_id", u.ID()))

	if err := h.userRepo.Salvar(ctx, u); err != nil {
		span.RecordError(err)
		return err
	}

	span.SetAttributes(attribute.String("auth.result", "success"))

	// map_roles := mappers.StringToRole(u.Role())

	// // 🔐 GERAR TOKENS
	// access := &token_access.AccessToken{
	// 	UserID: u.ID(),
	// 	Roles: shared_types.Role{
	// 		Tipo: map_roles,
	// 	},
	// 	Autor: shared_types.Autor{
	// 		Tipo: shared_types.AutorHumano,
	// 		ID:   u.ID(),
	// 	},
	// 	Origem: shared_types.Origem{
	// 		Canal:   shared_types.CanalAPI,
	// 		Sistema: shared_types.Auth,
	// 	},
	// }

	// accessToken, err := h.accessTokenSvc.Generate(access)
	// if err != nil {
	// 	span.RecordError(err)
	// 	return nil, err
	// }

	// refreshTokenRaw, err := h.refreshTokenAppSvc.Create(ctx, u.ID())
	// if err != nil {
	// 	span.RecordError(err)
	// 	return nil, err
	// }

	// return &dto.TokenResponseDTO{
	// 	AccessToken:  accessToken,
	// 	RefreshToken: refreshTokenRaw,
	// }, nil
	return nil
}
