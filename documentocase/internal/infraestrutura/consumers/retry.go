package consumers

import "github.com/rabbitmq/amqp091-go"

func HandleRetry(
	msg amqp091.Delivery,
	err error,
	maxRetries int,
) {

	retry := 0

	if v, ok := msg.Headers["x-retry-count"]; ok {
		retry = int(v.(int32))
	}

	if retry >= maxRetries {

		msg.Nack(false, false)

		return
	}

	headers := msg.Headers
	headers["x-retry-count"] = retry + 1

	msg.Nack(false, true)
}
