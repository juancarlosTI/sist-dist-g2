package snapshotstore

import (
	"context"
	"database/sql"

	snapshot_store "github.com/juancarlosTI/monorepo-gestao-jur/backend/documentocase/internal/application/snapshot_store"
	"github.com/juancarlosTI/monorepo-gestao-jur/backend/documentocase/internal/dominio/documento"
	"github.com/juancarlosTI/monorepo-gestao-jur/backend/documentocase/internal/infraestrutura/mappers"
)

type SnapshotStoreRepositorySQL struct {
	db *sql.DB
}

func NewSnapshotStoreRepository(
	db *sql.DB,
) *SnapshotStoreRepositorySQL {

	return &SnapshotStoreRepositorySQL{
		db: db,
	}
}

func DeveCriarSnapshot(
	documento_versao int,
	versao_snapshot int,
) bool {

	const snapshot_interval = 50

	return documento_versao-versao_snapshot >= snapshot_interval
}

func CriarSnapshotDominio(
	d *documento.Documento,
) (*snapshot_store.DocumentoSnapshot, error) {
	to_application_snapshot, err := mappers.MapDocumentoParaSnapshot(d)
	if err != nil {
		return &snapshot_store.DocumentoSnapshot{}, nil
	}
	return &to_application_snapshot, nil
}

func (r *SnapshotStoreRepositorySQL) CarregarSnapshot(
	ctx context.Context,
	agregado_id string,
) (*snapshot_store.DocumentoSnapshot, error) {

	query := `
	SELECT
		id,
		estado,
		origem_canal,
		origem_sistema,
		autor_tipo,
		autor_id,
		versao,
		processos_ids
	FROM snapshot_documento
	WHERE id = $1
	`

	row := r.db.QueryRowContext(
		ctx,
		query,
		agregado_id,
	)

	var snapshot snapshot_store.DocumentoSnapshot

	err := row.Scan(
		&snapshot.ID,
		&snapshot.Estado,
		&snapshot.OrigemCanal,
		&snapshot.OrigemSistema,
		&snapshot.AutorTipo,
		&snapshot.AutorID,
		&snapshot.Versao,
		&snapshot.ProcessosIDs,
	)

	if err == sql.ErrNoRows {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return &snapshot, nil
}

func (r *SnapshotStoreRepositorySQL) CarregarSnapshotTx(
	ctx context.Context,
	tx *sql.Tx,
	agregado_id string,

) (*snapshot_store.DocumentoSnapshot, error) {

	query := `
	SELECT
		id,
		estado,
		origem_canal,
		origem_sistema,
		autor_tipo,
		autor_id,
		versao,
		processos_ids
	FROM snapshot_documento
	WHERE id = $1
	`

	row := tx.QueryRowContext(
		ctx,
		query,
		agregado_id,
	)

	var snapshot snapshot_store.DocumentoSnapshot

	err := row.Scan(
		&snapshot.ID,
		&snapshot.Estado,
		&snapshot.OrigemCanal,
		&snapshot.OrigemSistema,
		&snapshot.AutorTipo,
		&snapshot.AutorID,
		&snapshot.Versao,
		&snapshot.ProcessosIDs,
	)

	if err == sql.ErrNoRows {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return &snapshot, nil
}

func (r *SnapshotStoreRepositorySQL) SalvarSnapshot(
	ctx context.Context,
	tx *sql.Tx,
	snapshot snapshot_store.DocumentoSnapshot,

) error {

	query := `
	INSERT INTO snapshot_documento (
		id,
		estado,
		origem_canal,
		origem_sistema,
		autor_tipo,
		autor_id,
		versao,
		processos_ids
	)
	VALUES ($1,$2,$3,$4,$5,$6,$7,$8)
	ON CONFLICT (id)
	DO UPDATE SET
	versao = EXCLUDED.versao
	`

	_, err := tx.ExecContext(
		ctx,
		query,
		snapshot.ID,
		snapshot.Estado,
		snapshot.OrigemCanal,
		snapshot.OrigemSistema,
		snapshot.AutorTipo,
		snapshot.AutorID,
		snapshot.Versao,
		snapshot.ProcessosIDs,
	)

	return err
}
