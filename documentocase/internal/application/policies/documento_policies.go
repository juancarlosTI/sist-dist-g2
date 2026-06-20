package policies

import (
	"github.com/juancarlosTI/monorepo-gestao-jur/backend/documentocase/internal/application/auth"
	"github.com/juancarlosTI/monorepo-gestao-jur/backend/shared/types"
)

func podeGerenciarDocumento(authCtx auth.AuthContext) bool {

	switch authCtx.Autor.Tipo {

	case types.AutorHumano,
		types.AutorAgente,
		types.AutorSistema:

		return true

	default:
		return false
	}
}

type DocumentoPolicy interface {
	PodeCriarDocumento(authCtx auth.AuthContext) bool
	PodeArquivarDocumento(authCtx auth.AuthContext) bool
	PodeValidarDocumento(authCtx auth.AuthContext) bool
	PodeInvalidarDocumento(authCtx auth.AuthContext) bool
	PodeAssociarDocumentoAoProcesso(authCtx auth.AuthContext) bool
	PodeAtualizarDocumento(authCtx auth.AuthContext) bool
}
