package consumers

import (
	"context"
	"fmt"
	"log"

	"github.com/rabbitmq/amqp091-go"
)

type Handler interface {
	Handle(
		ctx context.Context,
		msg amqp091.Delivery,
	) error
}

type Dispatcher struct {
	handlers map[string]Handler
}

func NewDispatcher() *Dispatcher {

	return &Dispatcher{
		handlers: make(map[string]Handler),
	}
}

func (d *Dispatcher) Register(
	eventType string,
	h Handler,
) {
	d.handlers[eventType] = h
}

func (d *Dispatcher) Dispatch(
	ctx context.Context,
	msg amqp091.Delivery,
) error {

	eventType := msg.Type

	if eventType == "" {
		log.Println("mensagem sem type")
		return fmt.Errorf("mensagem sem type")
	}

	h, ok :=
		d.handlers[eventType]

	if !ok {
		log.Printf(
			"evento sem handler | type=%s routing=%s",
			msg.Type,
			msg.RoutingKey,
		)
		return fmt.Errorf("handler nao encontrado")
	}

	return h.Handle(ctx, msg)
}
