package middlewares

import (
	"context"
	"net/http"

	"github.com/juancarlosTI/monorepo-gestao-jur/backend/auth_service/internal/dominio/token_access"
	authtokenjwt "github.com/juancarlosTI/monorepo-gestao-jur/backend/auth_service/internal/infraestrutura/seguranca/auth-token-jwt"
)

type TokenValidator interface {
	ValidateToken(token string) (*token_access.AccessToken, error)
}

func GetToken(r *http.Request) (string, bool) {
	token, ok := r.Context().Value(TokenContextKey).(string)
	return token, ok
}

func TokenValidateMiddleware(validator authtokenjwt.JWTAccessTokenService) func(http.Handler) http.Handler {

	return func(next http.Handler) http.Handler {

		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			tokenStr, ok := GetToken(r)
			if !ok {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}

			claims, err := validator.Validate(ctx, tokenStr)
			if err != nil {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}

			ctx = context.WithValue(ctx, TokenClaimsKey, claims)

			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
