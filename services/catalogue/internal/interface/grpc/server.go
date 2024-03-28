package grpc

import (
	"fmt"
	"log/slog"
	"net"

	"google.golang.org/grpc"

	"github.com/hossein1376/BehKhan/catalogue/internal/domain/services"
)

type Server struct {
	srv *grpc.Server
}

func NewServer() *Server {
	srv := grpc.NewServer()
	return &Server{
		srv: srv,
	}
}

func (s *Server) Start(addr string) error {
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		return fmt.Errorf("new Listener: %w", err)
	}
	return s.srv.Serve(lis)
}

func (s *Server) Stop() error {
	s.srv.GracefulStop()
	return nil
}

func (s *Server) Mount(srvc services.Service, logger *slog.Logger) {
	NewBooksHndlr(s.srv, srvc, logger)
}
