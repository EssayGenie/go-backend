package http

import (
	"github.com/go-chi/chi/v5"
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
	r.chi.Route(pattern, func(r chi.Router) {
		fn(&Router{chi: r})
	})
}
