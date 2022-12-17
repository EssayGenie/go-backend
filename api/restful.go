package api

import (
	"context"
	"fmt"
	"go-backend/conf"
	ghttp "go-backend/http"
	"log"
	"net/http"
)

type Restful struct {
	config  *conf.Configuration
	handler http.Handler
}

func (rs *Restful) ListenAndServe(host string, port uint32) {
	server := &http.Server{
		Addr:    fmt.Sprintf("%s:%d", host, port),
		Handler: rs.handler,
	}

	done := make(chan struct{})
	defer close(done)

	go func() {
		ghttp.Terminate(done)
		ctx, cancel := context.WithTimeout(context.Background(), conf.TIMEOUT_IN_MINUTE)
		defer cancel()
		err := server.Shutdown(ctx)
		if err != nil {
			log.Fatal("Shutdown Error :::", err)
			return
		}
	}()
}

func RegisterEndpoints(gc *conf.Configuration) *Restful {
	rs := &Restful{
		config: gc,
	}
	r := ghttp.NewRouter()

	// registering middleware
	r.Use(rs.loadConfig)
	r.Use(ghttp.AddRequestId)
	r.Use(ghttp.Recoverer)

	// registering endpoint
	r.Route(V1_ENDPOINT, func(r *ghttp.Router) {
		r.Get(HEALTH_ENDPOINT, rs.HealthCheck)
	})

	// allocating handler
	rs.handler = r

	return rs
}
