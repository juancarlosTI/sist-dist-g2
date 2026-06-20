package mappers

import (
	"time"

	"github.com/juancarlosTI/monorepo-gestao-jur/backend/shared/kernel/eventos"
	outbox "github.com/juancarlosTI/monorepo-gestao-jur/backend/shared/kernel/outbox"
)

func MapEventoIntegracaoParaOutboxEvent(
	e eventos.EventoIntegracao,
	payload []byte,
) (outbox.OutboxEvent, error) {

	base := e.Base()

	return outbox.OutboxEvent{

		EventoID: base.EventoID.String(),

		EventoVersao: base.Versao,

		CorrelacaoID:  base.CorrelacaoID,
		CausalidadeID: base.CausalidadeID,

		RoutingKey: base.RoutingKey,

		AutorTipo: string(base.Autor.Tipo),
		AutorID:   base.Autor.ID,

		OrigemCanal:   string(base.Origem.Canal),
		OrigemSistema: string(base.Origem.Sistema),

		Payload: payload,

		OcorreuAs: base.OcorreuAs,
		CriadoAs:  time.Now(),
	}, nil
}
