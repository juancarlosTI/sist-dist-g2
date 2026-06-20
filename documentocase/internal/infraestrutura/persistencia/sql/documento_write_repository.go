package sql

import (
	"context"
	"database/sql"
	"log"

	application "github.com/juancarlosTI/monorepo-gestao-jur/backend/documentocase/internal/application/common"
	eventstore "github.com/juancarlosTI/monorepo-gestao-jur/backend/documentocase/internal/application/event_store"
	dominio_common "github.com/juancarlosTI/monorepo-gestao-jur/backend/documentocase/internal/dominio/common"
	"github.com/juancarlosTI/monorepo-gestao-jur/backend/documentocase/internal/dominio/documento"
	mappers "github.com/juancarlosTI/monorepo-gestao-jur/backend/documentocase/internal/infraestrutura/mappers"
	eventostore "github.com/juancarlosTI/monorepo-gestao-jur/backend/documentocase/internal/infraestrutura/persistencia/evento_store"
	snapshotstore "github.com/juancarlosTI/monorepo-gestao-jur/backend/documentocase/internal/infraestrutura/persistencia/snapshot_store"
)

type DocumentoSQLRepository struct {
	db             *sql.DB
	eventStoreRepo *eventostore.EventStoreRepositorySQL
	snapshotRepo   *snapshotstore.SnapshotStoreRepositorySQL
	snapshotPolicy application.SnapshotPolicy
}

func NewDocumentoRepositorySQLHandler(
	db *sql.DB,
	eventStoreRepo *eventostore.EventStoreRepositorySQL,
	snapshotRepo *snapshotstore.SnapshotStoreRepositorySQL,
	snapshotPolicy application.SnapshotPolicy,
) *DocumentoSQLRepository {
	return &DocumentoSQLRepository{
		db:             db,
		eventStoreRepo: eventStoreRepo,
		snapshotRepo:   snapshotRepo,
		snapshotPolicy: snapshotPolicy,
	}
}

func (r *DocumentoSQLRepository) Salvar(
	ctx context.Context,
	tx *sql.Tx,
	d *documento.Documento,
) ([]eventstore.EventStore, error) {

	versaoEsperada := d.Versao() - len(d.Eventos())

	snapshotAtual, err :=
		r.snapshotRepo.CarregarSnapshotTx(
			ctx,
			tx,
			d.ID(),
		)

	if err != nil {
		return nil, err
	}

	eventos := make([]eventstore.EventStore, 0)

	for i, evt := range d.Eventos() {

		esEvt, err :=
			mappers.MapEventoDominioParaEventStore(evt)

		if err != nil {
			return nil, err
		}

		esEvt.EventoVersao =
			versaoEsperada + i + 1

		eventos = append(eventos, esEvt)
	}

	// append event store
	err = r.eventStoreRepo.Append(
		ctx,
		tx,
		eventos,
		versaoEsperada,
	)

	if err != nil {
		return nil, err
	}

	// 2 — Criar snapshot se necessário

	// snapshot
	if r.snapshotPolicy.DeveCriarSnapshot(
		d.Versao(),
		snapshotAtual,
		len(d.Eventos()),
	) {

		snapDominio := d.GerarSnapshot()

		snapApp, err :=
			mappers.MapSnapshotDominioParaApplication(
				&snapDominio,
			)

		if err != nil {
			return nil, err
		}

		err = r.snapshotRepo.SalvarSnapshot(
			ctx,
			tx,
			snapApp,
		)

		log.Println("Write model salvar: ", err)

		if err != nil {
			return nil, err
		}
	}

	d.LimparEventos()

	return eventos, nil
}

func (r *DocumentoSQLRepository) PorID(
	ctx context.Context,
	id dominio_common.DocumentoID,
) (*documento.Documento, error) {

	snapshot, err := r.snapshotRepo.CarregarSnapshot(ctx, id.String())
	if err != nil {
		return nil, err
	}

	// Application -> Dominio

	snapshotDominio, err := mappers.MapSnapshotApplicationParaDominio(snapshot)
	if err != nil {
		return nil, err
	}

	var eventos []eventstore.EventStore

	if snapshotDominio != nil {

		eventos, err = r.eventStoreRepo.CarregarDaVersao(
			ctx,
			id.String(),
			snapshotDominio.Versao,
		)

	} else {

		eventos, err = r.eventStoreRepo.Carregar(
			ctx,
			id.String(),
		)
	}

	if snapshotDominio == nil && len(eventos) == 0 {
		return nil, application.ErrNotFound
	}

	eventosDominio, err := mappers.MapEventStoreParaDominio(eventos)
	if err != nil {
		return nil, err
	}

	return documento.ReconstruirAgregado(
		snapshotDominio,
		eventosDominio,
	)
}
