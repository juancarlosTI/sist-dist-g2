package mensageria

import (
	"github.com/juancarlosTI/monorepo-gestao-jur/backend/shared/kernel/outbox"
	"github.com/rabbitmq/amqp091-go"
)

type RabbitMQOutboxPublisher struct {
	channel  *amqp091.Channel
	exchange string
	appID    string
}

func NewRabbitMQOutboxPublisher(
	channel *amqp091.Channel,
	exchange string,
) *RabbitMQOutboxPublisher {

	return &RabbitMQOutboxPublisher{
		channel:  channel,
		exchange: exchange,
	}
}

func (p *RabbitMQOutboxPublisher) Publish(e outbox.OutboxEvent) error {
	return p.channel.Publish(
		p.exchange,
		e.RoutingKey,
		false,
		false,
		amqp091.Publishing{
			ContentType: "application/json",
			Body:        e.Payload,
			MessageId:   e.EventoID,
			Timestamp:   e.OcorreuAs,
			Type:        e.EventoNome,
			AppId:       "documento-application-service",
			Headers: amqp091.Table{
				"schema": e.EventoVersao,
			},
		},
	)
}
