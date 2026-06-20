package eventstore

import "time"

type EventStore struct {
	EventoID     string
	EventoVersao int
	EventoNome   string

	AgregadoID   string
	AgregadoTipo string

	CorrelacaoID  string
	CausalidadeID string

	Payload []byte

	AutorTipo     string
	AutorID       string
	OrigemCanal   string
	OrigemSistema string
	RoleTipo      string

	OcorreuAs time.Time
	CriadoAs  time.Time
}
