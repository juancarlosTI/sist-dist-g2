package mappers

import (
	"net/http"

	"github.com/juancarlosTI/monorepo-gestao-jur/backend/documentocase/internal/application/auth"
	shared_types "github.com/juancarlosTI/monorepo-gestao-jur/backend/shared/types"
)

func AuthContextFromRequest(r *http.Request) (auth.AuthContext, error) {

	roleTipo, err := shared_types.ParseRoleTipo(
		r.Header.Get("X-User-Role"),
	)
	if err != nil {
		return auth.AuthContext{}, err
	}

	return auth.AuthContext{
		Autor: shared_types.Autor{
			ID:   r.Header.Get("X-User-ID"),
			Tipo: shared_types.AutorHumano,
		},
		Roles: shared_types.Role{
			Tipo: roleTipo,
		},
		Origem: shared_types.Origem{
			Canal:   shared_types.CanalAPI,
			Sistema: shared_types.Legal,
		},
	}, nil
}

// func EventoBaseFrom(eventCtx shared_context.EventContext) shared_evento.EventoBase {
// 	return shared_evento.EventoBase{
// 		UserID:    r.Header.Get("X-User-Id"),
// 		ActorType: r.Header.Get("X-Actor-Type"),
// 		Roles:     parseRoles(r.Header.Get("X-Roles")),
// 		Canal:     "api",
// 		Sistema:   "bc-legal",
// 	}
// }
