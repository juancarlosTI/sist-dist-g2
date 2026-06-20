package middlewares

import (
	"encoding/json"
	"net/http"

	appErr "github.com/juancarlosTI/monorepo-gestao-jur/backend/shared/errors"
)

func HandleError(w http.ResponseWriter, err error) {

	if err == nil {
		return
	}

	if e, ok := err.(*appErr.AppError); ok {

		writeJSON(w, e.Status, e)

		return
	}

	writeJSON(
		w,
		http.StatusInternalServerError,
		map[string]string{
			"code":    "internal_error",
			"message": "internal server error",
		},
	)
}

func writeJSON(
	w http.ResponseWriter,
	status int,
	body interface{},
) {

	w.Header().
		Set("Content-Type", "application/json")

	w.WriteHeader(status)

	_ = json.NewEncoder(w).
		Encode(body)
}
