package http

import (
	"context"
	"log"
	"net/http"
	"runtime/debug"
)

// Recoverer is a middleware that recovers from panics, logs the panic (and a backtrace)
// then returns a HTTP 500 (Internal Server HTTPError) status if possible.
// Recoverer prints a request ID if one is provided.
func Recoverer(w http.ResponseWriter, r *http.Request) (context.Context, error) {
	defer func() {
		if healer := recover(); healer != nil {
			log.Printf("PANIC: %v", debug.Stack())

			he := &HTTPError{
				Code:    http.StatusInternalServerError,
				Message: http.StatusText(http.StatusInternalServerError),
			}

			handleError(he, w, r)
		}
	}()

	return nil, nil
}

func handleError(err error, w http.ResponseWriter, r *http.Request) {
	errorId := GetRequestId(r.Context())
	switch e := err.(type) {
	case *HTTPError:
		if e.Code >= http.StatusInternalServerError {
			log.Printf("HTTPError: %s", e.InternalError)
			e.ErrorID = errorId
			log.Println("ERROR :", e.Cause())
		} else {
			log.Println("INFO :", e.Cause())
		}
		if jsonErr := SendJSON(w, e.Code, e); jsonErr != nil {
			log.Println("ERROR :", jsonErr)
			handleError(jsonErr, w, r)
		}
	default:
		log.Println("UNHANDLED FATAL ERROR :", e)
		w.WriteHeader(http.StatusInternalServerError)
		if _, writeErr := w.Write([]byte(http.StatusText(http.StatusInternalServerError))); writeErr != nil {
			log.Println("ERROR WRITING GENERIC ERROR MESSAGE :", err)
		}
	}
}
