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
	"github.com/juancarlosTI/monorepo-gestao-jur/backend/documentocase/internal/infraestrutura/mensageria"
	outbox_repo "github.com/juancarlosTI/monorepo-gestao-jur/backend/documentocase/internal/infraestrutura/outbox"
	workers "github.com/juancarlosTI/monorepo-gestao-jur/backend/documentocase/internal/infraestrutura/workers"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Captura sinais em paralelo
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		s := <-sigChan
		log.Println("Shutdown recebido… cancelando contexto:", s)
		cancel()
	}()

	// Db
	cfg := config.Load()

	db, err := sql.Open("pgx", cfg.PostgresDSN())
	if err != nil {
		log.Fatal(err)
	}

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	rabbit, err := mensageria.NewRabbitMQ(cfg.RabbitURL)
	if err != nil {
		log.Fatal(err)
	}

	defer rabbit.Close()

	if err := mensageria.BootstrapDocumentoMessaging(
		rabbit,
	); err != nil {
		log.Fatal(err)
	}

	outboxRepo :=
		outbox_repo.NewOutboxRepositorySQL(db)

	publisher :=
		mensageria.NewRabbitMQOutboxPublisher(
			rabbit.PublisherChannel,
			"documentos.eventos",
		)

	worker :=
		workers.NewOutboxWorker(
			outboxRepo,
			publisher,
		)

	if err := worker.Run(ctx); err != nil {
		log.Fatal(err)
	}
}
