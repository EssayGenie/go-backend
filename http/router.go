package http

import (
	"github.com/go-chi/chi/v5"
	"net/http"
)

type Router struct {
	chi chi.Router
}

func NewRouter() *Router {
	return &Router{
		chi: chi.NewRouter(),
	}
}

func (r *Router) Route(pattern string, fn func(*Router)) {
	r.chi.Route(pattern, func(c chi.Router) {
		fn(&Router{chi: c})
	})
}

func (r *Router) Get(pattern string, fn ApiHandler) {
	r.chi.Get(pattern, Handler(fn))
}

func (r *Router) Post(pattern string, fn ApiHandler) {
	r.chi.Post(pattern, Handler(fn))
}
func (r *Router) Put(pattern string, fn ApiHandler) {
	r.chi.Put(pattern, Handler(fn))
}

func (r *Router) Delete(pattern string, fn ApiHandler) {
	r.chi.Delete(pattern, Handler(fn))
}

func (r *Router) With(fn MiddlewareHandler) *Router {
	c := r.chi.With(GetMiddleware(fn))
	return &Router{chi: c}
}

func (r *Router) WithBypass(fn func(next http.Handler) http.Handler) *Router {
	c := r.chi.With(fn)
	return &Router{chi: c}
}

func (r *Router) Use(fn MiddlewareHandler) {
	r.chi.Use(GetMiddleware(fn))
}

func (r *Router) UseBypass(fn func(next http.Handler) http.Handler) {
	r.chi.Use(fn)
}

func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	r.chi.ServeHTTP(w, req)
}
