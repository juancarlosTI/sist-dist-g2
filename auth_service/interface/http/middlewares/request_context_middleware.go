package middlewares

import (
	"context"
	"net/http"

	"github.com/google/uuid"
	"github.com/juancarlosTI/monorepo-gestao-jur/backend/auth_service/internal/application/auth"
)

func GetRequestContext(r *http.Request) (auth.RequestContext, bool) {
	ctx := r.Context().Value(RequestContextKey)

	if ctx == nil {
		return auth.RequestContext{}, false
	}

	rc, ok := ctx.(auth.RequestContext)
	return rc, ok
}

const RequestContextKey contextKey = "request_context"

func RequestContextMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		reqID := r.Header.Get("X-Request-ID")
		if reqID == "" {
			reqID = uuid.NewString()
		}

		correlationID := r.Header.Get("X-Correlation-ID")
		if correlationID == "" {
			correlationID = reqID
		}

		ctx := auth.RequestContext{
			RequestID:     reqID,
			CorrelationID: correlationID,
		}

		newCtx := context.WithValue(r.Context(), RequestContextKey, ctx)

		next.ServeHTTP(w, r.WithContext(newCtx))
	})
}
