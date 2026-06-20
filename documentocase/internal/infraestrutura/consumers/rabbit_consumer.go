package consumers

import (
	"context"
	"log"

	"github.com/rabbitmq/amqp091-go"
)

type RabbitConsumer struct {
	channel    *amqp091.Channel
	queueName  string
	dispatcher *Dispatcher
}

func NewRabbitConsumer(
	channel *amqp091.Channel,
	queueName string,
	dispatcher *Dispatcher,
) *RabbitConsumer {

	return &RabbitConsumer{
		channel:    channel,
		queueName:  queueName,
		dispatcher: dispatcher,
	}
}

func (c *RabbitConsumer) Run(
	ctx context.Context,
) error {

	msgs, err := c.channel.Consume(
		c.queueName,
		"",
		false,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		return err
	}

	for {
		select {

		case <-ctx.Done():
			return nil

		case msg, ok := <-msgs:

			if !ok {
				return nil
			}

			err := c.dispatcher.Dispatch(
				ctx,
				msg,
			)

			if err != nil {

				log.Printf(
					"erro processando mensagem: %v",
					err,
				)

				msg.Nack(false, true)

				continue
			}

			msg.Ack(false)
		}
	}
}
