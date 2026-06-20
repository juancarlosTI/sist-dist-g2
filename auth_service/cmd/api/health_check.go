package main

import "net/http"

// @Summary Retorna status da API
// @Description Endpoint de healthcheck
// @Tags health
// @Success 200 {string} string "ok"
// @Router /health [get]
func HealthHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("ok"))
}
