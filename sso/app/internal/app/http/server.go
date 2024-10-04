package http

import (
	"context"
	"github.com/pkg/errors"

	"vss/sso/internal/transport/http"
)

type Server struct {
	userHandler http.UserHandler
	ctx         context.Context
}

func NewServer(ctx context.Context) (*Server, error) {
	srv := Server{}
	err := srv.MapHandler()

	if err != nil {
		return nil, errors.Wrapf(err, "could not map handlers")
	}

	return &Server{}, nil
}

func (s *Server) MustRun() {

}

func (s *Server) Stop() {

}
