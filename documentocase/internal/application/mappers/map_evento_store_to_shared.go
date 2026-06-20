package mappers

import (
	"encoding/json"
	"fmt"

	eventstore "github.com/juancarlosTI/monorepo-gestao-jur/backend/documentocase/internal/application/event_store"
	eventos_processor "github.com/juancarlosTI/monorepo-gestao-jur/backend/documentocase/internal/application/eventos_processor"
	"github.com/juancarlosTI/monorepo-gestao-jur/backend/documentocase/internal/application/ports"
	documento "github.com/juancarlosTI/monorepo-gestao-jur/backend/documentocase/internal/dominio/documento"
)

func MapToOutbox(
	evt eventstore.EventStore,
) (ports.OutboxMessage, error) {

	switch evt.EventoNome {

	case "DocumentoCriado":

		var payload documento.DocumentoCriadoPayload

		if err := json.Unmarshal(
			evt.Payload,
			&payload,
		); err != nil {
			return ports.OutboxMessage{}, err
		}

		routingKey := fmt.Sprintf(
			"documento.criado.processor.v%d",
			eventos_processor.DocumentoCriadoProcessorVersao,
		)

		return ports.OutboxMessage{
			EventoNome: evt.EventoNome,
			RoutingKey: routingKey,
			Payload: eventos_processor.DocumentoCriadoProcessorPayload{
				DocumentoID: payload.DocumentoID,
				ArquivoID:   payload.ArquivoID,
			},
		}, nil
	}

	return ports.OutboxMessage{}, nil
}
