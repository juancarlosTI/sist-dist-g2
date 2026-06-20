package main

import (
	"context"
	"database/sql"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/jackc/pgx/v5/stdlib"
	_ "github.com/juancarlosTI/monorepo-gestao-jur/backend/auth_service/cmd/api/docs"
	"github.com/juancarlosTI/monorepo-gestao-jur/backend/auth_service/cmd/api/routes"

	"github.com/juancarlosTI/monorepo-gestao-jur/backend/auth_service/interface/http/middlewares"
	"github.com/juancarlosTI/monorepo-gestao-jur/backend/auth_service/internal/application/handlers"
	app_services "github.com/juancarlosTI/monorepo-gestao-jur/backend/auth_service/internal/application/service"
	audit_service "github.com/juancarlosTI/monorepo-gestao-jur/backend/auth_service/internal/application/service"
	"github.com/juancarlosTI/monorepo-gestao-jur/backend/auth_service/internal/application/usecases"
	"github.com/juancarlosTI/monorepo-gestao-jur/backend/auth_service/internal/config"
	"github.com/juancarlosTI/monorepo-gestao-jur/backend/auth_service/internal/dominio/service"
	oidc "github.com/juancarlosTI/monorepo-gestao-jur/backend/auth_service/internal/federation/oidc"
	"github.com/juancarlosTI/monorepo-gestao-jur/backend/auth_service/internal/infraestrutura/persistencia/postgres"
	"github.com/juancarlosTI/monorepo-gestao-jur/backend/auth_service/internal/infraestrutura/seguranca/auditoria"
	jwt_service "github.com/juancarlosTI/monorepo-gestao-jur/backend/auth_service/internal/infraestrutura/seguranca/auth-token-jwt"
	"github.com/juancarlosTI/monorepo-gestao-jur/backend/auth_service/internal/infraestrutura/seguranca/hash"
	"github.com/juancarlosTI/monorepo-gestao-jur/backend/auth_service/internal/infraestrutura/seguranca/token_generator"
	httpSwagger "github.com/swaggo/http-swagger"
)

