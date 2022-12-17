package handler

import (
	"context"
	ghttp "go-backend/http"
	"net/http"
)

type MiddlewareHandler func(w http.ResponseWriter, r *http.Request) (context.Context, error)

func (mh MiddlewareHandler) serve(next http.Handler, w http.ResponseWriter, r *http.Request) {
	ctx, err := mh(w, r)
	if err != nil {
		// handling error
		ghttp.HandleError(err, w, r)
		return
	}
	if ctx != nil {
		r = r.WithContext(ctx)
	}
	next.ServeHTTP(w, r)
}

func (mh MiddlewareHandler) middlewareHandlerImpl(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		mh.serve(next, w, r)
	})
}

func GetMiddleware(fn MiddlewareHandler) func(http.Handler) http.Handler {
	return fn.middlewareHandlerImpl
}
