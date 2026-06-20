package eventosprocessor

import (
	dominio_common "github.com/juancarlosTI/monorepo-gestao-jur/backend/documentocase/internal/dominio/common"
)

type Envelope struct {
	EventoID string

	EventoNome string

	RoutingKey string

	Payload []byte

	CorrelacaoID  string
	CausalidadeID string
}

const DocumentoCriadoProcessorVersao = 1

type DocumentoCriadoProcessorPayload struct {
	DocumentoID dominio_common.DocumentoID `json:"documento_id"`
	ArquivoID   string                     `json:"arquivo_id"`
}

type DocumentoCriadoProcessor struct {
	base    Envelope
	payload DocumentoCriadoProcessorPayload
}

func (e DocumentoCriadoProcessor) Nome() string {
	return "DocumentoCriado"
}

func (e DocumentoCriadoProcessor) Base() Envelope {
	return e.base
}

func (e DocumentoCriadoProcessor) Payload() any {
	return e.payload
}
