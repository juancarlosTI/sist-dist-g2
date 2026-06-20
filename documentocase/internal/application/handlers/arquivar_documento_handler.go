package handlers

import (
	"context"

	"github.com/juancarlosTI/monorepo-gestao-jur/backend/documentocase/internal/application/auth"
	"github.com/juancarlosTI/monorepo-gestao-jur/backend/documentocase/internal/application/commands"
	application "github.com/juancarlosTI/monorepo-gestao-jur/backend/documentocase/internal/application/common"
	"github.com/juancarlosTI/monorepo-gestao-jur/backend/documentocase/internal/application/mappers"
	"github.com/juancarlosTI/monorepo-gestao-jur/backend/documentocase/internal/application/policies"
	"github.com/juancarlosTI/monorepo-gestao-jur/backend/documentocase/internal/dominio/documento"
)

type ArquivarDocumentoHandler struct {
	repo   documento.Repository
	policy policies.DocumentoPolicy
}

func NewArquivarDocumentoHandler(repo documento.Repository, policy policies.DocumentoPolicy) *ArquivarDocumentoHandler {
	return &ArquivarDocumentoHandler{
		repo:   repo,
		policy: policy,
	}
}

func (h *ArquivarDocumentoHandler) Handle(
	authCtx auth.AuthContext,
	ctx context.Context,
	cmd commands.ArquivarDocumentoCommand,
) error {

	doc, err := h.repo.PorID(ctx, cmd.DocumentoID)
	if err != nil {
		return err
	}

	if !h.policy.PodeArquivarDocumento(authCtx) {
		return application.ErrNaoAutorizado
	}

	evento_ctx := mappers.EventContextFromAuth(authCtx)

	if err := doc.ArquivarDocumento(evento_ctx); err != nil {
		return err
	}
	// h.repo.Salvar(ctx, doc)

	return nil
}
