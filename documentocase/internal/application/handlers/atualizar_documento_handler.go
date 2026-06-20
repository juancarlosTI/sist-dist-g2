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

type AtualizarDocumentoHandler struct {
	repo   documento.Repository
	policy policies.DocumentoPolicy
}

func NewAtualizarDocumentoHandler(repo documento.Repository,
	policy policies.DocumentoPolicy) *AtualizarDocumentoHandler {
	return &AtualizarDocumentoHandler{
		repo:   repo,
		policy: policy,
	}
}

func (h *AtualizarDocumentoHandler) Handle(cmd commands.AtualizarDocumentoCommand,
	authCtx auth.AuthContext, ctx context.Context) error {
	doc, err := h.repo.PorID(ctx, cmd.DocumentoID)
	if err != nil {
		return err
	}

	if !h.policy.PodeAtualizarDocumento(authCtx) {
		return application.ErrNaoAutorizado
	}

	evento_ctx := mappers.EventContextFromAuth(authCtx)

	if err := doc.AtualizarConteudo(evento_ctx); err != nil {
		return err
	}

	// h.repo.Salvar(ctx, doc)

	return nil
}
