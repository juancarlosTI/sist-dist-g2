package mappers

import (
	"net/http"

	"github.com/juancarlosTI/monorepo-gestao-jur/backend/auth_service/internal/application/auth"
)

func RequestContextFromRequest(r *http.Request) (auth.RequestContext, error) {
	return auth.RequestContext{
		RequestID:     r.Header.Get("X-Request-ID"),
		CorrelationID: r.Header.Get("X-Correlation-ID"),
	}, nil
}
