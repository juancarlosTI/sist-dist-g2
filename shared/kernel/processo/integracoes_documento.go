package processo

import shared_context "github.com/juancarlosTI/monorepo-gestao-jur/backend/shared/kernel/eventos"

const DocumentoAssociacaoSolicitadaVersao = 1

type DocumentoAssociacaoSolicitada struct {
	BaseEvento    shared_context.EventoIntegracaoBase
	PayloadEvento any
}

func (e DocumentoAssociacaoSolicitada) Nome() string {
	return "DocumentoAssociacaoSolicitada"
}

func (e DocumentoAssociacaoSolicitada) Base() shared_context.EventoIntegracaoBase {
	return e.BaseEvento
}

func (e DocumentoAssociacaoSolicitada) Payload() any {
	return e.PayloadEvento
}

const DocumentoCriacaoSolicitadaVersao = 1

type DocumentoCriacaoSolicitada struct {
	BaseEvento    shared_context.EventoIntegracaoBase
	PayloadEvento any
}

func (e DocumentoCriacaoSolicitada) Nome() string {
	return "DocumentoCriacaoSolicitada"
}

func (e DocumentoCriacaoSolicitada) Base() shared_context.EventoIntegracaoBase {
	return e.BaseEvento
}

func (e DocumentoCriacaoSolicitada) Payload() any {
	return e.PayloadEvento
}
