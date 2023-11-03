package Grpc

import (
	"fmt"
	"net"
	"syscall"

	"google.golang.org/grpc"

	"github.com/hossein1376/BehKhan/review/pkg/config"
)

func ServeGrpc(app *config.Application) {
	// catch panics, if any were to happen
	defer func() {
		if err := recover(); err != nil {
			app.Logger.Error("failed to recover in grpc goroutine", "error", err)
		}
	}()

	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", app.Settings.Grpc.Port))
	if err != nil {
		app.Logger.Error("failed to start a grpc listener", "error", err)
		return
	}

	s := grpc.NewServer()

	// graceful shutdown
	go func() {
		<-app.Signals.ShutdownGRPC
		app.Logger.Debug("GRPC graceful shutdown started")

		s.GracefulStop()

		app.Logger.Debug("GRPC graceful shutdown finished")
		app.Signals.ShutdownGRPC <- syscall.Signal(0)
	}()

	app.Logger.Info(fmt.Sprintf("starting gRPC server on port %s", app.Settings.Grpc.Port))
	err = s.Serve(lis)
	if err != nil {
		app.Logger.Error("failed to start grpc server", "error", err)
		return
	}
}
