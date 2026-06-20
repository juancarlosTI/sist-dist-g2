package middlewares

import (
	"context"
	"net/http"
	"strings"
)

type contextKey string

const TokenContextKey contextKey = "access_token"

func TokenExtractorMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		authHeader := r.Header.Get("Authorization")

		if authHeader != "" && strings.HasPrefix(authHeader, "Bearer ") {
			token := strings.TrimPrefix(authHeader, "Bearer ")

			ctx := context.WithValue(r.Context(), TokenContextKey, token)
			r = r.WithContext(ctx)
		}

		next.ServeHTTP(w, r)
	})
}
