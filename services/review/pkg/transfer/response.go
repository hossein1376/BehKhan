package transfer

import (
	"fmt"
	"log/slog"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/valyala/fasthttp"
)

type Response struct {
	*slog.Logger
}

func NewResponse(logger *slog.Logger) *Response {
	return &Response{Logger: logger}
}

func (res *Response) Respond(c *fiber.Ctx, statusCode int, message any) {
	if err := c.Status(statusCode).JSON(message); err != nil {
		res.logInternalError(c.Request(), err)
	}
}

// logInternalError logs the error details to the standard logger
func (res *Response) logInternalError(req *fasthttp.Request, err error) {
	res.Error("internal error",
		"error_message", err,
		"request_method", string(req.Header.Method()),
		"request_url", string(req.URI().Path()))
}

// StatusOKResponse means everything went as expected
func (res *Response) StatusOKResponse(c *fiber.Ctx, data any) error {
	return c.Status(http.StatusOK).JSON(data)
}

// StatusCreatedResponse indicates that requested resource(s) have been successfully created
func (res *Response) StatusCreatedResponse(c *fiber.Ctx, data any) error {
	return c.Status(http.StatusCreated).JSON(data)
}

// StatusNoContentResponse means the operation was successful, and server has nothing more to say about it
func (res *Response) StatusNoContentResponse(c *fiber.Ctx) error {
	return c.Status(http.StatusNoContent).JSON(nil)
}

// BadRequestResponse indicates that the request has been deemed unacceptable by server
func (res *Response) BadRequestResponse(c *fiber.Ctx, err error) error {
	msg := http.StatusText(http.StatusBadRequest)
	if err != nil {
		msg = err.Error()
	}

	response := fiber.Map{"message": msg}
	return c.Status(http.StatusBadRequest).JSON(response)
}

// UnauthorizedResponse responds when user is not authorized
func (res *Response) UnauthorizedResponse(c *fiber.Ctx) error {
	return c.Status(http.StatusUnauthorized).JSON(http.StatusText(http.StatusUnauthorized))
}

// ForbiddenResponse indicates that the action is not allowed
func (res *Response) ForbiddenResponse(c *fiber.Ctx) error {
	return c.Status(http.StatusForbidden).JSON(http.StatusText(http.StatusForbidden))
}

// NotFoundResponse will return 404 with, if provided, the error message.
// Otherwise, it'll return the classic 404 error message.
func (res *Response) NotFoundResponse(c *fiber.Ctx, err error) error {
	msg := http.StatusText(http.StatusNotFound)
	if err != nil {
		msg = err.Error()
	}

	response := fiber.Map{"message": msg}
	return c.Status(http.StatusNotFound).JSON(response)
}

// MethodNotAllowedResponse is returned when the request's method is not acceptable
func (res *Response) MethodNotAllowedResponse(c *fiber.Ctx) error {
	response := fiber.Map{"message": fmt.Sprintf("the %s method is not supported for this content", c.Request().Header.Method())}

	return c.Status(http.StatusMethodNotAllowed).JSON(response)
}

// DuplicateRequestResponse kicks in when the request was already handled before
func (res *Response) DuplicateRequestResponse(c *fiber.Ctx) error {
	response := fiber.Map{"message": "the content you're trying to add already exists"}

	return c.Status(http.StatusNotAcceptable).JSON(response)
}

// InternalServerErrorResponse indicates something has gone wrong unexpectedly
func (res *Response) InternalServerErrorResponse(c *fiber.Ctx, err error) error {
	res.logInternalError(c.Request(), err)

	response := fiber.Map{"message": http.StatusText(http.StatusInternalServerError)}
	return c.Status(http.StatusInternalServerError).JSON(response)
}
