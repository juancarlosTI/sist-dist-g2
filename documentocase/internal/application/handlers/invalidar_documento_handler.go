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

type InvalidarDocumentoHandler struct {
	repo   documento.Repository
	policy policies.DocumentoPolicy
}

func NewInvalidarDocumentoHandler(repo documento.Repository, policy policies.DocumentoPolicy) *InvalidarDocumentoHandler {
	return &InvalidarDocumentoHandler{
		repo:   repo,
		policy: policy,
	}
}

func (h *InvalidarDocumentoHandler) Handle(cmd commands.InvalidarDocumentoCommand,
	authCtx auth.AuthContext, ctx context.Context) error {
	doc, err := h.repo.PorID(ctx, cmd.DocumentoID)
	if err != nil {
		return err
	}

	if !h.policy.PodeInvalidarDocumento(authCtx) {
		return application.ErrNaoAutorizado
	}

	evento_ctx := mappers.EventContextFromAuth(authCtx)

	if err := doc.InvalidarDocumento(cmd.Motivo, evento_ctx); err != nil {
		return err
	}

	// h.repo.Salvar(ctx, doc)

	return nil
}
