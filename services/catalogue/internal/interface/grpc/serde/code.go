package serde

import (
	"errors"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/hossein1376/BehKhan/catalogue/pkg/errs"
)

// Code will return an error enriched by gRPC code and message, if it has been wrapped with `errs` package. In the
// case of an empty message, text of the gRPC code is used.
//
// If error input is nil, nil is returned. In case error was not wrapped by `errs` package, status 13 (Internal) and its
// text will be utilized by default.
func Code(err error) error {
	if err == nil {
		return nil
	}
	var e errs.Error
	if errors.As(err, &e) {
		msg := e.Message
		if msg == "" {
			msg = e.GrpcStatusCode.String()
		}
		return status.Error(e.GrpcStatusCode, msg)
	}
	return status.Error(codes.Internal, codes.Internal.String())
}
