package middlewares

import (
	"log"
	"net/http"
	"runtime/debug"
)

func RecoveryMiddleware(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		defer func() {

			if err := recover(); err != nil {

				log.Printf(
					"[PANIC] path=%s err=%v\n%s",
					r.URL.Path,
					err,
					string(debug.Stack()),
				)

				http.Error(
					w,
					"internal server error",
					http.StatusInternalServerError,
				)
			}
		}()

		next.ServeHTTP(w, r)
	})
}
