package routes

import (
	"net/http"

	"github.com/juancarlosTI/monorepo-gestao-jur/backend/auth_service/internal/config"
	jwt_service "github.com/juancarlosTI/monorepo-gestao-jur/backend/auth_service/internal/infraestrutura/seguranca/auth-token-jwt"

	controller "github.com/juancarlosTI/monorepo-gestao-jur/backend/auth_service/interface/http/controllers"
	"github.com/juancarlosTI/monorepo-gestao-jur/backend/auth_service/interface/http/middlewares"
)

func BuildAuthControllers(deps Dependencies, config *config.Config) *controller.AuthController {

	return controller.NewAuthController(
		deps.RegistroComEmailHandler,
		deps.LoginComEmailHandler,
		deps.LoginComGoogleHandler,
		deps.MeHandler,
		deps.LogoutHandler,
		deps.ValidarRefreshTokenHandler,
		deps.AtualizarRefreshToken,
		deps.RevogarRefreshTokenHandler,
		config,
	)
}

func RegisterAuthRoutes(mux *http.ServeMux, c APIControllers, accessSvc jwt_service.JWTAccessTokenService) {

	mux.Handle(
		"POST /auth/register",
		middlewares.Chain(
			http.HandlerFunc(c.AuthController.RegistroComEmail),
		),
	)

	mux.Handle(
		"POST /auth/login",
		middlewares.Chain(
			http.HandlerFunc(c.AuthController.LoginComEmail),
		),
	)

	mux.Handle(
		"GET /auth/logout",
		middlewares.Chain(
			http.HandlerFunc(c.AuthController.Logout),
			middlewares.TokenExtractorMiddleware,
			middlewares.TokenValidateMiddleware(accessSvc),
			middlewares.AuthContextMiddleware,
		),
	)

	mux.Handle(
		"POST /auth/token/refresh",
		middlewares.Chain(
			http.HandlerFunc(c.AuthController.Refresh),
			middlewares.TokenExtractorMiddleware,
			middlewares.TokenValidateMiddleware(accessSvc),
			middlewares.AuthContextMiddleware,
		),
	)

	mux.Handle(
		"POST /auth/token/revoke",
		middlewares.Chain(
			http.HandlerFunc(c.AuthController.Revoke),
			middlewares.TokenExtractorMiddleware,
			middlewares.TokenValidateMiddleware(accessSvc),
			middlewares.AuthContextMiddleware,
		),
	)

	mux.Handle(
		"POST /auth/google/login",
		middlewares.Chain(
			http.HandlerFunc(c.AuthController.GoogleLogin),
		),
	)

	mux.Handle(
		"GET /auth/me",
		middlewares.Chain(
			http.HandlerFunc(c.AuthController.Me),

			middlewares.TokenExtractorMiddleware,
			middlewares.TokenValidateMiddleware(accessSvc),
			middlewares.AuthContextMiddleware,
		),
	)
}
