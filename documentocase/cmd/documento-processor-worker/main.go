package main

import (
	"context"
	"database/sql"
	"log"
	"os"
	"os/signal"
	"syscall"

	_ "github.com/jackc/pgx/v5/stdlib"

	"github.com/juancarlosTI/monorepo-gestao-jur/backend/documentocase/internal/config"

	"github.com/juancarlosTI/monorepo-gestao-jur/backend/documentocase/internal/application/services"

	consumer_handlers "github.com/juancarlosTI/monorepo-gestao-jur/backend/documentocase/internal/infraestrutura/consumers/handlers"
	"github.com/juancarlosTI/monorepo-gestao-jur/backend/documentocase/internal/infraestrutura/inbox"
	"github.com/juancarlosTI/monorepo-gestao-jur/backend/documentocase/internal/infraestrutura/ocr"
	"github.com/juancarlosTI/monorepo-gestao-jur/backend/documentocase/internal/infraestrutura/storage"

	"github.com/juancarlosTI/monorepo-gestao-jur/backend/documentocase/internal/infraestrutura/consumers"

	"github.com/juancarlosTI/monorepo-gestao-jur/backend/documentocase/internal/infraestrutura/mensageria"
)

func main() {

	ctx, cancel :=
		context.WithCancel(context.Background())

	defer cancel()

	sigChan := make(chan os.Signal, 1)

	signal.Notify(
		sigChan,
		syscall.SIGINT,
		syscall.SIGTERM,
	)

	go func() {
		<-sigChan
		cancel()
	}()

	cfg := config.Load()
	log.Println("Config: ", cfg)

	db, err := sql.Open("pgx", cfg.PostgresDSN())
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	rabbit, err :=
		mensageria.NewRabbitMQ(
			cfg.RabbitURL,
		)

	if err != nil {
		log.Fatal(err)
	}

	defer rabbit.Close()

	if err := mensageria.BootstrapDocumentoMessaging(
		rabbit,
	); err != nil {
		log.Fatal(err)
	}

	storage, err :=
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

	// inboxRepo := NewInboxRepository
	inboxRepo :=
		inbox.NewInboxRepository(
			db,
		)

	ocrService := ocr.NewOCRService()

	processor :=
		services.NewDocumentoProcessor(
			storage,
			ocrService,
		)

	handler := consumer_handlers.NewDocumentoCriadoProcessorHandler(
		processor,
	)

	handlerComInbox :=
		consumers.NewIdempotencyHandler(
			handler,
			inboxRepo,
		)

	dispatcher :=
		consumers.NewDispatcher()

	dispatcher.Register(
		"DocumentoCriado",
		handlerComInbox,
	)

	consumer :=
		consumers.NewDocumentoProcessorConsumer(
			rabbit.ConsumerChannel,
			"documento.processor.queue",
			dispatcher,
			2,
		)

	log.Println(
		"documento processor iniciado",
	)

	if err := consumer.Run(ctx); err != nil {
		log.Fatal(err)
	}
}
