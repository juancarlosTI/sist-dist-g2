package commands

import (
	"github.com/juancarlosTI/monorepo-gestao-jur/backend/documentocase/internal/dominio/common"
	types "github.com/juancarlosTI/monorepo-gestao-jur/backend/shared/types"
)

type CriarDocumentoPorSolicitacaoCommand struct {
	PedidoDocumento *common.PedidoDocumento
	Autor           types.Autor
	Origem          types.Origem
}
