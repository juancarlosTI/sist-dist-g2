package ports

import (
	"context"
	"database/sql"

	eventstore "github.com/juancarlosTI/monorepo-gestao-jur/backend/documentocase/internal/application/event_store"
	"github.com/juancarlosTI/monorepo-gestao-jur/backend/shared/kernel/outbox"
)

type OutboxMessage struct {
	EventoNome string
	RoutingKey string
	Payload    any
}
type OutboxRepository interface {
	Salvar(
		ctx context.Context,
		tx *sql.Tx,
		origem eventstore.EventStore,
		msg OutboxMessage,
	) error

	NaoPublicado(
		ctx context.Context,
		tx *sql.Tx,
		limite int,
	) ([]outbox.OutboxEvent, error)

	Publicado(
		ctx context.Context,
		tx *sql.Tx,
		id string,
		eventoVersao int,
	) error

	IncrementarTentativa(
		ctx context.Context,
		tx *sql.Tx,
		id string,
		eventoVersao int,
		err error,
	) error

	BeginTx(
		ctx context.Context,
	) (*sql.Tx, error)
}
