package grpc

import (
	"context"
	"log/slog"

	grpcMiddleware "github.com/grpc-ecosystem/go-grpc-middleware/v2"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/logging"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/recovery"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/hossein1376/BehKhan/catalogue/pkg/reqID"
	"github.com/hossein1376/BehKhan/catalogue/pkg/slogger"
)

type interceptors struct {
	logger *slog.Logger
}

func newInterceptors(logger *slog.Logger) interceptors {
	return interceptors{logger: logger}
}

func (i interceptors) recoverHandler() recovery.Option {
	return recovery.WithRecoveryHandlerContext(func(ctx context.Context, p any) (err error) {
		i.logger.ErrorContext(ctx, "panic in gRPC server", slog.Any("msg", p))
		return status.Errorf(codes.Internal, codes.Internal.String())
	})
}

func (i interceptors) loggerHandler() logging.Logger {
	return func(l *slog.Logger) logging.Logger {
		return logging.LoggerFunc(func(ctx context.Context, lvl logging.Level, msg string, fields ...any) {
			l.Log(ctx, slog.Level(lvl), msg, fields...)
		})
	}(i.logger)
}

func (i interceptors) unaryRequestID(
	ctx context.Context, req any, _ *grpc.UnaryServerInfo, handler grpc.UnaryHandler,
) (any, error) {
	id, err := reqID.NewRequestID()
	if err != nil {
		i.logger.ErrorContext(ctx, "error generating request id", slog.Any("error", err))
		return handler(ctx, req)
	}

	// put request id inside context
	ctx = context.WithValue(ctx, reqID.RequestIDKey, id)

	// include request_id in logs
	ctx = slogger.WithAttrs(ctx, slog.String("request_id", id))

	return handler(ctx, req)
}

func (i interceptors) streamRequestID(
	srv any, stream grpc.ServerStream, _ *grpc.StreamServerInfo, handler grpc.StreamHandler,
) error {
	id, err := reqID.NewRequestID()
	if err != nil {
		i.logger.ErrorContext(stream.Context(), "error generating request id", slog.Any("error", err))
		return handler(srv, stream)
	}

	// put request id inside context
	ctx := context.WithValue(stream.Context(), reqID.RequestIDKey, id)

	// include request_id in logs
	ctx = slogger.WithAttrs(ctx, slog.String("request_id", id))

	// wrap new context with the gRPC stream's context
	wrappedStream := grpcMiddleware.WrapServerStream(stream)
	wrappedStream.WrappedContext = ctx
	return handler(srv, wrappedStream)
}
