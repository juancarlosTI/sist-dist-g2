package handlers

import (
	"context"
	"database/sql"

	"github.com/juancarlosTI/monorepo-gestao-jur/backend/documentocase/internal/application/auth"
	"github.com/juancarlosTI/monorepo-gestao-jur/backend/documentocase/internal/application/commands"
	application "github.com/juancarlosTI/monorepo-gestao-jur/backend/documentocase/internal/application/common"
	"github.com/juancarlosTI/monorepo-gestao-jur/backend/documentocase/internal/application/mappers"
	"github.com/juancarlosTI/monorepo-gestao-jur/backend/documentocase/internal/application/policies"
	"github.com/juancarlosTI/monorepo-gestao-jur/backend/documentocase/internal/application/ports"
	"github.com/juancarlosTI/monorepo-gestao-jur/backend/documentocase/internal/dominio/documento"
)

type CriarDocumentoHandler struct {
	uow    application.UnitOfWork
	repo   documento.Repository
	outbox ports.OutboxRepository
	policy policies.DocumentoPolicy
}

func NewCriarDocumentoHandler(
	uow application.UnitOfWork,
	outboxRepo ports.OutboxRepository,
	repo documento.Repository,
	policy policies.DocumentoPolicy,
) *CriarDocumentoHandler {
	return &CriarDocumentoHandler{
		uow:    uow,
		repo:   repo,
		outbox: outboxRepo,
		policy: policy,
	}
}

func (h *CriarDocumentoHandler) Handle(
	cmd commands.CriarDocumentoCommand,
	authCtx auth.AuthContext,
	ctx context.Context,
) error {

	if !h.policy.PodeCriarDocumento(authCtx) {
		return application.ErrNaoAutorizado
	}

	eventoCtx := mappers.EventContextFromAuth(authCtx)

	doc, err := documento.CriarDocumento(
		eventoCtx,
		cmd.ArquivoID,
	)
	if err != nil {
		return err
	}

	return h.uow.Do(ctx, func(tx *sql.Tx) error {

		events, err := h.repo.Salvar(
			ctx,
			tx,
			doc,
		)
		if err != nil {
			return err
		}

		for _, event := range events {

			msg, err := mappers.MapToOutbox(event)
			if err != nil {
				return err
			}

			err = h.outbox.Salvar(
				ctx,
				tx,
				event,
				msg,
			)
			if err != nil {
				return err
			}
		}

		return nil
	})
}
