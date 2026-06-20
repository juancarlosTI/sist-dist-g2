package mappers

import (
	"encoding/json"
	"time"

	eventstore "github.com/juancarlosTI/monorepo-gestao-jur/backend/documentocase/internal/application/event_store"
	documento "github.com/juancarlosTI/monorepo-gestao-jur/backend/documentocase/internal/dominio/documento"
)

func MapEventoDominioParaEventStore(
	e documento.EventoDocumento,
) (eventstore.EventStore, error) {

	payload, err := json.Marshal(e.Payload())
	if err != nil {
		return eventstore.EventStore{}, err
	}

	base := e.Base()

	return eventstore.EventStore{
		EventoID:     base.EventoID.String(),
		EventoVersao: 1,
		EventoNome:   e.Nome(),

		AgregadoID:   base.AgregadoID,
		AgregadoTipo: base.AgregadoTipo,

		CorrelacaoID:  base.CorrelacaoID,
		CausalidadeID: base.CausalidadeID,

		Payload:   payload,
		AutorTipo: base.Autor.Tipo.String(),
		AutorID:   base.Autor.ID,

		OrigemCanal:   base.Origem.Canal.String(),
		OrigemSistema: base.Origem.Sistema.String(),
		RoleTipo:      base.Role.Tipo.String(),

		OcorreuAs: base.OcorreuAs,
		CriadoAs:  time.Now(),
	}, nil
}
