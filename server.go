package smartbridge_service

import (
	"context"
	"net/http"
	"time"
)

type Server struct {
	httpServer *http.Server
}

type HttpsServer struct {
	Port string
	Cert string
	Key  string
}

func (s *Server) Run(httpServer HttpsServer, handler http.Handler) error {
	s.httpServer = &http.Server{
		Addr:           ":" + httpServer.Port,
		Handler:        handler,
		MaxHeaderBytes: 1 << 20,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
	}

	return s.httpServer.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
