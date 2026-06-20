package handlers

import (
	"context"

	"github.com/juancarlosTI/monorepo-gestao-jur/backend/auth_service/internal/application/auth"
	"github.com/juancarlosTI/monorepo-gestao-jur/backend/auth_service/internal/application/commands"
	app_services "github.com/juancarlosTI/monorepo-gestao-jur/backend/auth_service/internal/application/service"
	"github.com/juancarlosTI/monorepo-gestao-jur/backend/auth_service/internal/dominio/audit"
	"github.com/juancarlosTI/monorepo-gestao-jur/backend/auth_service/internal/dominio/identidade"
	services "github.com/juancarlosTI/monorepo-gestao-jur/backend/auth_service/internal/dominio/service"
	"go.opentelemetry.io/otel/attribute"
)

type LogoutHandler struct {
	passwordHashSvc services.PasswordService
	// accessTokenSvc  services.
	userRepo           identidade.UserRepository
	auditSvc           services.AuditService
	refreshTokenAppSvc *app_services.RefreshTokenAppService
}

func NewLogoutHandler(
	refreshTokenAppSvc *app_services.RefreshTokenAppService,
) *LogoutHandler {
	return &LogoutHandler{
		refreshTokenAppSvc: refreshTokenAppSvc,
	}
}

func (h *LogoutHandler) Handle(ctx context.Context, authCtx auth.AuthContext,
	cmd commands.RefreshTokenPayload) error {

	ctx, span := tracer.Start(ctx, "LogoutHandler.Handle")
	defer span.End()

	// span.SetAttributes(
	// 	attribute.String("auth.email", cmd.Email),
	// 	attribute.String("auth.flow", "logout"),
	// )

	// span.SetAttributes(
	// 	attribute.String("auth.email_hash", observability.Hash(cmd.Email)),
	// 	attribute.String("auth.email_mask", observability.MaskEmail(cmd.Email)),
	// )

	// // Verificar email

	// user, err := h.userRepo.FindByEmail(ctx, cmd.Email)
	// if err != nil {
	// 	span.RecordError(err)
	// 	// ✅ TRACE (evento técnico)
	// 	span.AddEvent("auth.user_not_found")
	// 	span.SetAttributes(attribute.String("auth.result", "user_not_found"))

	// 	h.auditSvc.Log(ctx, audit.AuditEvent{
	// 		EventType: "logout_failed",
	// 		Metadata: map[string]interface{}{
	// 			"reason":     "user_not_found",
	// 			"email_hash": observability.Hash(cmd.Email),
	// 		},
	// 	})
	// 	return fmt.Errorf("invalid credentials")
	// }
	// span.SetAttributes(attribute.String("auth.user_id", user.ID()))

	if err := h.refreshTokenAppSvc.Revoke(ctx, cmd.RefreshToken); err != nil {
		span.RecordError(err)
		span.AddEvent("auth.invalid_password")
		span.SetAttributes(attribute.String("auth.result", "invalid_password"))

		h.auditSvc.Log(ctx, audit.AuditEvent{
			EventType: "logout_failed",
			UserID:    "",
			Metadata: map[string]interface{}{
				"reason": "invalid refresh_token",
			},
		})
		return nil
	}

	return nil
}
