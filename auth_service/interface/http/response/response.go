package response

import (
	"encoding/json"
	"net/http"
)

type ErrorDetail struct {
	Field   string `json:"field,omitempty"`
	Message string `json:"message,omitempty"`
}

type ErrorBody struct {
	Code    string        `json:"code"`
	Message string        `json:"message"`
	Details []ErrorDetail `json:"details,omitempty"`
}

type ErrorResponse struct {
	Error ErrorBody `json:"error"`
}

type SuccessResponse struct {
	Data interface{} `json:"data"`
}

// ERROR RESPONSE
func WriteJSONError(
	w http.ResponseWriter,
	status int,
	message string,
	code string,
	details ...ErrorDetail,
) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	resp := ErrorResponse{
		Error: ErrorBody{
			Code:    code,
			Message: message,
			Details: details,
		},
	}

	_ = json.NewEncoder(w).Encode(resp)
}

// SUCCESS RESPONSE
func WriteJSON(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	_ = json.NewEncoder(w).Encode(SuccessResponse{
		Data: data,
	})
}

// (opcional) sucesso sem body (204 safe)
func WriteNoContent(w http.ResponseWriter) {
	w.WriteHeader(http.StatusNoContent)
}
