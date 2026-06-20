package outbox

import "time"

const OutboxEventVersao = 1

type OutboxEvent struct {
	EventoID     string
	EventoVersao int
	EventoNome   string

	CorrelacaoID  string
	CausalidadeID string

	RoutingKey string

	AutorTipo string
	AutorID   string

	OrigemCanal   string
	OrigemSistema string

	Payload []byte

	OcorreuAs time.Time
	CriadoAs  time.Time

	PublicadoAs *time.Time
	Tentativas  int
	UltimoErro  *string
}
