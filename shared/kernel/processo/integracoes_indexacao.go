package processo

import (
	shared_context "github.com/juancarlosTI/monorepo-gestao-jur/backend/shared/kernel/eventos"
)

const ProcessoCriadoIndexVersao = 1

type ProcessoCriadoIndexPayload struct {
	ProcessoID string `json:"processo_id"`

	Nome string `json:"nome"`

	CategoriaProcesso string `json:"categoria_processo"`

	Estado      int    `json:"estado"`
	EstadoLabel string `json:"estado_label"`

	ClienteID   *string `json:"cliente_id"`
	ClienteNome string  `json:"cliente_nome"`

	AdvogadoID   string `json:"advogado_id"`
	AdvogadoNome string `json:"advogado_nome"`
}

type ProcessoCriadoIndex struct {
	BaseEvento    shared_context.EventoIntegracaoBase
	PayloadEvento ProcessoCriadoIndexPayload
}

func (e ProcessoCriadoIndex) Nome() string {
	return "ProcessoCriadoIndex"
}

func (e ProcessoCriadoIndex) Base() shared_context.EventoIntegracaoBase {
	return e.BaseEvento
}

func (e ProcessoCriadoIndex) Payload() ProcessoCriadoIndexPayload {
	return e.PayloadEvento
}

// const DocumentoCriacaoSolicitadaVersao = 1

// type DocumentoCriacaoSolicitada struct {
// 	BaseEvento    shared_context.EventoIntegracaoBase
// 	PayloadEvento any
// }

// func (e DocumentoCriacaoSolicitada) Nome() string {
// 	return "DocumentoCriacaoSolicitada"
// }

// func (e DocumentoCriacaoSolicitada) Base() shared_context.EventoIntegracaoBase {
// 	return e.BaseEvento
// }

// func (e DocumentoCriacaoSolicitada) Payload() any {
// 	return e.PayloadEvento
// }
