package mensageria

import "github.com/rabbitmq/amqp091-go"

func DeclareDocumentoProcessorQueues(
	ch *amqp091.Channel,
) error {

	args := amqp091.Table{
		"x-dead-letter-exchange": "documentos.dlx",
	}

	q, err := ch.QueueDeclare(
		"documento.processor.queue",
		true,
		false,
		false,
		false,
		args,
	)

	if err != nil {
		return err
	}

	err = ch.QueueBind(
		q.Name,
		"documento.criado.processor.*",
		"documentos.eventos",
		false,
		nil,
	)

	if err != nil {
		return err
	}

	dlq, err := ch.QueueDeclare(
		"documento.processor.dlq",
		true,
		false,
		false,
		false,
		nil,
	)

	if err != nil {
		return err
	}

	return ch.QueueBind(
		dlq.Name,
		"documento.#",
		"documentos.dlx",
		false,
		nil,
	)
}
