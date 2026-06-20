package common

import (
	"time"

	types "github.com/juancarlosTI/monorepo-gestao-jur/backend/shared/types"
)

type EventoBase struct {
	EventoID types.EventoID
	Versao   int

	Hash         string
	AgregadoID   string
	AgregadoTipo string

	OcorreuAs     time.Time
	CorrelacaoID  string
	CausalidadeID string

	Origem types.Origem
	Autor  types.Autor
	Role   types.Role
}
