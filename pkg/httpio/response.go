package httpio

import (
	"encoding/json"
	"errors"
	"net/http"
)

// WriteJSON writes a JSON response with the given status code and payload.
func WriteJSON(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

// WriteError inspects the error type and responds accordingly:
//   - *AppError  → uses its Code and Message
//   - plain error → 500 Internal Server Error, internals are never leaked to the caller
func WriteError(w http.ResponseWriter, err error) {
	var appErr *AppError
	if errors.As(err, &appErr) {
		WriteJSON(w, appErr.Code, appErr)
		return
	}

	// Raw / unexpected error — never leak internals to the caller.
	WriteJSON(w, http.StatusInternalServerError, &AppError{
		Code:    http.StatusInternalServerError,
		Message: "internal server error",
	})
}
