package handler

import (
	ghttp "go-backend/http"
	"net/http"
)

type apiHandler func(w http.ResponseWriter, r *http.Request) error

func (ah apiHandler) serve(w http.ResponseWriter, r *http.Request) {
	if err := ah(w, r); err != nil {
		// handling error
		ghttp.HandleError(err, w, r)
	}
}

func handler(fn apiHandler) http.HandlerFunc {
	return fn.serve
}
