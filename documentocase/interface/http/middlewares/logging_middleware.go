package middlewares

import (
	"log"
	"net/http"
	"time"
)

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		start := time.Now()

		traceID := GetTraceID(r)

		log.Printf(
			"[REQ] trace=%s method=%s path=%s",
			traceID,
			r.Method,
			r.URL.Path,
		)

		next.ServeHTTP(w, r)

		log.Printf(
			"[RES] trace=%s method=%s path=%s duration=%s",
			traceID,
			r.Method,
			r.URL.Path,
			time.Since(start),
		)
	})
}
