package Grpc

import (
	"fmt"
	"net"

	"google.golang.org/grpc"

	"github.com/hossein1376/BehKhan/catalogue/internal/handlers"
	"github.com/hossein1376/BehKhan/catalogue/pkg/configs"
	"github.com/hossein1376/BehKhan/catalogue/proto/cataloguePB"
)

func ServeGrpc(app *configs.Application) {
	defer func() {
		if err := recover(); err != nil {
			app.Logger.Error("failed to recover in grpc goroutine", "error", err)
		}
	}()

	lis, err := net.Listen("tcp", ":8003")
	if err != nil {
		app.Logger.Error("failed to start a grpc listener", "error", err)
	}

	s := grpc.NewServer()
	cataloguePB.RegisterBookServiceServer(s, &handlers.Server{Application: app})

	app.Logger.Info(fmt.Sprintf("starting gRPC server on port %d", app.Settings.Grpc.Port))

	err = s.Serve(lis)
	if err != nil {
		app.Logger.Error("failed to start a grpc server", "error", err)
	}
}
