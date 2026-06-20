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

	"github.com/go-chi/chi/v5"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/jackc/pgx/v5/stdlib"
	_ "github.com/juancarlosTI/monorepo-gestao-jur/backend/documentocase/cmd/api/docs"
	"github.com/juancarlosTI/monorepo-gestao-jur/backend/documentocase/interface/http/middlewares"

	"github.com/juancarlosTI/monorepo-gestao-jur/backend/documentocase/cmd/api/routes"

	application "github.com/juancarlosTI/monorepo-gestao-jur/backend/documentocase/internal/application/common"
	appHandlers "github.com/juancarlosTI/monorepo-gestao-jur/backend/documentocase/internal/application/handlers"
	"github.com/juancarlosTI/monorepo-gestao-jur/backend/documentocase/internal/application/policies"

	"github.com/juancarlosTI/monorepo-gestao-jur/backend/documentocase/internal/config"

	// Infra
	// "github.com/juancarlosTI/monorepo-gestao-jur/backend/documentocase/internal/infraestrutura/consumers"
	// infraHandlers "github.com/juancarlosTI/monorepo-gestao-jur/backend/documentocase/internal/infraestrutura/consumers/handlers"

	outbox_repo "github.com/juancarlosTI/monorepo-gestao-jur/backend/documentocase/internal/infraestrutura/outbox"
	eventostore "github.com/juancarlosTI/monorepo-gestao-jur/backend/documentocase/internal/infraestrutura/persistencia/evento_store"
	snapshotstore "github.com/juancarlosTI/monorepo-gestao-jur/backend/documentocase/internal/infraestrutura/persistencia/snapshot_store"
	documento_sql "github.com/juancarlosTI/monorepo-gestao-jur/backend/documentocase/internal/infraestrutura/persistencia/sql"
	uow_sql "github.com/juancarlosTI/monorepo-gestao-jur/backend/documentocase/internal/infraestrutura/persistencia/sql"
	storage "github.com/juancarlosTI/monorepo-gestao-jur/backend/documentocase/internal/infraestrutura/storage"
	httpSwagger "github.com/swaggo/http-swagger"
)

// @title Documento Service API
// @version 1.0
// @description Documento Service
// @host localhost
// @BasePath /api/v1/documento
// @schemes https
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func main() {

	// =====================================================================
	// CONTEXT COM SHUTDOWN
	// =====================================================================

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		s := <-sigChan
		log.Println("Shutdown recebido:", s)
		cancel()
	}()

	// =====================================================================
	// CONFIG E BANCO
	// =====================================================================

	cfg := config.Load()

	db, err := sql.Open("pgx", cfg.PostgresDSN())
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatal("Falha ao conectar no banco:", err)
	}

	storageService, err :=
		storage.NewMinIOStorage(
			ctx,
			cfg.Endpoint,
			cfg.AccessKey,
			cfg.SecretKey,
			cfg.Bucket,
			false,
		)

	if err != nil {
		log.Fatal(err)
	}

	// =====================================================================
	// POLÍTICAS, STORES E REPOSITÓRIOS
	// =====================================================================

	documentoPolicy := policies.DocumentoPolicyImpl{}
	snapshotPolicy := application.Every50EventsPolicy{}

	sqlUow := uow_sql.NewSQLUnitOfWork(db)
	eventStoreRepo := eventostore.NewEventStoreRepository(db)
	snapshotRepo := snapshotstore.NewSnapshotStoreRepository(db)
	outboxRepo := outbox_repo.NewOutboxRepositorySQL(db)
	// documentoReadRepo := documento_sql.NewDocumentoReadSQLRepositoryHandler(db)
	documentoWriteRepo := documento_sql.NewDocumentoRepositorySQLHandler(
		db,
		eventStoreRepo,
		snapshotRepo,
		snapshotPolicy,
	)

	// =====================================================================
	// QUERIES
	// =====================================================================

	// buscarDocumentoHandler := documento_query.NewBuscarDocumentoPorIDHandler(documentoReadRepo)
	// listarDocumentosDoUsuarioHandler := documento_query.NewListarDocumentosDoUsuarioPorIDHandler(documentoReadRepo)

	// =====================================================================
	// COMMAND HANDLERS (APPLICATION)
	// =====================================================================

	criarDocumentoHandler := appHandlers.NewCriarDocumentoHandler(sqlUow, outboxRepo, documentoWriteRepo, documentoPolicy)
	arquivarDocumentoHandler := appHandlers.NewArquivarDocumentoHandler(documentoWriteRepo, documentoPolicy)
	atualizarDocumentoHandler := appHandlers.NewAtualizarDocumentoHandler(documentoWriteRepo, documentoPolicy)
	associarDocumentoHandler := appHandlers.NewAssociarDocumentoAoProcessoHandler(documentoWriteRepo, documentoPolicy)
	invalidarDocumentoHandler := appHandlers.NewInvalidarDocumentoHandler(documentoWriteRepo, documentoPolicy)

	// =====================================================================
	// HTTP SERVER
	// =====================================================================

	deps := routes.Dependencies{
		Storage: storageService,
		// DocumentoCriacaoSolicitada: appConsumer,
		// BuscarProcessoHandler:     buscarDocumentoHandler,
		// ListarDocumentosDoUsuario: listarDocumentosDoUsuarioHandler,
		ArquivarDocumentoHandler:  arquivarDocumentoHandler,
		AssociarDocumentoHandler:  associarDocumentoHandler,
		AtualizarDocumentoHandler: atualizarDocumentoHandler,
		CriacaoDeDocumentoHandler: criarDocumentoHandler,
		InvalidarDocumentoHandler: invalidarDocumentoHandler,
	}

	controllers := routes.BuildControllers(deps)

	// --- HTTP SERVER ---
	r := chi.NewRouter()

	// HANDLER FINAL (único ponto de entrada)
	r.Use(middlewares.RecoveryMiddleware)
	r.Use(middlewares.TraceMiddleware)
	r.Use(middlewares.LoggingMiddleware)

	routes.RegisterRoutes(r, controllers)

	r.Get("/health", HealthHandler)
	r.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL("/api/v1/documento/swagger/doc.json"),
	))

	server := &http.Server{
		Addr:    ":" + cfg.HTTPPort,
		Handler: r,
	}

	go func() {
		log.Printf("API iniciada na porta %s", cfg.HTTPPort)

		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal("Erro no servidor HTTP:", err)
		}
	}()

	// =====================================================================
	// AGUARDA ENCERRAMENTO
	// =====================================================================

	<-ctx.Done()

	log.Println("Finalizando documento-api…")

	shutdownCtx, cancelShutdown := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancelShutdown()

	if err := server.Shutdown(shutdownCtx); err != nil {
		log.Println("Erro durante graceful shutdown:", err)
	}

	log.Println("documento-api finalizado com sucesso.")
}
