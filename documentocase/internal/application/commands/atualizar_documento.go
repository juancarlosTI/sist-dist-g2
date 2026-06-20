package commands

import "github.com/juancarlosTI/monorepo-gestao-jur/backend/documentocase/internal/dominio/common"

type AtualizarDocumentoCommand struct {
	DocumentoID common.DocumentoID
}
