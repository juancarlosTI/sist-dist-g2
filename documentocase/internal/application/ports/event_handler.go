package ports

import kernel "github.com/juancarlosTI/monorepo-gestao-jur/backend/shared/kernel/processo"

type DocumentoCriacaoSolicitadaHandler interface {
	Handle(event kernel.DocumentoCriacaoSolicitada) error
}

type DocumentoAssociacaoSolicitada interface {
	Handle(event kernel.DocumentoAssociacaoSolicitada) error
}
