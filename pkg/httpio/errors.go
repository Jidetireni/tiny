package httpio

import (
	"fmt"
	"net/http"
)

// AppError is the shared error type used across all packages in tiny.
// Code maps directly to an HTTP status code, Message is what the caller sees.
type AppError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (e *AppError) Error() string {
	return fmt.Sprintf("(%d) %s", e.Code, e.Message)
}

// --- constructors -----------------------------------------------------------

func NewAppError(code int, message string) *AppError {
	return &AppError{Code: code, Message: message}
}

func BadRequest(message string) *AppError {
	return &AppError{Code: http.StatusBadRequest, Message: message}
}

func NotFound(message string) *AppError {
	return &AppError{Code: http.StatusNotFound, Message: message}
}

func Conflict(message string) *AppError {
	return &AppError{Code: http.StatusConflict, Message: message}
}

func UnprocessableEntity(message string) *AppError {
	return &AppError{Code: http.StatusUnprocessableEntity, Message: message}
}

func InternalServerError(message string) *AppError {
	return &AppError{Code: http.StatusInternalServerError, Message: message}
}
