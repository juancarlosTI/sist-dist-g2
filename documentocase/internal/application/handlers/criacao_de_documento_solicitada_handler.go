package handlers

import (
	"context"

	shared "github.com/juancarlosTI/monorepo-gestao-jur/backend/shared/kernel/processo"
)

type DocumentoCriacaoSolicitadaConsumer struct {
	criarDocumento *CriarDocumentoHandler
}

func NewDocumentoCriacaoSolicitadaConsumer(
	criarDocumento *CriarDocumentoHandler,
) *DocumentoCriacaoSolicitadaConsumer {
	return &DocumentoCriacaoSolicitadaConsumer{
		criarDocumento: criarDocumento,
	}
}

func (h *DocumentoCriacaoSolicitadaConsumer) Handle(
	ctx context.Context,
	event shared.DocumentoCriacaoSolicitada,
) error {

	// acl_map, err := mappers.MapCriacaoDocumento(event)
	// if err != nil {
	// 	return err
	// }

	// authCtx := auth.AuthContext{
	// 	Autor:  event.Base().Autor,
	// 	Roles:  event.Base().Role,
	// 	Origem: event.Base().Origem,
	// }

	// h.criarDocumento.Handle(acl_map, authCtx, ctx)

	return nil
}
