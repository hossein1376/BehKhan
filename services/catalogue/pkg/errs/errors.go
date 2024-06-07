// Package errs exposes Error object which implements error interface.
// It exports functions covering different scenarios, with the possibility of adding a second argument to provide a
// message to the clients.
package errs

import (
	"fmt"
	"net/http"

	"google.golang.org/grpc/codes"
)

// Error is an object used for communicating error's context between layers.
// It contains an error, a Message to return to the clients, alongside HTTP and gRPC status codes.
type Error struct {
	Err            error
	Message        string
	HttpStatusCode int
	GrpcStatusCode codes.Code
}

// Error returns error's text, prefixed by its HTTP and gRPC status codes.
// Example:
//
//	[400][3] Bad Input
//
// If no error is present, it defaults to HTTP status text.
func (e Error) Error() string {
	if e.Err == nil {
		return http.StatusText(e.HttpStatusCode)
	}
	return fmt.Sprintf("[%d][%d] %s", e.HttpStatusCode, e.GrpcStatusCode, e.Err)
}

// Unwrap returns the underlying Err.
func (e Error) Unwrap() error {
	return e.Err
}

// getMsg extracts the optional Message, or an empty string if there is none.
// If more than one message is provided, only the first one will be used.
func getMsg(msg []string) string {
	message := ""
	if len(msg) != 0 {
		message = msg[0]
	}
	return message
}

// BadRequest indicates client has provided invalid arguments, and must correct them before retrying.
//
// HTTP: 400
// GRPC: 3
func BadRequest(err error, msg ...string) Error {
	return Error{
		Err:            err,
		Message:        getMsg(msg),
		HttpStatusCode: http.StatusBadRequest,
		GrpcStatusCode: codes.InvalidArgument,
	}
}

// Unauthorized indicates the request does not have valid authentication credentials for the operation.
//
// HTTP: 401
// GRPC: 16
func Unauthorized(err error, msg ...string) Error {
	return Error{
		Err:            err,
		Message:        getMsg(msg),
		HttpStatusCode: http.StatusUnauthorized,
		GrpcStatusCode: codes.Unauthenticated,
	}
}

// Forbidden indicates the caller does not have permission to execute the specified operation.
//
// HTTP: 403
// GRPC: 7
func Forbidden(err error, msg ...string) Error {
	return Error{
		Err:            err,
		Message:        getMsg(msg),
		HttpStatusCode: http.StatusForbidden,
		GrpcStatusCode: codes.PermissionDenied,
	}
}

// NotFound means some requested entity was not found.
//
// HTTP: 404
// GRPC: 5
func NotFound(err error, msg ...string) Error {
	return Error{
		Err:            err,
		Message:        getMsg(msg),
		HttpStatusCode: http.StatusNotFound,
		GrpcStatusCode: codes.NotFound,
	}
}

// Exists means operation was unsuccessful because one or more such entities already existed.
//
// HTTP: 409
// GRPC: 6
func Exists(err error, msg ...string) Error {
	return Error{
		Err:            err,
		Message:        getMsg(msg),
		HttpStatusCode: http.StatusConflict,
		GrpcStatusCode: codes.AlreadyExists,
	}
}

// Conflict indicates operation was rejected because the request is in conflict with the system's current state.
//
// HTTP: 409
// GRPC: 9
func Conflict(err error, msg ...string) Error {
	return Error{
		Err:            err,
		Message:        getMsg(msg),
		HttpStatusCode: http.StatusConflict,
		GrpcStatusCode: codes.FailedPrecondition,
	}
}

// TooMany indicates some resource has been exhausted, and client may need to wait some time before retrying.
//
// HTTP: 429
// GRPC: 8
func TooMany(err error, msg ...string) Error {
	return Error{
		Err:            err,
		Message:        getMsg(msg),
		HttpStatusCode: http.StatusTooManyRequests,
		GrpcStatusCode: codes.ResourceExhausted,
	}
}

// Internal means something has gone wrong in the server's side.
//
// HTTP: 500
// GRPC: 19
func Internal(err error, msg ...string) Error {
	return Error{
		Err:            err,
		Message:        getMsg(msg),
		HttpStatusCode: http.StatusInternalServerError,
		GrpcStatusCode: codes.Internal,
	}
}

// Timeout means a timeout has been reached. The operation may have been completed successfully or not.
//
// HTTP: 504
// GRPC: 4
func Timeout(err error, msg ...string) Error {
	return Error{
		Err:            err,
		Message:        getMsg(msg),
		HttpStatusCode: http.StatusGatewayTimeout,
		GrpcStatusCode: codes.DeadlineExceeded,
	}
}
