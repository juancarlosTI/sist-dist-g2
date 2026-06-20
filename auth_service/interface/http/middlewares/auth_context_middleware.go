package middlewares

import (
	"context"
	"net/http"

	"github.com/juancarlosTI/monorepo-gestao-jur/backend/auth_service/internal/application/auth"
	"github.com/juancarlosTI/monorepo-gestao-jur/backend/auth_service/internal/dominio/token_access"
	"github.com/juancarlosTI/monorepo-gestao-jur/backend/shared/types"
)

const AuthContextKey contextKey = "auth_context"
const TokenClaimsKey contextKey = "token_claims"

// func mergeRoles(roles types.Role) []types.RoleTipo {
// 	if len(roles.Tipo) == 0 {
// 		return []types.RoleTipo{}
// 	}
// 	return roles.Tipo
// }

func AuthContextMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		claimsVal := r.Context().Value(TokenClaimsKey)
		if claimsVal == nil {
			http.Error(w, "unauthorized", http.StatusUnauthorized)
			return
		}

		at, ok := claimsVal.(*token_access.AccessToken)
		if !ok {
			http.Error(w, "invalid token claims", http.StatusUnauthorized)
			return
		}

		authCtx := auth.AuthContext{
			Autor: at.Autor,
			Roles: types.Role{
				Tipo: at.Roles.Tipo,
			},
			Origem: at.Origem,
		}

		ctx := context.WithValue(r.Context(), AuthContextKey, authCtx)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func GetAuthContext(r *http.Request) (auth.AuthContext, bool) {
	ctx := r.Context().Value(AuthContextKey)

	if ctx == nil {
		return auth.AuthContext{}, false
	}

	authCtx, ok := ctx.(auth.AuthContext)
	return authCtx, ok
}
