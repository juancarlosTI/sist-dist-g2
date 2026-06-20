package mappers

import (
	"encoding/json"

	eventstore "github.com/juancarlosTI/monorepo-gestao-jur/backend/documentocase/internal/application/event_store"
	"github.com/juancarlosTI/monorepo-gestao-jur/backend/documentocase/internal/dominio/documento"
)

func DeserializeEventoDominio(
	payload []byte,
) (documento.EventoDocumento, error) {

	var envelope documento.EventoDocumento

	err := json.Unmarshal(payload, &envelope)
	if err != nil {
		return nil, err
	}

	// switch envelope.Base().AgregadoTipo {

	// case "ProcessoCriado":

	// 	var e processo.ProcessoCriado

	// 	err := json.Unmarshal(payload, &e)

	// 	return e, err

	// case "PendenciaRegistrada":

	// 	var e processo.PendenciaRegistrada

	// 	err := json.Unmarshal(payload, &e)

	// 	return e, err

	// case "PendenciaResolvida":

	// 	var e processo.PendenciaResolvida

	// 	err := json.Unmarshal(payload, &e)

	// 	return e, err

	// case "ProcessoConcluido":

	// 	var e processo.ProcessoConcluido

	// 	err := json.Unmarshal(payload, &e)

	// 	return e, err

	// case "ProcessoConcluidoTecnico":

	// 	var e processo.ProcessoTecnicoConcluido

	// 	err := json.Unmarshal(payload, &e)

	// 	return e, err

	// default:

	// 	return nil, errors.New("evento desconhecido")
	// }
	return nil, err
}

func MapEventStoreParaDominio(
	events []eventstore.EventStore,
) ([]documento.EventoDocumento, error) {

	result := make([]documento.EventoDocumento, 0, len(events))

	for _, evtStore := range events {

		evtDominio, err :=
			DeserializeEventoDominio(evtStore.Payload)

		if err != nil {
			return nil, err
		}

		result = append(result, evtDominio)
	}

	return result, nil
}
