package serde

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/hossein1376/BehKhan/catalogue/pkg/errs"
)

// Status will return HTTP status code of error, if it has been wrapped with `errs` package.
// As a second value, error's contextual message wrapped in a gin.H object will be returned. In the case of an empty
// message, text of the HTTP status code is used.
//
// If input err is nil, status code 204 and nil response is returned. In case error was not wrapped by `errs` package,
// status 500 (InternalServerError) and its text will be returned as a default.
func Status(err error) (int, gin.H) {
	if err == nil {
		return http.StatusNoContent, nil
	}
	var e *errs.Error
	if errors.As(err, &e) {
		msg := e.Message
		if msg == "" {
			msg = http.StatusText(e.HttpStatusCode)
		}
		return e.HttpStatusCode, gin.H{"message": msg}
	}
	return http.StatusInternalServerError, gin.H{"message": http.StatusText(http.StatusInternalServerError)}
}
