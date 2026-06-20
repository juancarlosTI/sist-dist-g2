package middlewares

import (
	"net/http"

	"github.com/juancarlosTI/monorepo-gestao-jur/backend/documentocase/internal/application/auth"
)

func GetAuthContext(r *http.Request) (auth.AuthContext, bool) {

	authCtx, ok := r.Context().
		Value(AuthContextKey).(auth.AuthContext)

	return authCtx, ok
}

func MustGetAuthContext(r *http.Request) auth.AuthContext {

	authCtx, ok := GetAuthContext(r)
	if !ok {
		panic("auth context missing")
	}

	return authCtx
}
