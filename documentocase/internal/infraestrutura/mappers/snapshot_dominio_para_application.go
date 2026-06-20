package mappers

import (
	snapshotstore "github.com/juancarlosTI/monorepo-gestao-jur/backend/documentocase/internal/application/snapshot_store"
	documento "github.com/juancarlosTI/monorepo-gestao-jur/backend/documentocase/internal/dominio/documento"
)

func MapSnapshotDominioParaApplication(
	d *documento.DocumentoSnapshot,
) (snapshotstore.DocumentoSnapshot, error) {

	result := make([]string, len(d.ProcessosIDs))

	copy(result, d.ProcessosIDs)

	return snapshotstore.DocumentoSnapshot{
		ID:            d.ID,
		Estado:        d.Estado,
		OrigemCanal:   d.OrigemCanal,
		OrigemSistema: d.OrigemSistema,
		AutorTipo:     d.AutorTipo,
		AutorID:       d.AutorID,
		Versao:        d.Versao,
		ProcessosIDs:  result,
	}, nil
}
