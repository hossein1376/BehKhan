package transfer

import (
	"fmt"
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	*slog.Logger
}

func NewResponse(logger *slog.Logger) *Response {
	return &Response{Logger: logger}
}

func (res *Response) Responsd(c *gin.Context, statusCode int, message any) {
	c.JSON(statusCode, message)
}

// logInternalError logs the error details to the standard logger
func (res *Response) logInternalError(req *http.Request, err error) {
	res.Error("internal error",
		"error_message", err,
		"request_method", req.Method,
		"request_url", req.URL.String())
}

// StatusOKResponse means everything went as expected
func (res *Response) StatusOKResponse(c *gin.Context, data any) {
	c.JSON(http.StatusOK, data)
}

// StatusCreatedResponse indicates that requested resource(s) have been successfully created
func (res *Response) StatusCreatedResponse(c *gin.Context, data any) {
	c.JSON(http.StatusCreated, data)
}

// StatusNoContentResponse means the operation was successful, and server has nothing more to say about it
func (res *Response) StatusNoContentResponse(c *gin.Context) {
	c.JSON(http.StatusNoContent, nil)
}

// BadRequestResponse indicates that the request has been deemed unacceptable by server
func (res *Response) BadRequestResponse(c *gin.Context, err error) {
	msg := http.StatusText(http.StatusBadRequest)
	if err != nil {
		msg = err.Error()
	}

	response := gin.H{"message": msg}
	c.JSON(http.StatusBadRequest, response)
}

// UnauthorizedResponse responds when user is not authorized
func (res *Response) UnauthorizedResponse(c *gin.Context) {
	c.JSON(http.StatusUnauthorized, http.StatusText(http.StatusUnauthorized))
}

// ForbiddenResponse indicates that the action is not allowed
func (res *Response) ForbiddenResponse(c *gin.Context) {
	c.JSON(http.StatusForbidden, http.StatusText(http.StatusForbidden))
}

// NotFoundResponse will return 404 with, if provided, the error message.
// Otherwise, it'll return the classic 404 error message.
func (res *Response) NotFoundResponse(c *gin.Context, err error) {
	msg := http.StatusText(http.StatusNotFound)
	if err != nil {
		msg = err.Error()
	}

	response := gin.H{"message": msg}
	c.JSON(http.StatusNotFound, response)
}

// MethodNotAllowedResponse is returned when the request's method is not acceptable
func (res *Response) MethodNotAllowedResponse(c *gin.Context) {
	response := gin.H{"message": fmt.Sprintf("the %s method is not supported for this content", c.Request.Method)}
	c.JSON(http.StatusMethodNotAllowed, response)
}

// DuplicateRequestResponse kicks in when the request was already handled before
func (res *Response) DuplicateRequestResponse(c *gin.Context) {
	response := gin.H{"message": "the content you're trying to add already exists"}
	c.JSON(http.StatusNotAcceptable, response)
}

// InternalServerErrorResponse indicates something has gone wrong unexpectedly
func (res *Response) InternalServerErrorResponse(c *gin.Context, err error) {
	res.logInternalError(c.Request, err)

	response := gin.H{"message": http.StatusText(http.StatusInternalServerError)}
	c.JSON(http.StatusInternalServerError, response)
}
