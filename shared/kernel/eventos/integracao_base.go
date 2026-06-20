package eventos

import (
	"time"

	"github.com/juancarlosTI/monorepo-gestao-jur/backend/shared/types"
)

type EventoIntegracaoBase struct {
	EventoID      types.EventoID
	CorrelacaoID  string
	CausalidadeID string
	RoutingKey    string
	Versao        int
	Autor         types.Autor
	Origem        types.Origem
	Role          types.Role
	OcorreuAs     time.Time
	CriadoAs      time.Time
}

type EventoIntegracao interface {
	Nome() string
	Base() EventoIntegracaoBase
	Payload() any
}
