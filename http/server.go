package http

import (
	"context"
	"github.com/jaevor/go-nanoid"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func Termination(done <-chan struct{}) {
	signalChan := make(chan os.Signal)
	signal.Notify(signalChan, syscall.SIGTERM, syscall.SIGINT)
	select {
	case sig := <-signalChan:
		log.Println("Received signal:", sig)
	case <-done:
		log.Println("Received done signal, Shutting down...")
	}
}

func AddRequestId(w http.ResponseWriter, r *http.Request) (context.Context, error) {
	canonicId, err := nanoid.Standard(21)
	if err != nil {
		return nil, err
	}
	id := canonicId()
	return WithRequestId(r.Context(), id), nil
}
