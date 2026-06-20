package handlers

import (
	"context"
	"fmt"
	"log"

	"github.com/juancarlosTI/monorepo-gestao-jur/backend/auth_service/internal/application/auth"
	dto "github.com/juancarlosTI/monorepo-gestao-jur/backend/auth_service/internal/application/dto"
	"github.com/juancarlosTI/monorepo-gestao-jur/backend/auth_service/internal/dominio/identidade"
)

type MeHandler struct {
	userRepo identidade.UserRepository
}

func NewMeHandler(
	userRepo identidade.UserRepository,
) *MeHandler {
	return &MeHandler{
		userRepo: userRepo,
	}
}

func (h *MeHandler) Handle(
	ctx context.Context,
	authCtx auth.AuthContext) (*dto.MeResponseDTO, error) {

	user, err := h.userRepo.FindByID(ctx, authCtx.Autor.ID)
	if err != nil {
		return nil, err
	}

	log.Println("\n\nUser:", user)
	log.Println("\n\nErr: ", err)

	if user == nil {
		return nil, fmt.Errorf("invalid credentials")
	}

	return &dto.MeResponseDTO{
		ID:    user.ID(),
		Email: user.Email(),
		Nome:  user.Nome(),
		Roles: user.Role(),
	}, nil
}
