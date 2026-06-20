package consumers

import (
	"context"

	"github.com/juancarlosTI/monorepo-gestao-jur/backend/documentocase/internal/infraestrutura/inbox"
	"github.com/rabbitmq/amqp091-go"
)

type InboxHandler struct {
	next Handler
	repo *inbox.InboxRepository
}

func NewIdempotencyHandler(
	next Handler,
	repo *inbox.InboxRepository,
) *InboxHandler {

	return &InboxHandler{
		next: next,
		repo: repo,
	}
}

func (h *InboxHandler) Handle(
	ctx context.Context,
	msg amqp091.Delivery,
) error {

	eventID := msg.MessageId

	if eventID == "" {
		return h.next.Handle(ctx, msg)
	}

	processed, err := h.repo.JaProcessado(ctx, eventID)
	if err != nil {
		return err
	}

	if processed {
		return nil
	}

	err = h.next.Handle(ctx, msg)
	if err != nil {
		return err
	}

	return h.repo.MarcarComoProcessado(ctx, eventID)
}
