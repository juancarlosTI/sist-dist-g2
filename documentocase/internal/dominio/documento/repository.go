package documento

import (
	"context"
	"database/sql"

	eventstore "github.com/juancarlosTI/monorepo-gestao-jur/backend/documentocase/internal/application/event_store"
	"github.com/juancarlosTI/monorepo-gestao-jur/backend/documentocase/internal/dominio/common"
)

type Repository interface {
	Salvar(
		ctx context.Context,
		tx *sql.Tx,
		doc *Documento,
	) ([]eventstore.EventStore, error)
	PorID(ctx context.Context, id common.DocumentoID) (*Documento, error)
}
