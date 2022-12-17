package http

import (
	"context"
	"log"
	"net/http"
	"runtime/debug"
)

// Recoverer is a middleware that recovers from panics, logs the panic (and a
// backtrace), and returns a HTTP 500 (Internal Server Error) status if
// possible. Recoverer prints a request ID if one is provided.
func Recoverer(w http.ResponseWriter, r *http.Request) (context.Context, error) {
	defer func() {
		if healer := recover(); healer != nil {
			log.Printf("PANIC: %v", debug.Stack())

			he := &Error{
				Code:    http.StatusInternalServerError,
				Message: http.StatusText(http.StatusInternalServerError),
			}

			handleError(he, w, r)
		}
	}()

	return nil, nil
}

func handleError(err *Error, w http.ResponseWriter, r *http.Request) {
	// write more...
}
