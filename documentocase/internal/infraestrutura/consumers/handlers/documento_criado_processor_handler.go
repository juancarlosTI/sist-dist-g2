package handlers

import (
	"context"
	"encoding/json"
	"log"

	"github.com/rabbitmq/amqp091-go"

	"github.com/juancarlosTI/monorepo-gestao-jur/backend/documentocase/internal/application/services"
	documento "github.com/juancarlosTI/monorepo-gestao-jur/backend/documentocase/internal/dominio/documento"
)

type DocumentoCriadoProcessorHandler struct {
	processor services.DocumentoProcessor
}

func NewDocumentoCriadoProcessorHandler(
	processor services.DocumentoProcessor,
) *DocumentoCriadoProcessorHandler {

	return &DocumentoCriadoProcessorHandler{
		processor: processor,
	}
}

func (h *DocumentoCriadoProcessorHandler) Handle(
	ctx context.Context,
	msg amqp091.Delivery,
) error {

	var payload documento.DocumentoCriadoPayload

	if err := json.Unmarshal(
		msg.Body,
		&payload,
	); err != nil {
		return err
	}

	log.Println(
		"payload deserializado:",
		payload,
	)

	return h.processor.ProcessarDocumento(
		ctx,
		payload.DocumentoID.String(),
		payload.ArquivoID,
	)
}
