package mappers

import (
	snapshotstore "github.com/juancarlosTI/monorepo-gestao-jur/backend/documentocase/internal/application/snapshot_store"
	documento "github.com/juancarlosTI/monorepo-gestao-jur/backend/documentocase/internal/dominio/documento"
)

func MapDocumentoParaSnapshot(
	d *documento.Documento,
) (snapshotstore.DocumentoSnapshot, error) {
	return snapshotstore.DocumentoSnapshot{

		ID:            d.ID(),
		Estado:        int(d.Estado()),
		Versao:        d.Versao(),
		OrigemCanal:   d.OrigemCanal(),
		OrigemSistema: d.OrigemSistema(),
		AutorTipo:     d.AutorTipo(),
		AutorID:       d.AutorID(),
		ProcessosIDs:  d.Processos(),
	}, nil
}
