package consumers

// import (
// 	"context"
// 	"errors"
// 	"log"

// 	"github.com/rabbitmq/amqp091-go"
// )

// type ProcessoConsumer struct {
// 	channel    *amqp091.Channel
// 	dispatcher *Dispatcher
// }

// func NewProcessoConsumer(
// 	ch *amqp091.Channel,
// 	dispatcher *Dispatcher,
// ) *ProcessoConsumer {

// 	return &ProcessoConsumer{
// 		channel:    ch,
// 		dispatcher: dispatcher,
// 	}
// }

// func (c *ProcessoConsumer) Run(ctx context.Context) error {
// 	msgs, err := c.channel.Consume(
// 		"documento.processo.queue",
// 		"",
// 		false,
// 		false,
// 		false,
// 		false,
// 		nil,
// 	)

// 	if err != nil {
// 		return err
// 	}

// 	for {

// 		select {

// 		case <-ctx.Done():
// 			return nil

// 		case msg, ok := <-msgs:
// 			if !ok {
// 				return errors.New("channel fechado")
// 			}
// 			err := c.dispatcher.Dispatch(
// 				ctx,
// 				msg,
// 			)

// 			if err != nil {

// 				msg.Nack(false, true)

// 				log.Println("Erro processando mensagem:", err)

// 				continue
// 			}

// 			msg.Ack(false)

// 		}
// 	}
// }
