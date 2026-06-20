package mensageria

import "fmt"

func BootstrapDocumentoMessaging(
	rabbit *RabbitMQ,
) error {

	if rabbit == nil {
		return fmt.Errorf("rabbitmq nil")
	}

	if rabbit.AdminChannel == nil {
		return fmt.Errorf("admin channel nil")
	}

	if err := DeclareDocumentoExchange(
		rabbit.AdminChannel,
	); err != nil {
		return fmt.Errorf(
			"declare documento exchange: %w",
			err,
		)
	}

	if err := DeclareDocumentoDLX(
		rabbit.AdminChannel,
	); err != nil {
		return fmt.Errorf(
			"declare documento dlx: %w",
			err,
		)
	}

	if err := DeclareDocumentoProcessorQueues(
		rabbit.AdminChannel,
	); err != nil {
		return fmt.Errorf(
			"declare documento queues: %w",
			err,
		)
	}

	return nil
}
