package errs

import "net/http"

type AppError struct {
	Message string
	Code    int
}

func (e *AppError) Error() string {
	return e.Message
}

func New(message string, code int) AppError {
	return AppError{
		Message: message,
		Code:    code,
	}
}

func NewBadRequestError(message string) AppError {
	return New(message, http.StatusBadRequest)
}

func NewNotFoundError(message string) AppError {
	return New(message, http.StatusNotFound)
}

func NewInternalServerError(message string) AppError {
	return New(message, http.StatusInternalServerError)
}

func NewUnauthorizedError(message string) AppError {
	return New(message, http.StatusUnauthorized)
}

func NewForbiddenError(message string) AppError {
	return New(message, http.StatusForbidden)
}

func NewConflictError(message string) AppError {
	return New(message, http.StatusConflict)
}

func NewBadGatewayError(message string) AppError {
	return New(message, http.StatusBadGateway)
}

func NewServiceUnavailableError(message string) AppError {
	return New(message, http.StatusServiceUnavailable)
}
