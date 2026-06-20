package mappers

import (
	"github.com/juancarlosTI/monorepo-gestao-jur/backend/auth_service/internal/application/auth"
	shared_context "github.com/juancarlosTI/monorepo-gestao-jur/backend/shared/common"
	"github.com/juancarlosTI/monorepo-gestao-jur/backend/shared/types"
)

func EventContextFromRequest(rc auth.RequestContext, opts ...string) shared_context.EventContext {

	roleStr := string(types.RoleProfissional) // default

	if len(opts) > 0 {
		roleStr = opts[0]
	}

	map_account_type := StringToRole(roleStr)
	return shared_context.EventContext{
		CorrelacaoID:  rc.CorrelationID,
		CausalidadeID: rc.RequestID,
		Autor: types.Autor{
			Tipo: types.AutorHumano,
			ID:   "UUID", // importante: usuário ainda não existe
		},
		Origem: types.Origem{
			Canal:   types.CanalAPI,
			Sistema: types.Auth,
		},
		Role: types.Role{
			Tipo: map_account_type,
		},
	}
}
