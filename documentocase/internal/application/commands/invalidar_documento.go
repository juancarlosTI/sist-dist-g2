package commands

import "github.com/juancarlosTI/monorepo-gestao-jur/backend/documentocase/internal/dominio/common"

type InvalidarDocumentoCommand struct {
	DocumentoID common.DocumentoID
	Motivo      string
}
