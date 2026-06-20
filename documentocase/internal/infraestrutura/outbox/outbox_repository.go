package outbox

import (
	"context"
	"database/sql"
	"encoding/json"
	"log"
	"time"

	eventstore "github.com/juancarlosTI/monorepo-gestao-jur/backend/documentocase/internal/application/event_store"
	"github.com/juancarlosTI/monorepo-gestao-jur/backend/documentocase/internal/application/ports"
	"github.com/juancarlosTI/monorepo-gestao-jur/backend/shared/kernel/outbox"
)

type OutboxRepositorySQL struct {
	db *sql.DB
}

func NewOutboxRepositorySQL(db *sql.DB) *OutboxRepositorySQL {
	return &OutboxRepositorySQL{
		db: db,
	}
}

func (r *OutboxRepositorySQL) BeginTx(
	ctx context.Context,
) (*sql.Tx, error) {
	return r.db.BeginTx(ctx, nil)
}

func (r *OutboxRepositorySQL) Salvar(
	ctx context.Context,
	tx *sql.Tx,
	origem eventstore.EventStore,
	msg ports.OutboxMessage,
) error {

	log.Printf(
		"salvando outbox evento_origem=%s evento_destino=%s routing=%s",
		origem.EventoNome,
		msg.EventoNome,
		msg.RoutingKey,
	)

	payload, err := json.Marshal(msg.Payload)
	if err != nil {
		return err
	}

	outboxID, err := outbox.NewOutboxID()
	if err != nil {
		return err
	}

	query := `
	INSERT INTO outbox (
		id,
		evento_id,
		evento_versao,
		evento_nome,
		correlacao_id,
		causalidade_id,
		payload,
		routing_key,
		autor_tipo,
		autor_id,
		origem_canal,
		origem_sistema,
		ocorreu_as,
		criado_as,
		publicado_as,
		tentativas,
		ultimo_erro
	)
	VALUES (
		$1,$2,$3,$4,$5,$6,$7,$8,
		$9,$10,$11,$12,$13,$14,$15,$16,$17
	)
	`

	_, err = tx.ExecContext(
		ctx,
		query,

		outboxID,

		origem.EventoID,

		outbox.OutboxEventVersao,

		msg.EventoNome,

		origem.CorrelacaoID,
		origem.CausalidadeID,

		payload,

		msg.RoutingKey,

		origem.AutorTipo,
		origem.AutorID,

		origem.OrigemCanal,
		origem.OrigemSistema,

		origem.OcorreuAs,

		time.Now().UTC(),

		nil, // publicado_as

		0, // tentativas

		nil, // ultimo_erro
	)

	if err != nil {
		log.Printf(
			"erro ao salvar outbox evento=%s: %v",
			msg.EventoNome,
			err,
		)
		return err
	}

	return nil
}

func (r *OutboxRepositorySQL) NaoPublicado(
	ctx context.Context,
	tx *sql.Tx,
	limite int,
) ([]outbox.OutboxEvent, error) {

	query := `
	SELECT
		evento_id,
		evento_versao,
		evento_nome,
		correlacao_id,
		causalidade_id,
		routing_key,
		autor_tipo,
		autor_id,
		origem_canal,
		origem_sistema,
		payload,
		ocorreu_as,
		criado_as,
		publicado_as,
		tentativas,
		ultimo_erro
	FROM outbox
	WHERE publicado_as IS NULL
	ORDER BY criado_as
	FOR UPDATE SKIP LOCKED
	LIMIT $1
	`

	rows, err := tx.QueryContext(
		ctx,
		query,
		limite,
	)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var eventos []outbox.OutboxEvent

	for rows.Next() {

		var (
			eventoID      string
			eventoVersao  int
			eventoNome    string
			correlacaoID  string
			causalidadeID string
			routingKey    string
			autorTipo     string
			autorID       string
			origemCanal   string
			origemSistema string
			payload       []byte
			ocorreuAs     time.Time
			criadoAs      time.Time
			publicadoAs   *time.Time
			tentativas    int
			ultimoErro    *string
		)

		if err := rows.Scan(
			&eventoID,
			&eventoVersao,
			&eventoNome,
			&correlacaoID,
			&causalidadeID,
			&routingKey,
			&autorTipo,
			&autorID,
			&origemCanal,
			&origemSistema,
			&payload,
			&ocorreuAs,
			&criadoAs,
			&publicadoAs,
			&tentativas,
			&ultimoErro,
		); err != nil {
			return nil, err
		}

		eventos = append(eventos, outbox.OutboxEvent{
			EventoID:      eventoID,
			EventoVersao:  eventoVersao,
			EventoNome:    eventoNome,
			CorrelacaoID:  correlacaoID,
			CausalidadeID: causalidadeID,
			RoutingKey:    routingKey,
			AutorTipo:     autorTipo,
			AutorID:       autorID,
			OrigemCanal:   origemCanal,
			OrigemSistema: origemSistema,
			Payload:       payload,
			OcorreuAs:     ocorreuAs,
			CriadoAs:      criadoAs,
			PublicadoAs:   publicadoAs,
			Tentativas:    tentativas,
			UltimoErro:    ultimoErro,
		})
	}

	return eventos, nil
}

func (r *OutboxRepositorySQL) Publicado(
	ctx context.Context,
	tx *sql.Tx,
	id string,
	eventoVersao int,
) error {

	query := `
	UPDATE outbox
	SET publicado_as = NOW(),
	    ultimo_erro = NULL
	WHERE evento_id = $1
	  AND evento_versao = $2
	`

	_, err := tx.ExecContext(
		ctx,
		query,
		id,
		eventoVersao,
	)

	return err
}

func (r *OutboxRepositorySQL) IncrementarTentativa(
	ctx context.Context,
	tx *sql.Tx,
	id string,
	eventoVersao int,
	errMsg error,
) error {

	query := `
	UPDATE outbox
	SET tentativas = tentativas + 1,
	    ultimo_erro = $3
	WHERE evento_id = $1
	  AND evento_versao = $2
	`

	_, err := tx.ExecContext(
		ctx,
		query,
		id,
		eventoVersao,
		errMsg.Error(),
	)

	return err
}
