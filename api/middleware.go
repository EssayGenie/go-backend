package api

import (
	"context"
	ghttp "go-backend/http"
	"net/http"
)

func (rs *Restful) loadConfig(w http.ResponseWriter, r *http.Request) (context.Context, error) {
	ctx := r.Context()
	ctx = ghttp.WithComponent(ctx, "api")
	return ctx, nil
}
