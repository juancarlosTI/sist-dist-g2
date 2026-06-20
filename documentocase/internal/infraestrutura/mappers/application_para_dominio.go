package mappers

import (
	snapshotstore "github.com/juancarlosTI/monorepo-gestao-jur/backend/documentocase/internal/application/snapshot_store"
	documento "github.com/juancarlosTI/monorepo-gestao-jur/backend/documentocase/internal/dominio/documento"
)

func MapSnapshotApplicationParaDominio(
	d *snapshotstore.DocumentoSnapshot,
) (*documento.DocumentoSnapshot, error) {

	if d == nil {
		return nil, nil
	}

	result := make([]string, len(d.ProcessosIDs))

	copy(result, d.ProcessosIDs)

	snapDominio := documento.DocumentoSnapshot{
		ID:            d.ID,
		Estado:        d.Estado,
		OrigemCanal:   d.OrigemCanal,
		OrigemSistema: d.OrigemSistema,
		AutorTipo:     d.AutorTipo,
		AutorID:       d.AutorID,
		Versao:        d.Versao,
		ProcessosIDs:  result,
	}
	return &snapDominio, nil
}
