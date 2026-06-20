package identidade

import (
	dominio_common "github.com/juancarlosTI/monorepo-gestao-jur/backend/auth_service/internal/dominio/common"
	shared_types "github.com/juancarlosTI/monorepo-gestao-jur/backend/shared/types"
)

type EventoAuth interface {
	Nome() string
	Base() dominio_common.EventoBase
	Payload() any
}

type UsuarioCriadoPayload struct {
	UserID dominio_common.UserID `json:"user_id"`
	Email  string                `json:"user_email"`
	Nome   string                `json:"user_nome"`
	Roles  shared_types.Role     `json:"user_roles"`
}

type UsuarioCriado struct {
	base dominio_common.EventoBase

	payload UsuarioCriadoPayload
}

func (e UsuarioCriado) Nome() string {
	return "UsuarioCriado"
}

func (e UsuarioCriado) Base() dominio_common.EventoBase {
	return e.base
}

func (e UsuarioCriado) Payload() any {
	return e.payload
}
