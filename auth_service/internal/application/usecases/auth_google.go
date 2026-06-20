package usecases

// import (
// 	"context"

// 	"github.com/juancarlosTI/monorepo-gestao-jur/backend/auth_service/internal/application/auth"
// 	"github.com/juancarlosTI/monorepo-gestao-jur/backend/auth_service/internal/application/commands"
// 	"github.com/juancarlosTI/monorepo-gestao-jur/backend/auth_service/internal/application/mappers"
// 	"github.com/juancarlosTI/monorepo-gestao-jur/backend/auth_service/internal/dominio/identidade"
// 	service "github.com/juancarlosTI/monorepo-gestao-jur/backend/auth_service/internal/dominio/service"
// 	"github.com/juancarlosTI/monorepo-gestao-jur/backend/auth_service/internal/dominio/token_access"
// 	shared_types "github.com/juancarlosTI/monorepo-gestao-jur/backend/shared/types"
// )

// type GoogleAuthUseCase struct {
// 	accessTokenService  service.AccessTokenService
// 	refreshTokenService service.RefreshTokenService
// 	userRepo            identidade.UserRepository
// }

// func NewGoogleAuthUseCase(access service.AccessTokenService, refresh service.RefreshTokenService,
// 	user identidade.UserRepository) *GoogleAuthUseCase {
// 	return &GoogleAuthUseCase{
// 		accessTokenService:  access,
// 		refreshTokenService: refresh,
// 		userRepo:            user,
// 	}
// }

// func (uc *OIDCLoginUseCase) Execute(ctx context.Context, reqCtx auth.RequestContext, cmd commands.LoginComOIDCCommand) (string, string, string, error) {

// 	p, err := uc.providers.Get(cmd.Provider)
// 	if err != nil {
// 		return "", "", "", err
// 	}

// 	tokenResp, err := p.ExchangeCode(ctx, cmd.Code, cmd.RedirectURI)
// 	if err != nil {
// 		return "", "", "", err
// 	}

// 	claims, err := p.VerifyIDToken(ctx, tokenResp.IDToken)
// 	if err != nil {
// 		return "", "", "", err
// 	}

// 	// 🔍 busca identidade externa
// 	ei, _ := uc.externalRepo.FindByProviderAndProviderID(ctx, cmd.Provider, claims.Subject)

// 	var user *identidade.User
// 	eventCtx := mappers.EventContextFromRequest(reqCtx)

// 	if ei != nil {
// 		user, err = uc.userRepo.FindByID(ctx, ei.UserID())
// 		if err != nil {
// 			return "", "", "", err
// 		}
// 	} else {
// 		// fallback por email
// 		user, _ = uc.userRepo.FindByEmail(ctx, claims.Email)

// 		if user == nil {
// 			user, err = identidade.RegistrarUsuarioOIDC(claims.Email, eventCtx)
// 			if err != nil {
// 				return "", "", "", err
// 			}
// 			if err := uc.userRepo.Salvar(ctx, user); err != nil {
// 				return "", "", "", err
// 			}
// 		}

// 		eiNew, _ := identidade.NovoExternalIdentity(user.ID(), cmd.Provider, claims.Subject)
// 		_ = uc.externalRepo.Salvar(ctx, eiNew)
// 	}

// 	map_roles, err := mappers.StringsToRoles(user.Roles())
// 	if err != nil {
// 		return "", "", "", err
// 	}

// 	access := &token_access.AccessToken{
// 		UserID: user.ID(),
// 		Roles:  map_roles,
// 		Autor: shared_types.Autor{
// 			Tipo: shared_types.AutorHumano,
// 			ID:   user.ID(),
// 		},
// 		Origem: shared_types.Origem{
// 			Canal:   shared_types.CanalAPI,
// 			Sistema: shared_types.Auth,
// 		},
// 	}

// 	// 🔐 gerar tokens
// 	at, err := uc.accessSvc.Generate(access)
// 	if err != nil {
// 		return "", "", "", err
// 	}

// 	rt, err := uc.refreshSvc.Create(ctx, user.ID())
// 	if err != nil {
// 		return "", "", "", err
// 	}

// 	return at, rt, user.ID(), nil
// }
