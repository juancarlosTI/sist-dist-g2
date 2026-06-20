package workers

import (
	"context"
	"log"
	"time"

	outbox "github.com/juancarlosTI/monorepo-gestao-jur/backend/documentocase/internal/application/ports"
	"github.com/juancarlosTI/monorepo-gestao-jur/backend/documentocase/internal/infraestrutura/mensageria"
)

type OutboxWorker struct {
	repo      outbox.OutboxRepository
	publisher *mensageria.RabbitMQOutboxPublisher
}

func NewOutboxWorker(
	repo outbox.OutboxRepository,
	publisher *mensageria.RabbitMQOutboxPublisher,
) *OutboxWorker {

	return &OutboxWorker{
		repo:      repo,
		publisher: publisher,
	}
}

func (w *OutboxWorker) processarBatch(ctx context.Context) {
	log.Printf("processando batch")

	tx, err := w.repo.BeginTx(ctx)
	if err != nil {
		return
	}

	defer func() {
		if err != nil {
			_ = tx.Rollback()
		}
	}()

	events, err := w.repo.NaoPublicado(ctx, tx, 5)
	if err != nil {
		return
	}

	for _, e := range events {

		log.Printf(
			"publicando evento=%s routing=%s",
			e.EventoNome,
			e.RoutingKey,
		)

		err = w.publisher.Publish(e)

		if err != nil {
			log.Printf("erro publish: %v", err)
			w.repo.IncrementarTentativa(ctx, tx, e.EventoID, e.EventoVersao, err)
			continue
		}

		w.repo.Publicado(ctx, tx, e.EventoID, e.EventoVersao)
	}

	tx.Commit()
}

func (w *OutboxWorker) Run(ctx context.Context) error {
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return nil
		case <-ticker.C:
			w.processarBatch(ctx)

		}
	}
}
