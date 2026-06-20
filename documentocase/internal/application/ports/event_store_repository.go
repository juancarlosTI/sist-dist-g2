package ports

import (
	"context"
	"database/sql"

	eventstore "github.com/juancarlosTI/monorepo-gestao-jur/backend/documentocase/internal/application/event_store"
)

type EventStoreRepository interface {
	Append(
		ctx context.Context,
		tx *sql.Tx,
		evento eventstore.EventStore,
	) error
	LoadStream(
		ctx context.Context,
		agregadoID string,
	) ([]eventstore.EventStore, error)
	LoadFromVersion(
		ctx context.Context,
		agregadoID string,
		versao int,
	) ([]eventstore.EventStore, error)
	GetLastVersion(
		ctx context.Context,
		agregado_id string,
	) (eventstore.EventStore, error)
}
