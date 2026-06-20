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

type AssociarDocumentoAoProcessoHandler struct {
	repo   documento.Repository
	policy policies.DocumentoPolicy
}

func NewAssociarDocumentoAoProcessoHandler(repo documento.Repository, policy policies.DocumentoPolicy) *AssociarDocumentoAoProcessoHandler {
	return &AssociarDocumentoAoProcessoHandler{
		repo:   repo,
		policy: policy,
	}
}

func (h *AssociarDocumentoAoProcessoHandler) Handle(cmd commands.AssociarDocumentoAoProcessoCommand,
	authCtx auth.AuthContext, ctx context.Context) error {
	doc, err := h.repo.PorID(ctx, cmd.DocumentoID)
	if err != nil {
		return err
	}

	if !h.policy.PodeAssociarDocumentoAoProcesso(authCtx) {
		return application.ErrNaoAutorizado
	}

	evento_ctx := mappers.EventContextFromAuth(authCtx)

	if err := doc.AssociarDocumentoAoProcesso(cmd.ProcessoID,
		evento_ctx); err != nil {
		return err
	}

	// h.repo.Salvar(ctx, doc)

	return nil
}
