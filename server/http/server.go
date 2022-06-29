package http

import (
	"context"
	"net/http"

	"github.com/bricks-cloud/noname-gateway-project/common"
	"github.com/rs/zerolog/log"
)

// Server contains a http server
type Server struct {
	server *http.Server
}

// NewServer creates a new gateway server
func NewServer(hd http.Handler, sc common.ServerConfiguration) *Server {
	return &Server{
		server: &http.Server{
			Addr:              ":8080",
			Handler:           hd,
			ReadTimeout:       sc.GetReadTimeout(),
			ReadHeaderTimeout: sc.GetReadHeaderTimeout(),
			WriteTimeout:      sc.GetWriteTimeout(),
			IdleTimeout:       sc.GetIdleTimeout(),
		},
	}
}

// Start starts the webserver for the gateway
func (s *Server) Start() error {
	log.Info().Msg("starting the gateway http server...")
	return s.server.ListenAndServe()
}

// Stop stops the webserver for the gateway
func (s *Server) Stop(ctx context.Context) error {
	log.Ctx(ctx).Info().Msg("stopping the gateway http server...")
	return s.server.Shutdown(ctx)
}
