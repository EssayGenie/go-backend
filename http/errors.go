package http

import (
	"fmt"
	"log"
	"net/http"
)

type Error struct {
	Code            int
	Message         string
	InternalMessage string `json:"-"`
	InternalError   error  `json:"-"`
}

func NewError(code int, message string, args ...interface{}) *Error {
	return &Error{
		Code:            code,
		Message:         fmt.Sprintf(message, args...),
		InternalMessage: fmt.Sprintf(message, args...),
		InternalError:   fmt.Errorf(message, args...),
	}
}

func (e *Error) Error() error {
	log.Printf("Error: %s", e.InternalError)
	return e.InternalError
}

func (e *Error) ErrorMessage() string {
	return e.Message
}

func InternalServerError(message string, args ...interface{}) *Error {
	return NewError(http.StatusInternalServerError, message, args...)
}

func BadRequestError(message string, args ...interface{}) *Error {
	return NewError(http.StatusBadRequest, message, args...)
}

func UnauthorizedError(message string, args ...interface{}) *Error {
	return NewError(http.StatusUnauthorized, message, args...)
}

func NotFoundError(message string, args ...interface{}) *Error {
	return NewError(http.StatusNotFound, message, args...)
}
