package serverAPI

import (
	"context"
	"net/http"
	"time"
)

type ServerApi struct {
	server *http.Server
}

func (s *ServerApi) Run(handler http.Handler, port string) error {
	s.server = &http.Server{
		Addr:           port,
		Handler:        handler,
		MaxHeaderBytes: 1 << 20,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
	}
	return s.server.ListenAndServe()
}

func (s *ServerApi) Stop(ctx context.Context) error {
	return s.server.Shutdown(ctx)
}
