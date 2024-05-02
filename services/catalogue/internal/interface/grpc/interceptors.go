package grpc

import (
	"context"
	"log/slog"

	grpcMiddleware "github.com/grpc-ecosystem/go-grpc-middleware/v2"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/logging"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/recovery"
	"github.com/oklog/ulid/v2"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/hossein1376/BehKhan/catalogue/pkg/slogger"
)

const RequestID = "request_id"

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
	id, err := ulid.New(ulid.Now(), ulid.DefaultEntropy())
	if err != nil {
		i.logger.ErrorContext(ctx, "error generating ulid", slog.Any("err", err))
		return handler(ctx, req)
	}
	reqID := id.String()
	ctx = context.WithValue(ctx, RequestID, reqID)
	return handler(slogger.WithAttrs(ctx, slog.String("request_id", reqID)), req)
}

func (i interceptors) streamRequestID(
	srv any, stream grpc.ServerStream, _ *grpc.StreamServerInfo, handler grpc.StreamHandler,
) error {
	id, err := ulid.New(ulid.Now(), ulid.DefaultEntropy())
	if err != nil {
		i.logger.ErrorContext(stream.Context(), "error generating ulid", slog.Any("err", err))
		return handler(srv, stream)
	}
	reqID := id.String()

	// put request_id inside context
	ctx := context.WithValue(stream.Context(), RequestID, reqID)

	// include request_id in logs
	newCtx := slogger.WithAttrs(ctx, slog.String("request_id", reqID))

	// wrap new context with the gRPC stream's context
	wrappedStream := grpcMiddleware.WrapServerStream(stream)
	wrappedStream.WrappedContext = newCtx
	return handler(srv, wrappedStream)
}