// @title Auth Service API
// @version 1.0
// @description Auth Service
// @host localhost
// @BasePath /api/v1/auth-service
// @schemes https

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func main() {

	// --- CONTEXTO PRINCIPAL ---
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Captura sinais
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		s := <-sigChan
		log.Println("Shutdown recebido:", s)
		cancel()
	}()

	// --- CONFIG E DB ---
	cfg := config.Load()

	db, err := sql.Open("pgx", cfg.PostgresDSN())
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatal("Erro conectando no banco:", err)
	}

	userRepo := postgres.NewUserRepositorySQL(db)
	tokenRepo := postgres.NewRefreshTokenRepository(db)
	auditRepo := auditoria.NewAuditRepository(db)
	// External identities repo
	externalRepo := postgres.NewExternalIdentityRepositorySQL(db)

	// Services
	auditSvc := audit_service.NewAuditService(auditRepo)
	hashSvc := hash.NewSHA256TokenHasher()
	tokenGeneratorSvc := token_generator.NewSecureTokenGenerator(32)
	refreshTokenSvc := service.NewRefreshTokenService(tokenRepo, hashSvc, tokenGeneratorSvc)
	refreshTokenAppSvc := app_services.NewRefreshTokenAppService(refreshTokenSvc, hashSvc)
	accessSvc := jwt_service.NewJWTAccessTokenService(cfg.JWTSecret, time.Duration(cfg.AccessTokenTTL))
	pwHashSvc := hash.NewArgon2Service()

	// Handlers
	registroComEmailHandler := handlers.NewRegistroComEmailHandler(userRepo, pwHashSvc, accessSvc, refreshTokenAppSvc)
	loginComEmailHandler := handlers.NewLoginComEmailHandler(userRepo, pwHashSvc, accessSvc, refreshTokenAppSvc, auditSvc)
	meHandler := handlers.NewMeHandler(userRepo)
	// OIDC providers (Google)
	googleCfg := oidc.GoogleProviderConfig{
		ClientID:      cfg.OIDCGoogleClientID,
		ClientSecret:  cfg.OIDCGoogleClientSecret,
		Issuer:        cfg.OIDCGoogleIssuer,
		TokenEndpoint: cfg.OIDCGoogleTokenURL,
		JWKSURI:       cfg.OIDCGoogleJWKSURL,
		HTTPTimeout:   time.Duration(cfg.HTTPTimeoutSeconds) * time.Second,
		JWKSCacheTTL:  10 * time.Minute,
	}

	googleProvider, err := oidc.NewGoogleProvider(googleCfg)
	if err != nil {
		log.Println("warning: failed to create google oidc provider:", err)
	}

	providers := oidc.NewProviderRegistry()
	if googleProvider != nil {
		providers.Register(googleProvider)
	}

	loginComGoogleHandler := usecases.NewOIDCLoginUseCase(providers, accessSvc, refreshTokenSvc, userRepo, externalRepo)

	// OIDC login usecase
	logoutHandler := handlers.NewLogoutHandler(refreshTokenAppSvc)
	validarHandler := handlers.NewValidateTokenHandler(accessSvc, refreshTokenAppSvc)
	atualizarTokenHandler := handlers.NewUpdateRefreshTokenHandler(accessSvc, refreshTokenAppSvc, userRepo)
	revogarTokenHandler := handlers.NewRevokeTokenHandler(refreshTokenAppSvc)

	deps := routes.Dependencies{
		RegistroComEmailHandler: registroComEmailHandler,
		LoginComEmailHandler:    loginComEmailHandler,
		LoginComGoogleHandler:   loginComGoogleHandler,
		LogoutHandler:           logoutHandler,

		MeHandler: meHandler,

		ValidarRefreshTokenHandler: validarHandler,
		AtualizarRefreshToken:      atualizarTokenHandler,
		RevogarRefreshTokenHandler: revogarTokenHandler,
	}
	controllers := routes.BuildControllers(deps, cfg)

	// --- HTTP SERVER ---
	mux := http.NewServeMux()

	// Rotas principais
	routes.RegisterRoutes(mux, controllers, *accessSvc)

	// Rota interna (nginx)
	mux.Handle(
		"GET /auth/validate",
		middlewares.Chain(
			http.HandlerFunc(controllers.AuthController.Validate),

			middlewares.TokenExtractorMiddleware,
			middlewares.TokenValidateMiddleware(*accessSvc),
			middlewares.AuthContextMiddleware,
		),
	)

	// Health
	mux.HandleFunc("/health", HealthHandler)

	// Swagger
	mux.Handle(
		"/swagger/",
		httpSwagger.Handler(
			httpSwagger.URL("/api/v1/auth-service/swagger/doc.json"),
		),
	)

	// HANDLER FINAL (único ponto de entrada)
	var handler http.Handler = middlewares.Chain(
		mux,
		middlewares.RecoveryMiddleware,
		middlewares.RequestContextMiddleware,
		middlewares.LoggingMiddleware,
		middlewares.TracingMiddleware,
	)

	server := &http.Server{
		Addr:    ":" + cfg.HTTPPort,
		Handler: handler,
	}

	go func() {
		log.Printf("API iniciada na porta %s", cfg.HTTPPort)

		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal("Erro iniciando servidor HTTP:", err)
		}
	}()

	// Espera cancelamento
	<-ctx.Done()
	log.Println("Encerrando API…")

	// Timeout para shutdown suave
	shutdownCtx, cancelShutdown := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancelShutdown()

	if err := server.Shutdown(shutdownCtx); err != nil {
		log.Println("Erro no shutdown HTTP:", err)
	}

	log.Println("API finalizada com sucesso.")

}
