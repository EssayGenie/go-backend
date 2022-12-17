package http

import (
	"net/http"
)

type ApiHandler func(w http.ResponseWriter, r *http.Request) error

func (ah ApiHandler) serve(w http.ResponseWriter, r *http.Request) {
	if err := ah(w, r); err != nil {
		// handling error
		HandleError(err, w, r)
	}
}

func Handler(fn ApiHandler) http.HandlerFunc {
	return fn.serve
}
