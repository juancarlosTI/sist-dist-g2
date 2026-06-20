package application

import snapshotstore "github.com/juancarlosTI/monorepo-gestao-jur/backend/documentocase/internal/application/snapshot_store"

type SnapshotPolicy interface {
	DeveCriarSnapshot(
		agregado_versao int,
		snapshotAtual *snapshotstore.DocumentoSnapshot,
		uncommitted_events int,
	) bool
}

type Every50EventsPolicy struct{}

func (p Every50EventsPolicy) DeveCriarSnapshot(
	agregadoVersao int,
	snapshotAtual *snapshotstore.DocumentoSnapshot,
	uncommittedEvents int,
) bool {

	// nada novo para persistir
	if uncommittedEvents == 0 {
		return false
	}

	// primeiro snapshot
	if snapshotAtual == nil {
		return true
	}

	return (agregadoVersao - snapshotAtual.Versao) >= 50
}
