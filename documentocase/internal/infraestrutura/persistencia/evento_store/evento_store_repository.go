package eventostore

import (
	"context"
	"database/sql"
	"errors"

	"github.com/jackc/pgconn"
	application "github.com/juancarlosTI/monorepo-gestao-jur/backend/documentocase/internal/application/common"
	eventstore "github.com/juancarlosTI/monorepo-gestao-jur/backend/documentocase/internal/application/event_store"
)

type EventStoreRepositorySQL struct {
	db *sql.DB
}

func NewEventStoreRepository(
	db *sql.DB,
) *EventStoreRepositorySQL {

	return &EventStoreRepositorySQL{
		db: db,
	}
}

func (r *EventStoreRepositorySQL) Append(
	ctx context.Context,
	tx *sql.Tx,
	eventos []eventstore.EventStore,
	versao_esperada int,
) error {

	if len(eventos) == 0 {
		return nil
	}

	query := `
	INSERT INTO evento_store (
		evento_id,
		agregado_id,
		agregado_tipo,
		evento_versao,
		evento_nome,
		correlacao_id,
		causalidade_id,
		payload,
		autor_tipo,
		autor_id,
		origem_canal,
		origem_sistema,
		role_tipo,
		ocorreu_as,
		criado_as
	)
	VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13,$14,$15)
	`

	for i, event := range eventos {

		version := versao_esperada + i + 1

		_, err := tx.ExecContext(
			ctx,
			query,
			event.EventoID,
			event.AgregadoID,
			event.AgregadoTipo,
			version,
			event.EventoNome,
			event.CorrelacaoID,
			event.CausalidadeID,
			event.Payload,
			event.AutorTipo,
			event.AutorID,
			event.OrigemCanal,
			event.OrigemSistema,
			event.RoleTipo,
			event.OcorreuAs,
			event.CriadoAs,
		)

		if err != nil {

			// Detecta violação de unique index
			var pgErr *pgconn.PgError
			if errors.As(err, &pgErr) && pgErr.Code == "23505" {
				return application.ErrConcurrencyConflict
			}

			return err
		}

	}

	return nil
}

func (r *EventStoreRepositorySQL) Carregar(
	ctx context.Context,
	agregado_id string,
) ([]eventstore.EventStore, error) {

	query := `
	SELECT
		evento_id,
		agregado_id,
		agregado_tipo,
		evento_versao,
		evento_nome,
		correlacao_id,
		causalidade_id,
		payload,
		autor_tipo,
		autor_id,
		origem_canal,
		origem_sistema,
		ocorreu_as,
		criado_as
	FROM event_store
	WHERE agregado_id = $1
	ORDER BY evento_versao ASC
	`

	rows, err := r.db.QueryContext(ctx, query, agregado_id)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	events := []eventstore.EventStore{}

	for rows.Next() {

		var evt eventstore.EventStore

		err := rows.Scan(
			&evt.EventoID,
			&evt.AgregadoID,
			&evt.AgregadoTipo,
			&evt.EventoVersao,
			&evt.EventoNome,
			&evt.CorrelacaoID,
			&evt.CausalidadeID,
			&evt.Payload,
			&evt.AutorTipo,
			&evt.AutorID,
			&evt.OrigemCanal,
			&evt.OrigemSistema,
			&evt.OcorreuAs,
			&evt.CriadoAs,
		)

		if err != nil {
			return nil, err
		}

		events = append(events, evt)
	}

	return events, nil
}

func (r *EventStoreRepositorySQL) CarregarDaVersao(
	ctx context.Context,
	agregado_id string,
	versao int,
) ([]eventstore.EventStore, error) {

	query := `
	SELECT
		evento_id,
		agregado_id,
		agregado_tipo,
		evento_versao,
		evento_nome,
		correlacao_id,
		causalidade_id,
		payload,
		autor_tipo,
		autor_id,
		origem_canal,
		origem_sistema,
		ocorreu_as,
		criado_as
	FROM event_store
	WHERE agregado_id = $1
	AND evento_versao > $2
	ORDER BY evento_versao ASC
	`

	rows, err := r.db.QueryContext(
		ctx,
		query,
		agregado_id,
		versao,
	)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var events []eventstore.EventStore

	for rows.Next() {

		var evt eventstore.EventStore

		err := rows.Scan(
			&evt.EventoID,
			&evt.AgregadoID,
			&evt.AgregadoTipo,
			&evt.EventoVersao,
			&evt.EventoNome,
			&evt.CorrelacaoID,
			&evt.CausalidadeID,
			&evt.Payload,
			&evt.OcorreuAs,
			&evt.CriadoAs,
		)

		if err != nil {
			return nil, err
		}

		events = append(events, evt)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return events, nil
}
