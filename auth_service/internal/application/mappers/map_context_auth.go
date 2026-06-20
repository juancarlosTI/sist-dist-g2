package mappers

import (
	"github.com/juancarlosTI/monorepo-gestao-jur/backend/auth_service/internal/application/auth"
	shared_context "github.com/juancarlosTI/monorepo-gestao-jur/backend/shared/common"
	shared_types "github.com/juancarlosTI/monorepo-gestao-jur/backend/shared/types"
)

func EventContextFromAuth(
	authCtx auth.AuthContext,
) shared_context.EventContext {

	return shared_context.EventContext{
		Origem: shared_types.Origem{
			Canal:   MapCanal(string(authCtx.Origem.Canal)),
			Sistema: MapSistema(string(authCtx.Origem.Sistema)),
		},
		Autor: shared_types.Autor{
			Tipo: MapAutorTipo(string(authCtx.Autor.Tipo)),
			ID:   authCtx.Autor.ID,
		},
	}
}
