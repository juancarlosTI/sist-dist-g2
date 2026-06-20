package mensageria

import "github.com/rabbitmq/amqp091-go"

func DeclareDocumentoExchange(
	ch *amqp091.Channel,
) error {

	return ch.ExchangeDeclare(
		"documentos.eventos",
		"topic",
		true,
		false,
		false,
		false,
		nil,
	)
}

func DeclareDocumentoDLX(
	ch *amqp091.Channel,
) error {

	return ch.ExchangeDeclare(
		"documentos.dlx",
		"topic",
		true,
		false,
		false,
		false,
		nil,
	)
}
