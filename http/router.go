package http

import (
	"github.com/go-chi/chi/v5"
	"go-backend/http/handler"
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

func (r *Router) Get(pattern string, fn handler.ApiHandler) {
	r.chi.Get(pattern, handler.Handler(fn))
}

func (r *Router) Post(pattern string, fn handler.ApiHandler) {
	r.chi.Post(pattern, handler.Handler(fn))
}
func (r *Router) Put(pattern string, fn handler.ApiHandler) {
	r.chi.Put(pattern, handler.Handler(fn))
}

func (r *Router) Delete(pattern string, fn handler.ApiHandler) {
	r.chi.Delete(pattern, handler.Handler(fn))
}

func (r *Router) With(fn handler.MiddlewareHandler) *Router {
	c := r.chi.With(handler.GetMiddleware(fn))
	return &Router{chi: c}
}

func (r *Router) WithBypass(fn func(next http.Handler) http.Handler) *Router {
	c := r.chi.With(fn)
	return &Router{chi: c}
}

func (r *Router) Use(fn handler.MiddlewareHandler) {
	r.chi.Use(handler.GetMiddleware(fn))
}

func (r *Router) UseBypass(fn func(next http.Handler) http.Handler) {
	r.chi.Use(fn)
}

func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	r.chi.ServeHTTP(w, req)
}
