package middlewares

import (
	"context"
	"log"
	"net/http"

	"github.com/juancarlosTI/monorepo-gestao-jur/backend/documentocase/interface/http/mappers"
)

type AuthContextKeyType struct{}

var AuthContextKey = AuthContextKeyType{}

func AuthMiddleware(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		authCtx, err := mappers.AuthContextFromRequest(r)
		if err != nil {

			log.Printf(
				"[AUTH] unauthorized path=%s err=%v",
				r.URL.Path,
				err,
			)

			http.Error(
				w,
				"unauthorized",
				http.StatusUnauthorized,
			)

			return
		}

		ctx := context.WithValue(
			r.Context(),
			AuthContextKey,
			authCtx,
		)

		next.ServeHTTP(
			w,
			r.WithContext(ctx),
		)
	})
}
