package commands

import "github.com/juancarlosTI/monorepo-gestao-jur/backend/documentocase/internal/dominio/common"

type AssociarDocumentoAoProcessoCommand struct {
	DocumentoID common.DocumentoID
	ProcessoID  common.ExternalProcessoID
}
