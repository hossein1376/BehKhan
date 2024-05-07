package grpc

import (
	"fmt"
	"log/slog"
	"net"

	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/logging"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/recovery"
	"google.golang.org/grpc"

	"github.com/hossein1376/BehKhan/catalogue/internal/domain/service"
	"github.com/hossein1376/BehKhan/catalogue/internal/interface/config"
	"github.com/hossein1376/BehKhan/catalogue/internal/interface/grpc/bookshndlr"
)

type Server struct {
	srv  *grpc.Server
	addr string
}

func NewServer(srvc service.Service, logger *slog.Logger, cfg config.GRPC) *Server {
	// first, manage gRPC interceptors
	var unary []grpc.UnaryServerInterceptor
	var stream []grpc.StreamServerInterceptor
	i := newInterceptors(logger)

	// requestID interceptor must go first, if enabled
	if cfg.RequestID {
		unary = append(unary, i.unaryRequestID)
		stream = append(stream, i.streamRequestID)
	}

	// optional logger interceptor
	if cfg.Logger {
		unary = append(unary, logging.UnaryServerInterceptor(i.loggerHandler()))
		stream = append(stream, logging.StreamServerInterceptor(i.loggerHandler()))
	}

	// default recover interceptor
	unary = append(unary, recovery.UnaryServerInterceptor(i.recoverHandler()))
	stream = append(stream, recovery.StreamServerInterceptor(i.recoverHandler()))

	// create a new gRPC server instance
	srv := grpc.NewServer(
		grpc.ChainUnaryInterceptor(unary...),
		grpc.ChainStreamInterceptor(stream...),
	)

	// mount services
	bookshndlr.New(srv, srvc, logger)

	return &Server{
		srv:  srv,
		addr: cfg.Addr,
	}
}

func (s *Server) Start() error {
	lis, err := net.Listen("tcp", s.addr)
	if err != nil {
		return fmt.Errorf("new Listener: %w", err)
	}
	return s.srv.Serve(lis)
}

func (s *Server) Stop() {
	s.srv.GracefulStop()
}
