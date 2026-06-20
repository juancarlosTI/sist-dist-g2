package common

import "github.com/juancarlosTI/monorepo-gestao-jur/backend/shared/types"

type EventContext struct {
	CorrelacaoID  string
	CausalidadeID string
	Origem        types.Origem
	Autor         types.Autor
	Role          types.Role
}
