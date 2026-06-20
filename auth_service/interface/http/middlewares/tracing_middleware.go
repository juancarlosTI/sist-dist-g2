package middlewares

import (
	"context"
	"net/http"

	"github.com/google/uuid"
)

type Trace struct {
	TraceID string
	SpanID  string
}

const TraceKey contextKey = "trace"

func TracingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		trace := Trace{
			TraceID: uuid.NewString(),
			SpanID:  uuid.NewString(),
		}

		ctx := context.WithValue(r.Context(), TraceKey, trace)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
