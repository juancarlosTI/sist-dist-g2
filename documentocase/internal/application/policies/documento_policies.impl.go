package policies

import (
	"github.com/juancarlosTI/monorepo-gestao-jur/backend/documentocase/internal/application/auth"
)

type DocumentoPolicyImpl struct{}

func (p DocumentoPolicyImpl) PodeCriarDocumento(authCtx auth.AuthContext) bool {
	return podeGerenciarDocumento(authCtx)
}

func (p DocumentoPolicyImpl) PodeArquivarDocumento(authCtx auth.AuthContext) bool {
	return podeGerenciarDocumento(authCtx)
}

func (p DocumentoPolicyImpl) PodeValidarDocumento(authCtx auth.AuthContext) bool {
	return podeGerenciarDocumento(authCtx)
}

func (p DocumentoPolicyImpl) PodeInvalidarDocumento(authCtx auth.AuthContext) bool {
	return podeGerenciarDocumento(authCtx)
}

func (p DocumentoPolicyImpl) PodeAssociarDocumentoAoProcesso(authCtx auth.AuthContext) bool {
	return podeGerenciarDocumento(authCtx)
}

func (p DocumentoPolicyImpl) PodeAtualizarDocumento(authCtx auth.AuthContext) bool {
	return podeGerenciarDocumento(authCtx)
}
