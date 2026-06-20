package routes

import (
	"net/http"

	"github.com/juancarlosTI/monorepo-gestao-jur/backend/auth_service/interface/http/controllers"
	"github.com/juancarlosTI/monorepo-gestao-jur/backend/auth_service/internal/application/handlers"
	"github.com/juancarlosTI/monorepo-gestao-jur/backend/auth_service/internal/application/usecases"
	"github.com/juancarlosTI/monorepo-gestao-jur/backend/auth_service/internal/config"
	jwt_service "github.com/juancarlosTI/monorepo-gestao-jur/backend/auth_service/internal/infraestrutura/seguranca/auth-token-jwt"
)

type APIControllers struct {
	AuthController *controllers.AuthController
	Config         *config.Config
}

type Dependencies struct {
	RegistroComEmailHandler *handlers.RegistroComEmailHandler

	LoginComEmailHandler  *handlers.LoginComEmailHandler
	LoginComGoogleHandler *usecases.OIDCLoginUseCase
	LogoutHandler         *handlers.LogoutHandler

	MeHandler *handlers.MeHandler

	ValidarRefreshTokenHandler *handlers.ValidateRefreshTokenHandler
	AtualizarRefreshToken      *handlers.UpdateRefreshTokenHandler
	RevogarRefreshTokenHandler *handlers.RevokeTokenHandler
}

func BuildControllers(deps Dependencies, config *config.Config) APIControllers {

	return APIControllers{
		AuthController: BuildAuthControllers(deps, config),
	}
}

func RegisterRoutes(mux *http.ServeMux, c APIControllers, accessSvc jwt_service.JWTAccessTokenService) {

	RegisterAuthRoutes(mux, c, accessSvc)
}
