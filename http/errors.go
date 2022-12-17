package http

import (
	"fmt"
	"log"
	"net/http"
)

type HTTPError struct {
	Code            int
	Message         string
	InternalMessage string `json:"-"`
	InternalError   error  `json:"-"`
	ErrorID         string `json:"error_id,omitempty"`
}

func NewHTTPError(code int, message string, args ...interface{}) *HTTPError {
	return &HTTPError{
		Code:            code,
		Message:         fmt.Sprintf(message, args...),
		InternalMessage: fmt.Sprintf(message, args...),
		InternalError:   fmt.Errorf(message, args...),
	}
}

func (e *HTTPError) Error() string {
	log.Printf("HTTPError: %s", e.InternalError)
	return e.Message
}

func (e *HTTPError) Cause() error {
	return e.InternalError
}

func InternalServerError(message string, args ...interface{}) *HTTPError {
	return NewHTTPError(http.StatusInternalServerError, message, args...)
}

func BadRequestError(message string, args ...interface{}) *HTTPError {
	return NewHTTPError(http.StatusBadRequest, message, args...)
}

func UnauthorizedError(message string, args ...interface{}) *HTTPError {
	return NewHTTPError(http.StatusUnauthorized, message, args...)
}

func NotFoundError(message string, args ...interface{}) *HTTPError {
	return NewHTTPError(http.StatusNotFound, message, args...)
}
