package middlewares

import (
	"encoding/json"
	"net/http"
	"time"
)

type LogEntry struct {
	Timestamp     time.Time `json:"timestamp"`
	Method        string    `json:"method"`
	Path          string    `json:"path"`
	Status        int       `json:"status"`
	DurationMs    int64     `json:"duration_ms"`
	RequestID     string    `json:"request_id,omitempty"`
	CorrelationID string    `json:"correlation_id,omitempty"`
	UserID        string    `json:"user_id,omitempty"`
}

type responseWriter struct {
	http.ResponseWriter
	status int
}

func (rw *responseWriter) WriteHeader(code int) {
	rw.status = code
	rw.ResponseWriter.WriteHeader(code)
}

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		start := time.Now()

		rw := &responseWriter{ResponseWriter: w, status: 200}

		next.ServeHTTP(rw, r)

		duration := time.Since(start)

		// pega contextos
		reqCtx, _ := GetRequestContext(r)
		authCtx, _ := GetAuthContext(r)

		entry := LogEntry{
			Timestamp:     time.Now(),
			Method:        r.Method,
			Path:          r.URL.Path,
			Status:        rw.status,
			DurationMs:    duration.Milliseconds(),
			RequestID:     reqCtx.RequestID,
			CorrelationID: reqCtx.CorrelationID,
			UserID:        authCtx.Autor.ID,
		}

		jsonLog, _ := json.Marshal(entry)
		println(string(jsonLog)) // depois pluga em stdout → ELK / Loki
	})
}
