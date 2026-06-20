package consumers

import (
	"context"
	"encoding/json"
	"log"
	"os"
	"runtime"
	"sync"
	"time"

	"github.com/rabbitmq/amqp091-go"
)

type ProcessingMetrics struct {
	MessageID string `json:"message_id"`

	WorkerID int `json:"worker_id"`

	StartedAt  time.Time `json:"started_at"`
	FinishedAt time.Time `json:"finished_at"`

	DurationMs float64 `json:"duration_ms"`

	GoroutinesBefore int `json:"goroutines_before"`
	GoroutinesAfter  int `json:"goroutines_after"`

	AllocBeforeMB float64 `json:"alloc_before_mb"`
	AllocAfterMB  float64 `json:"alloc_after_mb"`

	HeapBeforeMB float64 `json:"heap_before_mb"`
	HeapAfterMB  float64 `json:"heap_after_mb"`

	NumCPU     int `json:"num_cpu"`
	GOMAXPROCS int `json:"gomaxprocs"`

	Success bool `json:"success"`
}

type ConsumerStats struct {
	Processed uint64
	Failed    uint64
}

type DocumentoProcessorConsumer struct {
	channel    *amqp091.Channel
	queueName  string
	dispatcher *Dispatcher

	workers int

	stats ConsumerStats
	mutex sync.Mutex
}

func SaveMetrics(
	m ProcessingMetrics,
) error {

	f, err :=
		os.OpenFile(
			"metrics.jsonl",
			os.O_APPEND|
				os.O_CREATE|
				os.O_WRONLY,
			0644,
		)

	if err != nil {
		return err
	}

	defer f.Close()

	data, err :=
		json.Marshal(m)

	if err != nil {
		return err
	}

	_, err =
		f.Write(
			append(data, '\n'),
		)

	return err
}

func NewDocumentoProcessorConsumer(
	channel *amqp091.Channel,
	queueName string,
	dispatcher *Dispatcher,
	workers int,
) *DocumentoProcessorConsumer {

	return &DocumentoProcessorConsumer{
		channel:    channel,
		queueName:  queueName,
		dispatcher: dispatcher,
		workers:    workers,
	}
}

func (c *DocumentoProcessorConsumer) logMetrics(
	m ProcessingMetrics,
) {

	log.Printf(
		`
================ PROCESSING METRICS ================
worker=%d
message=%s
success=%v

started_at=%s
finished_at=%s
duration/ms=%f

goroutines_before=%d
goroutines_after=%d

alloc_before=%fMB
alloc_after=%fMB

heap_before=%fMB
heap_after=%fMB
===================================================
`,
		m.WorkerID,
		m.MessageID,
		m.Success,

		m.StartedAt,
		m.FinishedAt,
		m.DurationMs,

		m.GoroutinesBefore,
		m.GoroutinesAfter,

		m.AllocBeforeMB,
		m.AllocAfterMB,

		m.HeapBeforeMB,
		m.HeapAfterMB,
	)
}

func (c *DocumentoProcessorConsumer) Run(
	ctx context.Context,
) error {

	err := c.channel.Qos(
		c.workers,
		0,
		false,
	)

	if err != nil {
		return err
	}

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

	jobs := make(chan amqp091.Delivery)

	var wg sync.WaitGroup

	// ==================================================
	// MÉTRICAS PERIÓDICAS
	// ==================================================

	// go func() {

	// 	ticker :=
	// 		time.NewTicker(
	// 			30 * time.Second,
	// 		)

	// 	defer ticker.Stop()

	// 	for {

	// 		select {

	// 		case <-ctx.Done():
	// 			return

	// 		case <-ticker.C:

	// 			c.mutex.Lock()

	// 			log.Printf(
	// 				"[THROUGHPUT] processed=%d failed=%d goroutines=%d",
	// 				c.stats.Processed,
	// 				c.stats.Failed,
	// 				runtime.NumGoroutine(),
	// 			)

	// 			c.mutex.Unlock()
	// 		}
	// 	}
	// }()

	// ==================================================
	// WORKER POOL
	// ==================================================

	for i := 0; i < c.workers; i++ {

		wg.Add(1)

		go func(workerID int) {

			defer wg.Done()

			log.Printf(
				"worker %d iniciado",
				workerID,
			)

			for {

				select {

				case <-ctx.Done():
					return

				case msg, ok := <-jobs:

					if !ok {
						return
					}

					start := time.Now()

					var before runtime.MemStats
					runtime.ReadMemStats(&before)

					goroutinesBefore :=
						runtime.NumGoroutine()

					log.Printf(
						"worker=%d message=%s processando",
						workerID,
						msg.MessageId,
					)

					err :=
						c.dispatcher.Dispatch(
							ctx,
							msg,
						)

					var after runtime.MemStats
					runtime.ReadMemStats(&after)

					goroutinesAfter :=
						runtime.NumGoroutine()

					metrics := ProcessingMetrics{
						WorkerID:  workerID,
						MessageID: msg.MessageId,

						StartedAt:  start,
						FinishedAt: time.Now(),
						DurationMs: float64(
							time.Since(start).Nanoseconds(),
						) / 1e6,

						GoroutinesBefore: goroutinesBefore,
						GoroutinesAfter:  goroutinesAfter,

						AllocBeforeMB: float64(before.Alloc) / 1024.0 / 1024.0,
						AllocAfterMB:  float64(after.Alloc) / 1024.0 / 1024.0,

						HeapBeforeMB: float64(before.HeapAlloc) / 1024.0 / 1024.0,
						HeapAfterMB:  float64(after.HeapAlloc) / 1024.0 / 1024.0,

						Success: err == nil,
					}

					if err != nil {

						c.mutex.Lock()
						c.stats.Failed++
						c.mutex.Unlock()

						log.Printf(
							"worker=%d erro=%v",
							workerID,
							err,
						)

						c.logMetrics(metrics)

						msg.Nack(
							false,
							true,
						)

						continue
					}

					c.mutex.Lock()
					c.stats.Processed++
					c.mutex.Unlock()

					msg.Ack(false)

					c.logMetrics(metrics)

					log.Printf(
						"worker=%d message=%s concluído",
						workerID,
						msg.MessageId,
					)
					SaveMetrics(metrics)
				}
			}

		}(i + 1)
	}

	// ==================================================
	// RECEBIMENTO DE MENSAGENS
	// ==================================================

	for {

		select {

		case <-ctx.Done():

			log.Println(
				"encerrando consumer...",
			)

			close(jobs)

			wg.Wait()

			log.Println(
				"todos os workers finalizados",
			)

			return nil

		case msg, ok := <-msgs:

			if !ok {

				log.Println(
					"canal RabbitMQ encerrado",
				)

				close(jobs)

				wg.Wait()

				return nil
			}

			jobs <- msg
		}
	}
}
