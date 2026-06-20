package mensageria

import (
	amqp "github.com/rabbitmq/amqp091-go"
)

type RabbitMQ struct {
	Conn *amqp.Connection

	AdminChannel     *amqp.Channel
	PublisherChannel *amqp.Channel
	ConsumerChannel  *amqp.Channel
}

func NewRabbitMQ(url string) (*RabbitMQ, error) {

	conn, err := amqp.Dial(url)
	if err != nil {
		return nil, err
	}

	adminCh, err := conn.Channel()
	if err != nil {
		conn.Close()
		return nil, err
	}

	publisherCh, err := conn.Channel()
	if err != nil {
		adminCh.Close()
		conn.Close()
		return nil, err
	}

	consumerCh, err := conn.Channel()
	if err != nil {
		publisherCh.Close()
		adminCh.Close()
		conn.Close()
		return nil, err
	}

	return &RabbitMQ{
		Conn:             conn,
		AdminChannel:     adminCh,
		PublisherChannel: publisherCh,
		ConsumerChannel:  consumerCh,
	}, nil
}

func (r *RabbitMQ) Close() {

	if r.ConsumerChannel != nil {
		_ = r.ConsumerChannel.Close()
	}

	if r.PublisherChannel != nil {
		_ = r.PublisherChannel.Close()
	}

	if r.AdminChannel != nil {
		_ = r.AdminChannel.Close()
	}

	if r.Conn != nil {
		_ = r.Conn.Close()
	}
}
