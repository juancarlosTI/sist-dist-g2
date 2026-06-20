package handlers

import (
	"context"
	"encoding/json"

	application "github.com/juancarlosTI/monorepo-gestao-jur/backend/documentocase/internal/application/handlers"
	shared "github.com/juancarlosTI/monorepo-gestao-jur/backend/shared/kernel/processo"

	"github.com/rabbitmq/amqp091-go"
)

type DocumentoCriacaoSolicitadaConsumer struct {
	appHandler *application.DocumentoCriacaoSolicitadaConsumer
}

func NewDocumentoCriacaoSolicitadaConsumer(
	appHandler *application.DocumentoCriacaoSolicitadaConsumer,
) *DocumentoCriacaoSolicitadaConsumer {
	return &DocumentoCriacaoSolicitadaConsumer{
		appHandler: appHandler,
	}
}

func (c *DocumentoCriacaoSolicitadaConsumer) Handle(
	ctx context.Context,
	msg amqp091.Delivery,
) error {

	var event shared.DocumentoCriacaoSolicitada

	if err := json.Unmarshal(msg.Body, &event); err != nil {
		return err
	}

	// Agora chamamos o application handler corretamente
	return c.appHandler.Handle(ctx, event)
}
